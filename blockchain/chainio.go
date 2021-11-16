package blockchain

type SpentTxOut struct {
	Amount int64

	PkScript []byte

	Height int32

	IsCoinBase bool
}

type SequenceLock struct {
	Seconds     int64
	BlockHeight int32
}
