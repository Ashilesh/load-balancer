package protocol

import (
	"fmt"
	"io"
	"net"
)

type Websocket struct {
	networkType string
}

func (w *Websocket) Default_values() {
	w.networkType = "tcp"
}

func (w *Websocket) CreateDuplexConnection(clientConn net.Conn, serverUrl string) {
	serverConn, err := net.Dial(w.networkType, serverUrl)
	if err != nil {
		// TODO: handle error
		fmt.Println("error", err)
		return
	}

	// if server disconnects then send data on this channel
	closeConnChan := make(chan int)

	go func() {
		fmt.Println("sending data to server")
		count, err := io.Copy(serverConn, clientConn)
		fmt.Println("client -> server ", count)
		if err != nil {
			fmt.Println("error while copying to target")
		}
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

	// wait for connection from LB to server disconnects
	<-closeConnChan

	serverConn.Close()
	clientConn.Close()

	fmt.Println("Complete")
}

var websocket *Websocket

func GetWebsocketFactory() *Websocket {
	if websocket != nil {
		return websocket
	}
	w := &Websocket{}
	setWebsocket(w)
	websocket.Default_values()

	return websocket
}

func setWebsocket(w *Websocket) {
	websocket = w
}
