package peer

import (
	"btc-node/chaincfg"
	"btc-node/chaincfg/chainhash"
	"btc-node/wire"
	"net"
	"sync"
	"time"

	"github.com/decred/dcrd/lru"
)

const (
	defaultTrickleInterval = 10 * time.Second
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

	OnHeaders func(p *Peer, msg *wire.MsgHeaders)

	OnNotFound func(p *Peer, msg *wire.MsgNotFound)

	OnGetData func(p *Peer, msg *wire.MsgGetData)

	OnGetBlocks func(p *Peer, msg *wire.MsgGetBlocks)

	OnGetHeaders func(p *Peer, msg *wire.MsgGetHeaders)

	OnGetCFCheckpt func(p *Peer, msg *wire.MsgGetCFCheckpt)

	OnFeeFilter func(p *Peer, msg *wire.MsgFeeFilter)

	OnFilterAdd func(p *Peer, msg *wire.MsgFilterAdd)

	OnFilterClear func(p *Peer, msg *wire.MsgFilterClear)

	OnFilterLoad func(p *Peer, msg *wire.MsgFilterLoad)

	OnMerkleBlock func(p *Peer, msg *wire.MsgMerkleBlock)

	OnVersion func(p *Peer, msg *wire.MsgVersion) *wire.MsgReject

	OnVerAck func(p *Peer, msg *wire.MsgVerAck)

	OnReject func(p *Peer, msg *wire.MsgReject)

	OnSendHeaders func(p *Peer, msg *wire.MsgSendHeaders)

	OnRead func(p *Peer, bytesRead int, msg wire.Message, err error)

	OnWrite func(p *Peer, bytesWritten int, msg wire.Message, err error)
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

	Listeners MessageListeners

	TrickleInterval time.Duration

	AllowSelfConns bool
}

type outMsg struct {
	msg      wire.Message
	doneChan chan<- struct{}
	encoding wire.MessageEncoding
}

type stallControlCmd uint8

type stallControlMsg struct {
	command stallControlCmd
	message wire.Message
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

	addr    string
	cfg     Config
	inbound bool

	flagsMtx             sync.Mutex
	na                   *wire.NetAddress
	id                   int32
	userAgent            string
	services             wire.ServiceFlag
	versionKnown         bool
	advertisedProtoVer   uint32
	protocolVersion      uint32
	sendHeadersPreferred bool
	verAckReceived       bool
	witnessEnabled       bool

	wireEncoding wire.MessageEncoding

	knownInventory     lru.Cache
	prevGetBlocksMtx   sync.Mutex
	prevGetBlocksBegin *chainhash.Hash
	prevGetBlocksStop  *chainhash.Hash
	prevGetHdrsMtx     sync.Mutex
	prevGetHdrsBegin   *chainhash.Hash
	prevGetHdrsStop    *chainhash.Hash

	statsMtx           sync.RWMutex
	timeOffset         int64
	timeConnected      time.Time
	startingHeight     int32
	lastBlock          int32
	lastAnnouncedBlock *chainhash.Hash
	lastPingNonce      uint64
	lastPingTime       time.Time
	lastPingMicros     int64

	stallControl  chan stallControlMsg
	outputQueue   chan outMsg
	sendQueue     chan outMsg
	sendDoneQueue chan struct{}
	outputInvChan chan *wire.InvVect
	inQuit        chan struct{}
	queueQuit     chan struct{}
	outQuit       chan struct{}
	quit          chan struct{}
}
