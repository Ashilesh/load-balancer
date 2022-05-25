package app

import (
	"fmt"
	"net"

	"github.com/Ashilesh/balancer/src/algo"
	"github.com/Ashilesh/balancer/src/config"
	"github.com/Ashilesh/balancer/src/protocol"
)

var configuration config.Configuration
var balancingAlgo algo.Algo

func init() {
	configuration = config.GetConfig()
	balancingAlgo = algo.GetAlgo()

	for _, url := range configuration.Nodes {
		balancingAlgo.Add(url)
	}
}

func Run() {
	createServer()
}

func createServer() {

	fmt.Println("networkType : ", configuration.NetworkType)
	fmt.Println("host: ", configuration.Host)
	ln, err := net.Listen(configuration.NetworkType, configuration.Host+":"+configuration.Port)
	if err != nil {
		panic("ERROR: unable to listen")
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERROR: unable to accept connection", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(clientConn net.Conn) {
	fmt.Println("client IP: ", clientConn.RemoteAddr().String())
	nodeUrl := balancingAlgo.GetUrl(clientConn.RemoteAddr().String())
	fmt.Println("INFO: Redirecting to", nodeUrl)

	proto := protocol.GetProto(configuration.Protocol)

	proto.CreateDuplexConnection(clientConn, nodeUrl)
}
