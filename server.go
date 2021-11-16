package main

import (
	"btc-node/addrmgr"
	"btc-node/blockchain"
	"btc-node/blockchain/indexers"
	"btc-node/chaincfg"
	"btc-node/chaincfg/chainhash"
	"btc-node/connmgr"
	"btc-node/database"
	"btc-node/mempool"
	"btc-node/mining/cpuminer"
	"btc-node/netsync.go"
	"btc-node/peer"
	"btc-node/txscript"
	"btc-node/wire"
	"sync"

	"github.com/btcsuite/btcutil/bloom"
)

type serverPeer struct {
	feeFilter int64

	*peer.Peer

	connReq        *connmgr.ConnReq
	server         *server
	persistent     bool
	continueHash   chainhash.Hash
	relayMtx       sync.Mutex
	disableRelayTx bool
	sentAddrs      bool
	isWhitelisted  bool
	filter         *bloom.Filter
	addressesMtx   sync.Mutex
	knownAddresses map[string]struct{}
	banScore       connmgr.DynamicBanScore
	quit           chan struct{}

	txProcessed    chan struct{}
	blockProcessed chan struct{}
}

type broadcastMsg struct {
	message      wire.Message
	excludePeers []*serverPeer
}

type relayMsg struct {
	invVect *wire.InvVect
	data    interface{}
}

type updatePeerHeightsMsg struct {
	newHash    *chainhash.Hash
	newHeight  int32
	originPeer *peer.Peer
}

type cfHeaderKV struct {
	blockHash    chainhash.Hash
	filterHeader chainhash.Hash
}

type server struct {
	bytesReceived uint64
	bytesSent     uint64
	started       int32
	shutdown      int32
	shutdownSched int32
	startupTime   int64

	chainParams          *chaincfg.Params
	addrManager          *addrmgr.AddrManager
	connManager          *connmgr.ConnManager
	sigCache             *txscript.SigCache
	hashCache            *txscript.HashCache
	syncManager          *netsync.SyncManager
	chain                *blockchain.BlockChain
	txMemPool            *mempool.TxPool
	cpuminer             *cpuminer.CPUMiner
	modifyRebroadcastInv chan interface{}
	newPeers             chan *serverPeer
	donePeers            chan *serverPeer
	banPeers             chan *serverPeer
	query                chan interface{}
	relayInv             chan relayMsg
	broadcast            chan broadcastMsg
	updatePeerHeightsMsg chan updatePeerHeightsMsg
	wg                   sync.WaitGroup
	quit                 chan struct{}
	nat                  NAT
	db                   database.DB
	timeSource           blockchain.MedianTimeSource
	services             wire.ServiceFlag

	txIndex   *indexers.TxIndex
	addrIndex *indexers.AddrIndex
	cfIndex   *indexers.CfIndex

	feeEstimator *mempool.FeeEstimator

	cfCheckptCaches    map[wire.FilterType][]cfHeaderKV
	cfCheckptCachesMtx sync.RWMutex

	agentBlacklist []string

	agentWhitelist []string
}
