package blocklist_contributor

import "strings"

type AdditionalBlockListContributor struct {
}

func (a AdditionalBlockListContributor) Contribute(blockLists []string) []string {
	lists := *getEnvConfig(a, "lists", "")
	if lists == "" {
		return blockLists
	}

	return append(blockLists, strings.Split(lists, ",")...)
}

func (a AdditionalBlockListContributor) Priority() int {
	return PriorityAdd
}

func (a AdditionalBlockListContributor) Name() string {
	return "additional"
}
