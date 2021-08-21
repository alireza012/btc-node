package wire

import (
	"btc-node/chaincfg/chainhash"
	"time"
)

type BlockHeader struct {
	Version int32

	PrevBlock chainhash.Hash

	MerkleRoot chainhash.Hash

	Timestamp time.Time

	Bits uint32

	Nonce uint32
}
