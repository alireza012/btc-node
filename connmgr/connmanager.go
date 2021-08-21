package connmgr

import (
	"net"
	"sync"
	"time"
)

type ConnState uint8

type ConnReq struct {
	id uint64

	Addr      net.Addr
	Permanent bool

	conn       net.Conn
	state      ConnState
	stateMtx   sync.RWMutex
	retryCount uint32
}

type Config struct {
	Listeners []net.Listener

	OnAccept func(net.Conn)

	TargetOutbound uint32

	RetryDuration time.Duration

	OnConnection func(*ConnReq, net.Conn)

	OnDisconnection func(*ConnReq)

	GetNewAddress func() (net.Addr, error)

	Dial func(net.Addr) (net.Conn, error)
}

type ConnManager struct {
	connReqCount uint64
	start        int32
	stop         int32

	cfg            Config
	wg             sync.WaitGroup
	failedAttempts uint64
	requests       chan interface{}
	quit           chan struct{}
}
