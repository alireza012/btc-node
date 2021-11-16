package mempool

import (
	"btc-nod/txscript"
	"btc-node/blockchain"
	"btc-node/blockchain/indexers"
	"btc-node/chaincfg"
	"btc-node/chaincfg/chainhash"
	"btc-node/mining"
	"btc-node/wire"
	"sync"
	"time"

	"github.com/btcsuite/btcutil"
)

const (
	DefaultBlockPrioritySize = 50000

	orphanTTL = time.Minute * 15

	orphanExpireScanInterval = time.Minute * 5

	MaxRBFSequence = 0xfffffffd

	MaxReplacementEvictions = 100
)

type Tag uint64

type Config struct {
	Policy Policy

	ChainParams *chaincfg.Params

	FetchUtxoView func(*btcutil.Tx) (*blockchain.UtxoViewpoint, error)

	BestHeight func() int32

	MedianTimePast func() time.Time

	CalcSequenceLock func(*btcutil.Tx, *blockchain.UtxoViewpoint) (*blockchain.SequenceLock, error)

	IsDeploymentActive func(deploymentID uint32) (bool, error)

	SigCache *txscript.SigCache

	HashCache txscript.HashCache

	AddrIndex *indexers.AddrIndex

	FeeEstimator *FeeEstimator
}

type Policy struct {
	MaxTxVersion int32

	DisableRelayPriority bool

	AcceptNonStd bool

	FreeTxRelayLimit float64

	MaxOrphanTxs int

	MaxOrphanTxSize int

	MaxSigOpCostPerTx int

	minRelayTxFee btcutil.Amount

	RejectReplacement bool
}

type TxDesc struct {
	mining.TxDesc

	StartingPriority float64
}

type orphanTx struct {
	tx         *btcutil.Tx
	tag        Tag
	expiration time.Time
}

type TxPool struct {
	lastUpdated int64

	mtx           sync.RWMutex
	cfg           Config
	pool          map[chainhash.Hash]*TxDesc
	orphans       map[chainhash.Hash]*orphanTx
	orphansByPrev map[wire.OutPoint]map[chainhash.Hash]*btcutil.Tx
	outpoints     map[wire.OutPoint]*btcutil.Tx
	pennyTotal    float64
	lastPennyUnix int64

	nextExpireScan time.Time
}
