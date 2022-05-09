package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("Initializing Server...")

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started at port", 8000)

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(conn net.Conn) {
			fmt.Println(conn.RemoteAddr().String())
			io.Copy(conn, conn)
			conn.Close()
		}(conn)
	}
}
