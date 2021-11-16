package wire

type BloomUpdateType uint8

type MsgFilterLoad struct {
	Filter    []byte
	HashFuncs uint32
	Tweak     uint32
	Flags     BloomUpdateType
}
