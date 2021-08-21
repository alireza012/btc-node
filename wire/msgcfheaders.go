package wire

import "btc-node/chaincfg/chainhash"

type MsgCFHeaders struct {
	FilterType       FilterType
	StopHash         chainhash.Hash
	PrevFilterHeader chainhash.Hash
	FilterHashes     []*chainhash.Hash
}
