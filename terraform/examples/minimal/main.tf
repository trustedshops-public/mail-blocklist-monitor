terraform {
  required_providers {
    pagerduty = {
      source  = "PagerDuty/pagerduty"
      version = "2.16.2"
    }
  }
  required_version = ">= 1.3"
}

module "lambda" {
  source = "../../modules/lambda"

  kms_key_arn = "key-arn-here"
  environment_variables = {
    LOG_LEVEL                          = "INFO"
    MAX_CLASSIFICATION                 = "yellow"
    EXPORTER_PAGERDUTY_ENABLED         = "true"
    EXPORTER_PAGERDUTY_INTEGRATION_KEY = "integration-key-here"
    IP_PROVIDER_SES_REGIONS            = "eu-central-1"
  }
}

module "cron" {
  source               = "../../modules/cron"
  lambda_arn           = module.lambda.lambda_arn
  lambda_function_name = module.lambda.lambda_function_name
}
