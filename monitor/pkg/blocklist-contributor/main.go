package blocklist_contributor

import (
	"com.trustedshops/mail-blocklist-monitor/pkg/util"
	"fmt"
	"sort"
	"strings"
)

const PriorityAdd = 0
const PriorityRemove = 10

// BlocklistContributor defines a component that modifies the block list and returns the modified version
type BlocklistContributor interface {
	// Contribute does the logic on the given blocklist and returns a new copy
	Contribute(blockLists []string) []string
	// Priority specifies when the contributor be called
	Priority() int
	// Name for the contributor
	Name() string
}

var _contributor []BlocklistContributor

func registerProvider(i BlocklistContributor) {
	_contributor = append(_contributor, i)
}

// getEnvConfig loads the env variable in format BLOCKLIST_CONTRIBUTOR_<<PROVIDERNAME>>_<<KEY>>
func getEnvConfig(b BlocklistContributor, key string, defaultVal string) *string {
	val := util.GetEnvWithDefault(strings.ToUpper(fmt.Sprintf("BLOCKLIST_CONTRIBUTOR_%s_%s", b.Name(), key)), defaultVal)
	return &val
}

// Deduplicate string slice
func Deduplicate(s []string) []string {
	var result []string
	seen := make(map[string]bool)
	for _, val := range s {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = true
		}
	}
	return result
}

// AggregateBlocklist returns a list with all blacklist contributors called
func AggregateBlocklist() []string {
	sort.Slice(_contributor, func(i, j int) bool {
		return _contributor[i].Priority() < _contributor[j].Priority()
	})
	blockList := make([]string, 0)

	for _, c := range _contributor {
		blockList = c.Contribute(blockList)
	}

	return Deduplicate(blockList)
}
