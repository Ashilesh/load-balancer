package main

import (
	"fmt"
	// "io"
	// "balancer/algo"
	"log"
	"net"

	"github.com/Ashilesh/balancer/src/algo"
	"github.com/Ashilesh/balancer/src/utils"
	// "github.com/Ashilesh/balancer/src/utils"
)

func main() {
	fmt.Println("Initializing Server...")

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started at port", 8000)

	defer ln.Close()

	consistentHash := algo.GetAlgo()
	consistentHash.Add("256")
	consistentHash.Add("afdfa")
	consistentHash.Add("4fnore")

	fmt.Println(consistentHash)

	consistentHash.Delete("3")

	fmt.Println(consistentHash)

	consistentHash.Add("localhost:1")
	consistentHash.Add("localhost:2")
	consistentHash.Add("localhost:3")

	fmt.Println(consistentHash.GetUrl("eewr"))
	fmt.Println(consistentHash.GetUrl("localho"))

	fmt.Println(consistentHash)
	fmt.Println(utils.GetHash("eewr"))
	fmt.Println(utils.GetHash("localho"))
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
