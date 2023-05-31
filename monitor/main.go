package main

import (
	"com.trustedshops/mail-blocklist-monitor/cmd"
	"com.trustedshops/mail-blocklist-monitor/pkg/util"
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
