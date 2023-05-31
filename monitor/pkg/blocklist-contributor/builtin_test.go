package blocklist_contributor

import (
	"strings"
	"testing"
)

var builtin = BuiltinBlocklistContributor{}

func TestDefaultBlocklists(t *testing.T) {
	for _, entry := range DefaultBlocklists {
		if strings.HasSuffix(entry, ".") || strings.HasPrefix(entry, ".") {
			t.Fatalf("Entry %s can not end or start with a dot", entry)
		}
	}
}

func TestBuiltinBlocklistContributor_Priority(t *testing.T) {
	testPriority(builtin, t, PriorityAdd)
}

func TestBuiltinBlocklistContributor_Contribute(t *testing.T) {
	testContribute(builtin, t, []string{}, DefaultBlocklists)
}

func TestBuiltinBlocklistContributor_Name(t *testing.T) {
	testProviderName(builtin, t, "builtin")
}
