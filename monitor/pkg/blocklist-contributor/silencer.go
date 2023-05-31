package blocklist_contributor

import "strings"

type SilencerBlocklistContributor struct {
}

func (s SilencerBlocklistContributor) Contribute(blockLists []string) []string {
	lists := *getEnvConfig(s, "lists", "")
	if lists == "" {
		return blockLists
	}

	toRemove := strings.Split(lists, ",")
	cleaned := make([]string, 0)
	remove := false

	for _, l := range blockLists {
		remove = false

		for _, r := range toRemove {
			if r == l {
				remove = true
				break
			}
		}

		if !remove {
			cleaned = append(cleaned, l)
		}
	}

	return cleaned
}

func (s SilencerBlocklistContributor) Priority() int {
	return PriorityRemove
}

func (s SilencerBlocklistContributor) Name() string {
	return "silencer"
}
