package wire

type Alert struct {
	Version int32

	RelayUntil int64

	Expiration int64

	ID int32

	Cancel int32

	SetCancel []int32

	MinVer int32

	MaxVer int32

	SetSubVer []string

	Priority int32

	Comment string

	StatusBar string

	Reserved string
}

type MsgAlert struct {
	SerializedPayload []byte

	Signature []byte

	Payload *Alert
}
