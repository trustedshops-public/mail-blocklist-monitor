package dns

import (
	"context"
	"net"
	"time"
)

func resolverDial(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{
		Timeout: time.Second,
	}
	return d.DialContext(ctx, "udp", "1.1.1.1:53")
}

var resolver = &net.Resolver{
	PreferGo: true,
	Dial:     resolverDial,
}
