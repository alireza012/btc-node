package wire

import "btc-node/chaincfg/chainhash"

type MsgGetCFCheckpt struct {
	FilterType FilterType
	StopHash   chainhash.Hash
}
