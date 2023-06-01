package tests

import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"testing"
)

func TestTerraformModulePagerDutyIntegration(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "pagerduty_integration",
		Vars:         map[string]interface{}{},
	})

	terraform.Init(t, terraformOptions)
	terraform.Validate(t, terraformOptions)
}
