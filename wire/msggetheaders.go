package wire

import "btc-node/chaincfg/chainhash"

type MsgGetHeaders struct {
	ProtocolVersion    uint32
	BlockLocatorHashes []*chainhash.Hash
	HashStop           chainhash.Hash
}
