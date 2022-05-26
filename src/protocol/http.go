package protocol

import (
	"fmt"
	"io"
	"net"
)

type Http struct {
	networkType string
}

func (h *Http) Default_values() {
	h.networkType = "tcp"
}

func (h *Http) CreateDuplexConnection(clientConn net.Conn, serverUrl string) {
	serverConn, err := net.Dial(h.networkType, serverUrl)
	if err != nil {
		// TODO: handle error
		fmt.Println("error", err)
		return
	}

	// if server disconnects then send data on this channel
	closeConnChan := make(chan int)
	closeClientConn := make(chan int)
	go func() {
		fmt.Println("sending data to server")
		count, err := io.Copy(serverConn, clientConn)
		closeClientConn <- 2
		if val := <-closeClientConn; val != 1 {
			// error msg
			fmt.Println("error while copying to target", err)
		}
		fmt.Println("client -> server ", count)

		// closeConnChan <- 1
		fmt.Println("client -> server complete")
	}()

	// use channel to communicate and close connections
	go func() {
		// var buf bytes.Buffer
		fmt.Println("receiving data from server")
		count, err := io.Copy(clientConn, serverConn)
		fmt.Println("server -> client ", count)

		if err != nil {
			fmt.Println("error while copying to client")
		}

		closeConnChan <- 1
		fmt.Println("server -> client complete")

	}()

	isServerDisconnected := false
	isClientDisconnected := false
	for {
		if isServerDisconnected && isClientDisconnected {
			break
		}

		select {
		case <-closeConnChan:
			isServerDisconnected = true
			serverConn.Close()
			clientConn.Close()
		case <-closeClientConn:
			isClientDisconnected = true
			if isServerDisconnected {
				closeClientConn <- 1
			} else {
				closeClientConn <- 2
			}
		}
	}

	fmt.Println("Complete")
}

var http *Http

func GetHttpFactory() *Http {
	if http != nil {
		return http
	}

	h := &Http{}
	setHttp(h)
	http.Default_values()

	return http
}

func setHttp(h *Http) {
	http = h
}
