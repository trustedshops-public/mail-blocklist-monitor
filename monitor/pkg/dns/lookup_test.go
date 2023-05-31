//go:build local
// +build local

package dns

import "testing"

func TestLookUp(t *testing.T) {
	result := LookUpBlockList("zen.spamhaus.org", "223.91.0.74")

	if !result.IsListed {
		t.Error("Expected IP to be listed on spam list")
	}

	if result.Details == "" {
		t.Error("Expected details to be present")
	}
}

func TestLookUpBlockList(t *testing.T) {
	listedResult := LookUpBlockList("zen.spamhaus.org", "127.0.0.2")
	if !listedResult.IsListed {
		t.Fatal("Expected IsListed to be true")
	}

	if listedResult.Details == "" {
		t.Fatal("Expected details to be not empty")
	}

	notListed := LookUpBlockList("zen.spamhaus.org", "-1")
	if notListed.IsListed {
		t.Fatal("Expected IsListed to be false")
	}
}
