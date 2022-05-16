package app

import (
	"fmt"
	"github.com/Ashilesh/balancer/src/config"
	"github.com/Ashilesh/balancer/src/utils"
	"io"
	"net"
)

var configuration config.Configuration

func Run() {
	// TODO: create struct to store strings for command fields ex. config = "-config"
	arg, err := utils.GetCmdArgs("-config")
	if err != nil {
		panic("Configuration file path argument not found")
	}

	fmt.Println(arg)
	// createServer()
}

func createServer() {
	networkType := "tcp"
	addr := "127.0.0.1:8080"

	ln, err := net.Listen(networkType, addr)
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
	target, err := net.Dial("tcp", ":7071")
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
