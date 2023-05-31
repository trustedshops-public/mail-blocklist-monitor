package cmd

import (
	"com.trustedshops/mail-blocklist-monitor/pkg/blocklist-contributor"
	"com.trustedshops/mail-blocklist-monitor/pkg/dns"
	"com.trustedshops/mail-blocklist-monitor/pkg/exporter"
	ipprovider "com.trustedshops/mail-blocklist-monitor/pkg/ip-provider"
	"com.trustedshops/mail-blocklist-monitor/pkg/processing"
	"com.trustedshops/mail-blocklist-monitor/pkg/util"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

func init() {
	logLevel := util.GetEnvWithDefault("LOG_LEVEL", "INFO")
	parsedLevel, err := log.ParseLevel(logLevel)
	if err != nil {
		logLevel = ""
	}

	if logLevel != "" {
		log.SetLevel(parsedLevel)
	}

	if log.GetLevel() < log.DebugLevel {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func Execute(ctx context.Context) (string, error) {
	maxClassification := dns.LvlTextToByte(util.GetEnvWithDefault("MAX_CLASSIFICATION", "yellow"))
	if maxClassification == 0 {
		log.Fatal("Please specify a valid classification in env var MAX_CLASSIFICATION")
	}

	ips := ipprovider.LoadAllIps(ctx)

	allResults := processing.Check(ips, blocklist_contributor.AggregateBlocklist())

	negativeResults := make(dns.LookupResults, 0)
	for _, result := range allResults {
		if result.IsListed && result.ClassificationLevel.Numeric > maxClassification {
			negativeResults = append(negativeResults, result)
		}
	}

	exporter.ExecuteAll(ctx, &allResults, &negativeResults)

	return "", nil
}

func RunLambda() {
	lambda.Start(Execute)
}

func RunStandalone() {
	_, err := Execute(context.Background())
	if err != nil {
		log.Fatal("Something crashed the standalone app: ", err)
	}
}

func HandleUnknownEnv() {
	log.Fatal("Please specify standalone or aws_lambda as environment")
}
