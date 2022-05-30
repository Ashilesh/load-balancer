package protocol

import (
	"io"
	"net"

	"github.com/Ashilesh/load-balancer/logs"
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
		// TODO: delete node url from structure
		logs.Error("unable to connect to server", err)
		return
	}

	// if server disconnects then send data on this channel
	closeConnChan := make(chan int)
	// if client disconnects send data to this channel
	closeClientConn := make(chan int)
	go func() {
		_, err := io.Copy(serverConn, clientConn)
		closeClientConn <- 2
		if val := <-closeClientConn; val != 1 && err != nil {
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
	logs.Info(clientConn.RemoteAddr().String(), "disconnected")
}

var http *Http

func GetHttpProtocol() *Http {
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
