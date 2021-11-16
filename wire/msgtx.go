package wire

import "btc-node/chaincfg/chainhash"

const (
	TxVersion = 1

	MaxTxInSequenceNum uint32 = 0xffffffff

	MaxPrevOutIndex uint32 = 0xffffffff

	SequenceLockTimeDisabled = 1 << 31

	SequenceLockTimeIsSeconds = 1 << 22

	SequenceLockTimeMask = 0x0000ffff

	SequenceLockTimeGranularity = 9

	defaultTxInOutAlloc = 15

	minTxInPayload = 9 + chainhash.HashSize

	maxTxInPerMessage = (MaxMessagePayload / minTxInPayload) + 1

	MinTxOutPayload = 9

	maxTxOutPerMessage = (MaxMessagePayload / MinTxOutPayload) + 1

	minTxPayload = 10

	freeListMaxScriptSize = 512

	freeListMaxItems = 12500

	maxWitnessItemsPerInput = 500000

	maxWitnessItemSize = 11000
)

type OutPoint struct {
	Hash  chainhash.Hash
	Index uint32
}

type TxIn struct {
	PreviousOutPoint OutPoint
	SignatureScript  []byte
	Witness          TxWitness
	Sequence         uint32
}

type TxWitness [][]byte

type TxOut struct {
	Value    int64
	PkScript []byte
}

type MsgTx struct {
	Version  int32
	TxIn     []*TxIn
	TxOut    []*TxOut
	LockTime uint32
}
