package wire

import "btc-node/chaincfg/chainhash"

type RejectCode uint8

type MsgReject struct {
	Cmd string

	Code RejectCode

	Reason string

	Hash chainhash.Hash
}
