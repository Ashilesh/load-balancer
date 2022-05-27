package protocol

import (
	"net"

	"github.com/Ashilesh/load-balancer/logs"
)

type IProtocol interface {
	CreateDuplexConnection(net.Conn, string)
}

func GetProtoFactory(protocolType string) IProtocol {
	switch protocolType {
	case "websocket":
		logs.Info("setting protocol to -> websocket")
		return GetWebsocketProtocol()
	default:
		logs.Info("setting protocol to -> http")
		return GetHttpProtocol()
	}
}
