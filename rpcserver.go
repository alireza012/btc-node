package main

import "net"

type rpcServer struct {
	started  int32
	shutdown int32
	cfg
}

type rpcserverPeer interface {
	ToPeer()
}

type rpcserverConnManager interface {
	Connect(addr string, permanent bool) error

	RemoveByID(id int32) error

	RemoveByAddr(addr string) error

	DisconnectByID(id int32) error

	DisconnectByAddr(addr string) error

	ConnectedCount() int32

	NetTotals() (uint64, uint64)
}

type rpcserverConfig struct {
	Listeners []net.Listener

	StartupTime int64

	connMgr
}
