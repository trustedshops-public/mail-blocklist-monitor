package dns

import (
	"net"
	"testing"
)

func TestCheckListingLevel(t *testing.T) {
	testCases := []struct {
		ip            net.IP
		expectedLevel string
	}{
		{
			net.IPv4(127, 0, 0, 1),
			"white",
		},
		{
			net.IPv4(127, 0, 0, 2),
			"black",
		},
		{
			net.IPv4(127, 0, 0, 3),
			"yellow",
		},
		{
			net.IPv4(127, 0, 0, 4),
			"brown",
		},
		{
			net.IPv4(127, 0, 0, 5),
			"no_check",
		},
		{
			net.IPv4(127, 0, 0, 10),
			"",
		},
		{
			net.IPv4(127, 1, 1, 1),
			"",
		},
	}

	for _, tc := range testCases {
		_, lvl := CheckListingLevel(tc.ip)
		if tc.expectedLevel != lvl {
			t.Errorf("Expected text %s, but got '%s'", tc.expectedLevel, lvl)
		}
	}
}

func TestFindLevel(t *testing.T) {
	ips := []net.IP{
		net.IPv4(127, 0, 0, 3),
		net.IPv4(127, 0, 1, 1),
	}
	_, lvl := FindLevel(ips)
	if lvl != "yellow" {
		t.Fatal("Expected level yellow, but got ", lvl)
	}
}

func TestLvlTextToByte(t *testing.T) {
	testCases := []struct {
		level byte
		text  string
	}{
		{1, "white"},
		{0, "n/a"},
	}

	for _, tc := range testCases {
		txt := LvlTextToByte(tc.text)
		if txt != tc.level {
			t.Fatalf("Expected text level to be %d, but %d", tc.level, txt)
		}
	}
}
