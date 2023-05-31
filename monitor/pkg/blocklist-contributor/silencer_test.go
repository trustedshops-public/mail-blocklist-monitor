package blocklist_contributor

import (
	"os"
	"testing"
)

var silencer = SilencerBlocklistContributor{}

func TestSilencerBlocklistContributor_Contribute(t *testing.T) {
	_ = os.Setenv("BLOCKLIST_CONTRIBUTOR_SILENCER_LISTS", "b,c")
	testContribute(silencer, t, []string{"a", "b"}, []string{"a"})

	_ = os.Setenv("BLOCKLIST_CONTRIBUTOR_SILENCER_LISTS", "")
	testContribute(silencer, t, []string{"a", "b"}, []string{"a", "b"})
}

func TestSilencerBlocklistContributor_Priority(t *testing.T) {
	testPriority(silencer, t, PriorityRemove)
}

func TestSilencerBlocklistContributor_Name(t *testing.T) {
	testProviderName(silencer, t, "silencer")
}
