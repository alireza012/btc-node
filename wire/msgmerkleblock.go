package wire

import "btc-node/chaincfg/chainhash"

type MsgMerkleBlock struct {
	Header       BlockHeader
	Transactions uint32
	Hashes       []*chainhash.Hash
	Flags        []byte
}
