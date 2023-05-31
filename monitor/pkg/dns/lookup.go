package dns

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
)

// LookupResult for a given IP and list
type LookupResult struct {
	// IP that has been checked
	IP string
	// IsListed states if a record has been found for the ip and the given list
	IsListed bool
	// Details about the listing
	Details string
	// List dns name
	List string
	// LookupResult contains the returned Ips for the query
	LookupResult []net.IP
	// ClassificationLevel contains the classification of the entry
	ClassificationLevel LookupResultClassification
}

// LookupResultClassification specifies the class for an lookup entry
type LookupResultClassification struct {
	// Text is human readable
	Text string
	// Numeric is parsed
	Numeric byte
}

// LookupResults represents multiple lookup results
type LookupResults []LookupResult

// LookUpBlockList checks a single list for the given IP
func LookUpBlockList(list string, ip string) LookupResult {
	reversedIp := ReverseIP(ip)
	entry := fmt.Sprintf("%s.%s", reversedIp, list)

	lookupResult, err := resolver.LookupIP(context.Background(), "ip4", entry)
	isListed := true
	details := ""

	if err != nil {
		log.Debug("Failed to lookup record for ", entry, ", setting as not listed")
		isListed = false
	} else {
		txtResult, err := net.LookupTXT(entry)
		if err == nil {
			details = strings.Join(txtResult, "\n")
		} else {
			log.Warning("No TXT record found for ", entry, ", ignoring.")
		}
	}

	lvl, txt := FindLevel(lookupResult)

	return LookupResult{
		List:                list,
		IP:                  ip,
		IsListed:            isListed,
		Details:             details,
		LookupResult:        lookupResult,
		ClassificationLevel: LookupResultClassification{txt, lvl},
	}
}
