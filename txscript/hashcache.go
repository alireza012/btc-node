package txscript

import (
	"btc-node/chaincfg/chainhash"
	"sync"
)

type TxSigHashes struct {
	HashPrevOuts chainhash.Hash
	HashSequence chainhash.Hash
	HashOutputs  chainhash.Hash
}

type HashCache struct {
	sigHashes map[chainhash.Hash]*TxSigHashes

	sync.RWMutex
}
