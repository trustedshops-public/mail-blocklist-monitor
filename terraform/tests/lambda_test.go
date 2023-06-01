package tests

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/gruntwork-io/terratest/modules/random"
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

	if err != nil {
		t.Fatal(err)
	}
}

func TestTerraformModuleLambda_Cron(t *testing.T) {
	suffix := random.UniqueId()
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "lambda-cron",
		Vars: map[string]interface{}{
			"suffix": suffix,
		},
	})

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	terraform.InitAndApply(t, terraformOptions)
	checkSyncInvocation(t, lambda.NewFromConfig(cfg), "mail-blocklist-monitor-"+suffix)

	defer terraform.Destroy(t, terraformOptions)
}
