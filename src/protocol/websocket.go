package protocol

import (
	"io"
	"net"

	"github.com/Ashilesh/load-balancer/logs"
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
		// TODO: remove url from structure
		logs.Error(err)
		return
	}

	// if server disconnects then send data on this channel
	closeConnChan := make(chan int)

	go func() {
		_, err := io.Copy(serverConn, clientConn)
		if err != nil {
			logs.Error(err)
		}
	}()

	// use channel to communicate and close connections
	go func() {
		_, err := io.Copy(clientConn, serverConn)

		if err != nil {
			logs.Error(err)
		}
		closeConnChan <- 1

	}()

	// wait for connection from LB to server disconnects
	<-closeConnChan

	serverConn.Close()
	clientConn.Close()

	logs.Info(clientConn.RemoteAddr().String(), "disconnected")
}

var websocket *Websocket

func GetWebsocketProtocol() *Websocket {
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
