package main

import (
	"fmt"
	// "io"
	// "balancer/algo"
	"log"
	"net"

	"github.com/Ashilesh/balancer/src/algo"
)

func main() {
	fmt.Println("Initializing Server...")

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started at port", 8000)

	defer ln.Close()

	// TODO: use interface so any algo can be used
	// ch := algo.GetConsistetnHash()
	// node := algo.GetNode("https://localhost:8080")

	// ch.Add(node)
	// ch.GetArray()

	// node = algo.GetNode("2")
	// ch.Add(node)
	// ch.GetArray()

	// node = algo.GetNode("10")
	// ch.Add(node)
	// ch.GetArray()

	// node = algo.GetNode("4")
	// ch.Add(node)
	// ch.GetArray()

	// node = algo.GetNode("5")
	// ch.Add(node)
	// ch.GetArray()

	// for {
	// 	conn, err := ln.Accept()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	go func(conn net.Conn) {
	// 		fmt.Println(conn.RemoteAddr().String())
	// 		io.Copy(conn, conn)
	// 		conn.Close()
	// 	}(conn)
	// }
}
