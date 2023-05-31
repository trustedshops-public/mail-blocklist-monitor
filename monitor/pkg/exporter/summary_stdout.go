package exporter

import (
	"com.trustedshops/mail-blocklist-monitor/pkg/dns"
	"context"
	log "github.com/sirupsen/logrus"
)

func init() {
	registerExporter(SummaryStdoutExporter{})
}

type SummaryStdoutExporter struct {
}

func (s SummaryStdoutExporter) OnlyOnFindings() bool {
	return false
}

func (s SummaryStdoutExporter) Name() string {
	return "summary_stdout"
}

func (s SummaryStdoutExporter) Export(ctx context.Context, results *dns.LookupResults) error {
	findings := 0
	for _, r := range *results {
		if r.IsListed {
			findings++
		}
	}
	log.WithField("queryCount", len(*results)).WithField("findings", findings).Info("Run completed")
	return nil
}
