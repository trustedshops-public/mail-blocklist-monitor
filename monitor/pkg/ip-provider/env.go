package ip_provider

import (
	"context"
	"fmt"
	"net"
	"strings"
)

type EnvIpProvider struct {
}

func (e EnvIpProvider) Name() string {
	return "env"
}

func (e EnvIpProvider) GetAvailableIps(ctx context.Context) ([]string, error) {
	rawIps := *getEnvConfig(e, "ips", "")
	ips := strings.Split(rawIps, ",")
	cleanedIps := make([]string, 0)
	for _, ip := range ips {
		if ip == "" {
			continue
		}

		ip := strings.TrimSpace(ip)

		if net.ParseIP(ip) == nil {
			return nil, fmt.Errorf("invalid ip %s", ip)
		}

		cleanedIps = append(cleanedIps, ip)
	}

	return cleanedIps, nil
}

func (e EnvIpProvider) ShouldFailSoft() bool {
	return false
}
