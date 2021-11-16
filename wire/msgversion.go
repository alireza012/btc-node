package wire

import "time"

type MsgVersion struct {
	ProtocolVersion int32

	Services ServiceFlag

	Timestamp time.Time

	AddrYou NetAddress

	AddrMe NetAddress

	Nonce uint64

	UserAgent string

	LastBlock int32

	DisableRelayTx bool
}
