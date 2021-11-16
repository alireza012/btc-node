package blockchain

import "btc-node/wire"

const (
	MaxBlockWeight = 4000000

	MaxBlockBaseSize = 1000000

	MaxBlockSigOpsCost = 80000

	WitnessScaleFactor = 4

	MinTxOutputWeight = WitnessScaleFactor * wire.MinTxOutPayload

	MaxOutputsPerBlock = MaxBlockWeight / MinTxOutputWeight
)
