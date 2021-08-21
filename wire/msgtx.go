package wire

import "btc-node/chaincfg/chainhash"

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
