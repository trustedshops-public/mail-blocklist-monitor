package processing

import (
	"github.com/gammazero/workerpool"
	log "github.com/sirupsen/logrus"
	"github.com/trustedshops-public/mail-blocklist-monitor/pkg/dns"
	"sort"
	"strings"
)

func ProcessEntryFunc(list string, ip string, resChan chan<- dns.LookupResult) func() {
	return func() {
		log.Debug("Add IP ", ip, " for blocklist ", list)
		resChan <- dns.LookUpBlockList(list, ip)
	}
}

func Check(ips []string, lists []string) dns.LookupResults {
	wp := workerpool.New(4)
	resChan := make(chan dns.LookupResult, len(ips)*len(lists))
	for _, list := range lists {
		for _, ip := range ips {
			wp.Submit(ProcessEntryFunc(list, ip, resChan))
		}
	}

	wp.StopWait()
	close(resChan)

	results := make(dns.LookupResults, 0)
	for r := range resChan {
		results = append(results, r)
	}

	sort.Slice(results, func(i, j int) bool {
		switch strings.Compare(results[i].List, results[j].List) {
		case -1:
			return true
		case 1:
			return false
		}

		return results[i].IP > results[j].IP
	})

	return results
}
