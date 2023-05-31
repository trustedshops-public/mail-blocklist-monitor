package dns

import "net"

// List of levels
var level = map[byte]string{
	1: "white",
	2: "black",
	3: "yellow",
	4: "brown",
	5: "no_check",
}

var _, reputationLevelsCIDR, _ = net.ParseCIDR("127.0.0.0/29")

// FindLevel tries to figure out the reputation level based on the given IPs
func FindLevel(ips []net.IP) (byte, string) {
	for _, ip := range ips {
		lvl, txt := CheckListingLevel(ip)
		if txt != "" {
			return lvl, txt
		}
	}
	return 0, ""
}

// CheckListingLevel returns the reputation level for a given ip,
// if it is not valid a empty string is returned
func CheckListingLevel(ip net.IP) (byte, string) {
	ip = ip.To16()
	if !reputationLevelsCIDR.Contains(ip) {
		return 0, ""
	}

	identifierBlock := ip[15]
	lvl, _ := level[identifierBlock]
	return identifierBlock, lvl
}

// LvlTextToByte converts a given level text to the numeric representation
func LvlTextToByte(text string) byte {
	for v, k := range level {
		if k == text {
			return v
		}
	}

	return 0
}
