package txscript

import (
	"btc-node/btcec"
	"btc-node/chaincfg/chainhash"
	"sync"
)

type sigCacheEntry struct {
	sig    *btcec.Signature
	pubkey *btcec.PublicKey
}

type SigCache struct {
	sync.RWMutex
	validSigs  map[chainhash.Hash]sigCacheEntry
	maxEntries uint
}
