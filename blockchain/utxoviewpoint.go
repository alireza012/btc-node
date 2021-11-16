package blockchain

import (
	"btc-node/chaincfg/chainhash"
	"btc-node/wire"
)

type txoFlags uint8

type UtxoEntry struct {
	amount      int64
	pkScript    []byte
	blockHeight int32

	packedFlags txoFlags
}

type UtxoViewpoint struct {
	entries  map[wire.OutPoint]*UtxoEntry
	bestHash chainhash.Hash
}
