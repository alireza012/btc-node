package wire

import "btc-node/chaincfg/chainhash"

type FilterType uint8

type MsgCFilter struct {
	FilterType FilterType
	BlockHash  chainhash.Hash
	Data       []byte
}
