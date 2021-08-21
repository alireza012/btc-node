package peer

import (
	"btc-node/chaincfg"
	"btc-node/chaincfg/chainhash"
	"btc-node/wire"
	"net"
)

type MessageListeners struct {
	OnGetAddr func(p *Peer, msg *wire.MsgGetAddr)

	OnAddr func(p *Peer, msg *wire.MsgAddr)

	OnPing func(p *Peer, msg *wire.MsgPing)

	OnPong func(p *Peer, msg *wire.MsgPong)

	OnMemPool func(p *Peer, msg *wire.MsgMemPool)

	OnTx func(p *Peer, msg *wire.MsgTx)

	OnBlock func(p *Peer, msg *wire.MsgBlock, buf []byte)

	OnCFilter func(p *Peer, msg *wire.MsgCFilter)

	OnCFHeaders func(p *Peer, msg *wire.MsgCFHeaders)

	OnCFCheckpt func(p *Peer, msg *wire.MsgCFCheckpt)

	OnInv func(p *Peer, msg *wire.MsgInv)
}

type Config struct {
	NewestBlock HashFunc

	HostToNetAddress HostToNetAddrFunc

	Proxy string

	UserAgent string

	UserAgentVersion string

	UserAgentComments []string

	ChainParams *chaincfg.Params

	Services wire.ServiceFlag

	ProtocolVersion uint32

	DisableRelayTx bool
}

type HashFunc func() (hash *chainhash.Hash, height int32, err error)

type HostToNetAddrFunc func(host string, port uint16, services wire.ServiceFlag) (*wire.NetAddress, error)

type Peer struct {
	bytesReceived uint64
	bytesSent     uint64
	lastRecv      int64
	lastSend      int64
	connected     int32
	disconnected  int32

	conn net.Conn

	addr string
	cfg
}
