package wire

import (
	"net"
	"time"
)

type NetAddress struct {
	Timestamp time.Time

	Services ServiceFlag

	IP net.IP

	Port uint16
}
