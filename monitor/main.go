package main

import (
	"github.com/trustedshops-public/mail-blocklist-monitor/cmd"
	"github.com/trustedshops-public/mail-blocklist-monitor/pkg/util"
)

func main() {
	env := util.GetEnvWithDefault("ENVIRONMENT", "standalone")

	switch env {
	case "aws_lambda":
		cmd.RunLambda()

	case "standalone":
		cmd.RunStandalone()

	default:
		cmd.HandleUnknownEnv()
	}
}
