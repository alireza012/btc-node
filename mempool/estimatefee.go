package mempool

import (
	"btc-node/chaincfg/chainhash"
	"sync"
)

const (
	estimateFeeDepth = 25
)

type SatoshiPerByte float64

type observedTransaction struct {
	hash chainhash.Hash

	feeRate SatoshiPerByte

	observed int32

	mined int32
}

type registeredBlock struct {
	hash         chainhash.Hash
	transactions []*observedTransaction
}

type FeeEstimator struct {
	maxRollback uint32
	binSize     int32

	maxReplacements int32

	minRegisteredBlocks uint32

	lastKnownHeight int32

	numBlocksRegistered uint32

	mtx      sync.RWMutex
	observed map[chainhash.Hash]*observedTransaction
	bin      [estimateFeeDepth][]*observedTransaction

	cached []SatoshiPerByte

	dropped []*registeredBlock
}
