package protocol

import (
	"fmt"
	"net"
)

type IProtocol interface {
	CreateDuplexConnection(net.Conn, string)
}

func GetProto(protocolType string) IProtocol {
	switch protocolType {
	case "websocket":
		fmt.Println("Protocol type: websocket")
		return GetWebsocketFactory()
	default:
		fmt.Println("Protocol type: http")
		return GetHttpFactory()
	}
}
