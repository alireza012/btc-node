package cpuminer

import (
	"btc-node/blockchain"
	"btc-node/chaincfg"
	"btc-node/mining"
	"sync"

	"github.com/btcsuite/btcutil"
)

type Config struct {
	ChainParams *chaincfg.Params

	BlockTemplateGenerator *mining.BlkTmplGenerator

	MiningAddrs []btcutil.Address

	ProcessBlock func(*btcutil.Block, blockchain.BehaviorFlags) (bool, error)

	ConnectedCount func() int32

	IsCurrent func() bool
}

type CPUMiner struct {
	sync.Mutex
	g                 *mining.BlkTmplGenerator
	cfg               Config
	numWorkers        uint32
	started           bool
	discreteMining    bool
	submitBlockLock   sync.Mutex
	wg                sync.WaitGroup
	workerWg          sync.WaitGroup
	updateNumWorkers  chan struct{}
	queryHashesPerSec chan float64
	updateHashes      chan uint64
	speedMonitorQuit  chan struct{}
	quit              chan struct{}
}
