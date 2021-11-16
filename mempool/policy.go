package mempool

import "github.com/btcsuite/btcutil"

const (
	maxStandardP2SHSigOps = 15

	maxStandardTxWeight = 400000

	maxStandardSigScriptSize = 1650

	DefaultMinRelayTxFee = btcutil.Amount(1000)

	maxStandardMultiSigKeys = 3
)
