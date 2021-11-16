package wire

import "btc-node/chaincfg/chainhash"

type MsgGetBlocks struct {
	ProtocolVersion    uint32
	BlockLocatorHashes []*chainhash.Hash
	HashStop           chainhash.Hash
}
