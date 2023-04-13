package internal

import (
	"github.com/soulnov23/go-tool/pkg/log"
	"github.com/soulnov23/go-tool/pkg/net"
)

type RPCServer struct {
	FrameLog log.Logger
	RunLog   log.Logger
}

func (svr *RPCServer) OnRead(conn *net.TcpConn) {
}
