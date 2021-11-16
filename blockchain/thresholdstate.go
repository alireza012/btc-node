package blockchain

import "btc-node/chaincfg/chainhash"

type ThresholdState byte

type thresholdStateCache struct {
	entries map[chainhash.Hash]ThresholdState
}
