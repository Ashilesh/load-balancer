package app

import (
	"fmt"
	"io"
	"net"

	"github.com/Ashilesh/balancer/src/algo"
	"github.com/Ashilesh/balancer/src/config"
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

func handleConnection(conn net.Conn) {
	nodeUrl := balancingAlgo.GetUrl(conn.RemoteAddr().String())
	fmt.Println("INFO: Redirecting to", nodeUrl)
	target, err := net.Dial(configuration.NetworkType, nodeUrl)
	if err != nil {
		panic(err)
	}

	go func(target, conn net.Conn) {
		fmt.Println("sending data to server")
		_, err := io.Copy(target, conn)
		if err != nil {
			fmt.Println("error while copying to target", err)
		}
	}(target, conn)

	// use channel to communicate and close connections
	go func() {
		fmt.Println("receiving data from server")
		_, err := io.Copy(conn, target)
		if err != nil {
			fmt.Println("error while copying to client")
		}
	}()
}
