package exporter

import (
	"context"
	"fmt"
	"github.com/PagerDuty/go-pagerduty"
	log "github.com/sirupsen/logrus"
	"github.com/trustedshops-public/mail-blocklist-monitor/pkg/dns"
	"strconv"
)

func init() {
	registerExporter(PagerDutyExporter{})
}

func createPagerDutyLink(text string, href string) interface{} {
	return struct {
		Href string `json:"href"`
		Text string `json:"text"`
	}{
		href, text,
	}
}

func (p PagerDutyExporter) triggerPagerDutyIncident(description string, entries dns.LookupResults) error {
	integrationKey := *getEnvConfig(p, "integration_key", "")
	severity := *getEnvConfig(p, "severity", "warning")

	res, err := pagerduty.ManageEvent(pagerduty.V2Event{
		RoutingKey: integrationKey,
		Action:     "trigger",
		DedupKey:   description,
		Payload: &pagerduty.V2Payload{
			Summary:  description,
			Source:   *getEnvConfig(p, "source", "blocklist-monitor"),
			Severity: severity,
			Details: map[string]interface{}{
				"severity": severity,
				"entries":  entries,
			},
		},
	})

	if err != nil {
		return err
	}

	log.Debug(res.Message)

	return nil
}

type PagerDutyExporter struct {
}

func (p PagerDutyExporter) OnlyOnFindings() bool {
	return true
}

func (p PagerDutyExporter) Name() string {
	return "pagerduty"
}

func (p PagerDutyExporter) aggregateByIP(results dns.LookupResults) map[string]dns.LookupResults {
	aggregate := map[string]dns.LookupResults{}

	for _, res := range results {
		_, found := aggregate[res.IP]
		if !found {
			aggregate[res.IP] = []dns.LookupResult{}
		}

		aggregate[res.IP] = append(aggregate[res.IP], res)
	}

	return aggregate
}

// Export PagerDuty events
func (p PagerDutyExporter) Export(ctx context.Context, results *dns.LookupResults) error {
	groupByList, _ := strconv.ParseBool(*getEnvConfig(p, "group_by_list", "false"))

	// Two possibilities
	// Grouping
	// - One alert per IP
	// No Grouping
	// One alert per IP and block list
	var exporter func(results *dns.LookupResults) error
	if groupByList {
		exporter = p.exportGroupedByList
	} else {
		exporter = p.exportForEachFinding
	}

	return exporter(results)
}

func (p PagerDutyExporter) exportForEachFinding(results *dns.LookupResults) error {
	for _, result := range *results {
		msg := fmt.Sprintf("%s found on blocklist %s", result.IP, result.List)
		if err := p.triggerPagerDutyIncident(msg, dns.LookupResults{result}); err != nil {
			return err
		}
	}
	return nil
}

func (p PagerDutyExporter) exportGroupedByList(results *dns.LookupResults) error {
	aggregated := p.aggregateByIP(*results)

	for ip, aggregatedResults := range aggregated {
		var links []interface{}

		for _, result := range aggregatedResults {
			links = append(links, createPagerDutyLink(result.IP+" - "+result.List+" - "+result.Details, result.Details))
		}

		if err := p.triggerPagerDutyIncident(fmt.Sprintf("%s found on %d mail blocklists\n", ip, len(aggregatedResults)), aggregatedResults); err != nil {
			return err
		}
	}
	return nil
}
