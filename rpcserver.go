package main

import (
	"btc-node/blockchain"
	"btc-node/blockchain/indexers"
	"btc-node/chaincfg"
	"btc-node/chaincfg/chainhash"
	"btc-node/database"
	"btc-node/mempool"
	"btc-node/mining"
	"btc-node/mining/cpuminer"
	"btc-node/peer"
	"btc-node/wire"
	"crypto/sha256"
	"net"
	"sync"
	"time"

	"github.com/btcsuite/btcutil"
)

type gbtWorkState struct {
	sync.Mutex
	lastTxUpdate  time.Time
	lastGenerated time.Time
	prevHash      *chainhash.Hash
	minTimestamp  time.Time
	template      *mining.BlockTemplate
	notifyMap     map[chainhash.Hash]map[int64]chan struct{}
	timeSource    blockchain.MedianTimeSource
}

type rpcServer struct {
	started                int32
	shutdown               int32
	cfg                    rpcserverConfig
	authsha                [sha256.Size]byte
	limitauthsha           [sha256.Size]byte
	ntfnMgr                *wsNotificationManager
	numClients             int32
	statusLines            map[int]string
	statusLock             sync.RWMutex
	wg                     sync.WaitGroup
	gbtWorkState           *gbtWorkState
	helpCacher             *helpCacher
	requestProcessShutdown chan struct{}
	quit                   chan int
}

type rpcserverPeer interface {
	ToPeer() *peer.Peer

	IsTxRelayDisabled() bool

	BanScore() uint32

	FeeFilter() int64
}

type rpcserverConnManager interface {
	Connect(addr string, permanent bool) error

	RemoveByID(id int32) error

	RemoveByAddr(addr string) error

	DisconnectByID(id int32) error

	DisconnectByAddr(addr string) error

	ConnectedCount() int32

	NetTotals() (uint64, uint64)

	ConnectedPeers() []rpcserverPeer

	PersistentPeers() []rpcserverPeer

	BroadcastMessage(msg wire.Message)

	AddRebroadcastInventory(iv *wire.InvVect, data interface{})

	RelayTransactions(txns []*mempool.TxDesc)

	NodeAddresses() []*wire.NetAddress
}

type rpcserverSyncManager interface {
	IsCurrent() bool

	SubmitBlock(block *btcutil.Block, flags blockchain.BehaviorFlags) (bool, error)

	Pause() chan<- struct{}

	SyncPeerID() int32

	LocateHeaders(locators []*chainhash.Hash, hashStop *chainhash.Hash) []wire.BlockHeader
}

type rpcserverConfig struct {
	Listeners []net.Listener

	StartupTime int64

	ConnMgr rpcserverConnManager

	SyncMgr rpcserverSyncManager

	TimeSource  blockchain.MedianTimeSource
	Chain       *blockchain.BlockChain
	ChainParams *chaincfg.Params
	DB          database.DB

	TxMemPool *mempool.TxPool

	Generator *mining.BlkTmplGenerator
	CPUMiner  *cpuminer.CPUMiner

	TxIndex   *indexers.TxIndex
	AddrIndex *indexers.AddrIndex
	CFIndex   *indexers.CfIndex

	FeeEstimator *mempool.FeeEstimator
}
