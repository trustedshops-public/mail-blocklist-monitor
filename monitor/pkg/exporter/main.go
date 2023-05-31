package exporter

import (
	"com.trustedshops/mail-blocklist-monitor/pkg/dns"
	"com.trustedshops/mail-blocklist-monitor/pkg/util"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

// Exporter provides the base functionality every exporter must support
type Exporter interface {
	// OnlyOnFindings specifies if the exporter should only be called for matches
	// if you return true the exporter will be called for all results, including negative ones
	OnlyOnFindings() bool

	// Name to identify the provider and used for env variable to check if the provider is enabled in format:
	// EXPORTER_<<NAME>>_ENABLED
	Name() string

	// Export the given lookup results as you like
	Export(ctx context.Context, results *dns.LookupResults) error
}

var _exporter []Exporter

func registerExporter(e Exporter) {
	_exporter = append(_exporter, e)
}

func GetExporter() []Exporter {
	return _exporter
}

func IsExporterActive(e Exporter) bool {
	result := getEnvConfig(e, "enabled", "false")
	switch *result {
	case "true":
		return true

	case "false":
		return false
	}

	return false
}

// getEnvConfig loads the env variable in format EXPORTER_<<EXPORTERNAME>>_<<KEY>>
func getEnvConfig(e Exporter, key string, defaultVal string) *string {
	val := util.GetEnvWithDefault(strings.ToUpper(fmt.Sprintf("EXPORTER_%s_%s", e.Name(), key)), defaultVal)
	return &val
}

func ExecuteAll(ctx context.Context, allResults *dns.LookupResults, negativeResults *dns.LookupResults) {
	for _, e := range GetExporter() {
		if !IsExporterActive(e) {
			log.Debug("Ignoring inactive exporter ", e.Name())
			continue
		}

		log.Debug("Execute exporter ", e.Name(), " OnlyOnFindings=", e.OnlyOnFindings())
		var err error

		if e.OnlyOnFindings() {
			err = e.Export(ctx, negativeResults)
		} else {
			err = e.Export(ctx, allResults)
		}

		if err != nil {
			log.Error("Error while calling exporter ", reflect.TypeOf(e).Name(), " ", err.Error())
		}
	}
}
