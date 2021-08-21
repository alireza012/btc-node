package wire

import "btc-node/chaincfg/chainhash"

type InvType uint32

type InvVect struct {
	Type InvType
	Hash chainhash.Hash
}
