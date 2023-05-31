package exporter

import (
	"context"
	"fmt"
	"github.com/trustedshops-public/mail-blocklist-monitor/pkg/dns"
	"os"
	"text/tabwriter"
)

func init() {
	registerExporter(TableStdoutExporter{})
}

// TableStdoutExporter writes all results to stdout in tabular form
type TableStdoutExporter struct {
}

func (t TableStdoutExporter) OnlyOnFindings() bool {
	return false
}

func (t TableStdoutExporter) Name() string {
	return "table_stdout"
}

func (t TableStdoutExporter) boolToStr(in bool, trueVal string, falseVal string) string {
	if in {
		return trueVal
	} else {
		return falseVal
	}
}

// Export with TabWriter to stdout
func (t TableStdoutExporter) Export(ctx context.Context, results *dns.LookupResults) error {
	w := tabwriter.NewWriter(os.Stdout, 10, 2, 3, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintln(w, "LISTED\tCLASSIFICATION\tLIST\tIP\tDETAILS")

	for _, r := range *results {
		_, _ = fmt.Fprintln(w, t.boolToStr(r.IsListed, "yes", "no")+"\t"+r.ClassificationLevel.Text+"\t"+r.List+"\t"+r.IP+"\t"+r.Details+" \t")
	}

	return w.Flush()
}
