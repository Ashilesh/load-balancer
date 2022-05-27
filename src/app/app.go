package app

import (
	"net"

	"github.com/Ashilesh/load-balancer/algo"
	"github.com/Ashilesh/load-balancer/config"
	"github.com/Ashilesh/load-balancer/logs"
	"github.com/Ashilesh/load-balancer/protocol"
)

var configuration config.Configuration
var balancingAlgo algo.Algo
var proto protocol.IProtocol

func init() {
	configuration = config.GetConfig()

	// TODO: pass algo type from config file
	balancingAlgo = algo.GetAlgoFactory()

	for _, url := range configuration.Nodes {
		balancingAlgo.Add(url)
	}

	proto = protocol.GetProtoFactory(configuration.Protocol)
}

func Run() {
	createServer()
}

func createServer() {

	logs.Info("Starting server")
	logs.Info("Host ->", configuration.Host)
	logs.Info("Port ->", configuration.Port)
	logs.Info("Protocol ->", configuration.Protocol)

	ln, err := net.Listen(configuration.NetworkType, configuration.Host+":"+configuration.Port)
	if err != nil {
		logs.Fatal("Unable to listen ", configuration.Port)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			logs.Error("Unable to accept connection")
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(clientConn net.Conn) {

	nodeUrl := balancingAlgo.GetUrl(clientConn.RemoteAddr().String())

	logs.Info("Redirecting client ->", clientConn.RemoteAddr().String(), " to url ->", nodeUrl)

	proto.CreateDuplexConnection(clientConn, nodeUrl)
}
