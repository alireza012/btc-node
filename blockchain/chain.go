package blockchain

import (
	"btc-nod/txscript"
	"btc-node/chaincfg"
	"btc-node/chaincfg/chainhash"
	"sync"
	"time"

	"github.com/btcsuite/btcd/database"
	"github.com/btcsuite/btcutil"
)

type orphanBlock struct {
	block      *btcutil.Block
	expiration time.Time
}

type BestState struct {
	Hash        chainhash.Hash
	Height      int32
	Bits        uint32
	BlockSize   uint64
	BlockWeight uint64
	NumTxns     uint64
	TotalTxns   uint64
	MedianTime  time.Time
}

type BlockChain struct {
	checkpoints         []chaincfg.Checkpoint
	checkpointsByHeight map[int32]*chaincfg.Checkpoint
	db                  database.DB
	chainParams         *chaincfg.Params
	timeSource          MedianTimeSource
	sigCache            *txscript.SigCache
	indexManager        IndexManager
	hashCache           *txscript.HashCache

	minRetargetTimespan int64
	maxRetargetTimespan int64
	blocksPerRetarget   int32

	chainLock sync.RWMutex

	index     *blockIndex
	bestChain *chainView

	orphanLock   sync.RWMutex
	orphans      map[chainhash.Hash]*orphanBlock
	prevOrphans  map[chainhash.Hash][]*orphanBlock
	oldestOrphan *orphanBlock

	nextCheckpoint *chaincfg.Checkpoint
	checkpointNode *blockNode

	stateLock     sync.RWMutex
	stateSnapshot *BestState

	warningsCaches   []thresholdStateCache
	deploymentCaches []thresholdStateCache

	unknownRulesWarned bool

	notificationsLock sync.RWMutex
	notifications     []NotificationCallback
}

type IndexManager interface {
	Init(*BlockChain, <-chan struct{}) error

	ConnectBlock(database.Tx, *btcutil.Block, []SpentTxOut) error

	DisconnectBlock(database.Tx, *btcutil.Block, []SpentTxOut) error
}
