package blocklist_contributor

import (
	"os"
	"testing"
)

var additional = AdditionalBlockListContributor{}

func TestAdditionalBlockListContributor_Contribute(t *testing.T) {
	testContribute(additional, t, []string{"a"}, []string{"a"})

	_ = os.Setenv("BLOCKLIST_CONTRIBUTOR_ADDITIONAL_LISTS", "c")
	testContribute(additional, t, []string{"a", "b"}, []string{"a", "b", "c"})
}

func TestAdditionalBlockListContributor_Name(t *testing.T) {
	testProviderName(additional, t, "additional")
}

func TestAdditionalBlockListContributor_Priority(t *testing.T) {
	testPriority(additional, t, PriorityAdd)
}
