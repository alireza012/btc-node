package netsync

import (
	"btc-node/blockchain"
	"btc-node/chaincfg"
	"btc-node/chaincfg/chainhash"
	"btc-node/mempool"
	peerpkg "btc-node/peer"
	"btc-node/wire"
	"container/list"
	"sync"
	"time"
)

type peerSyncState struct {
	syncCandidate   bool
	requestQueue    []*wire.InvVect
	requestedTxns   map[chainhash.Hash]struct{}
	requestedBlocks map[chainhash.Hash]struct{}
}

type SyncManager struct {
	PeerNotifier   PeerNotifier
	started        int32
	shutdown       int32
	chain          *blockchain.BlockChain
	txMemPool      *mempool.TxPool
	chainParams    *chaincfg.Params
	progressLogger *blockProgressLogger
	msgChan        chan interface{}
	wg             sync.WaitGroup
	quit           chan struct{}

	rejectedTxns     map[chainhash.Hash]struct{}
	requestedTxns    map[chainhash.Hash]struct{}
	requestedBlocks  map[chainhash.Hash]struct{}
	syncPeer         *peerpkg.Peer
	peerStates       map[*peerpkg.Peer]*peerSyncState
	lastProgressTime time.Time

	headersFirstMode bool
	headerList       *list.List
	startHeader      *list.Element
	nextCheckpoint   *chaincfg.Checkpoint

	feeEstimator *mempool.FeeEstimator
}
