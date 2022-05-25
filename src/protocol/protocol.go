package protocol

import "net"

type IProtocol interface {
	CreateDuplexConnection(net.Conn, string)
}

func GetProto(protocolType string) IProtocol {
	switch protocolType {
	case "websocket":
		return GetWebsocketFactory()
	// TODO: default is http
	default:
		return nil
	}
}
