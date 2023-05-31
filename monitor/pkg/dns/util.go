package dns

import (
	"strings"
)

// ReverseIP shuffles ip address blocks from last to first
func ReverseIP(ip string) string {
	parts := strings.Split(ip, ".")
	reversed := make([]string, len(parts))

	for i, part := range parts {
		reversed[len(parts)-1-i] = part
	}

	return strings.Join(reversed, ".")
}
