package wire

import "io"

const MaxMessagePayload = (1024 * 1024 * 32) // 32MB

type MessageEncoding uint32

type Message interface {
	BtcDecode(io.Reader, uint32, MessageEncoding) error
	BtcEncode(io.Writer, uint32, MessageEncoding) error
	Command() string
	MaxPayloadLength(uint32) uint32
}
