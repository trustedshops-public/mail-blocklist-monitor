package dns

import "testing"

func TestReverseIP(t *testing.T) {
	testCases := []struct {
		ip         string
		reversedIp string
	}{
		{
			ip:         "127.0.0.1",
			reversedIp: "1.0.0.127",
		},
	}

	for _, tc := range testCases {
		result := ReverseIP(tc.ip)
		if tc.reversedIp != result {
			t.Errorf("Expected IP '%s' but got '%s'", tc.reversedIp, result)
		}
	}
}
