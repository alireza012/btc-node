package wire

import "btc-node/chaincfg/chainhash"

type MsgCFCheckpt struct {
	FilterType    FilterType
	StopHash      chainhash.Hash
	FilterHeaders []*chainhash.Hash
}
