package tests

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkSyncInvocation(t *testing.T, lambdaClient *lambda.Client, functionName string) {
	payload, err := json.Marshal(struct{}{})
	assert.Nil(t, err)

	_, err = lambdaClient.Invoke(
		context.TODO(),
		&lambda.InvokeInput{
			FunctionName: aws.String(functionName),
			Payload:      payload,
		},
	)
	assert.Nil(t, err)
}

func TestTerraformModuleLambda_Cron(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "lambda-cron",
		Vars:         map[string]interface{}{},
	})

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	terraform.InitAndApply(t, terraformOptions)
	checkSyncInvocation(t, lambda.NewFromConfig(cfg), "mail-blocklist-monitor")

	defer terraform.Destroy(t, terraformOptions)
}
