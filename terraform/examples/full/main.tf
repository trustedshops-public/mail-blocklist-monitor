terraform {
  required_providers {
    pagerduty = {
      source  = "PagerDuty/pagerduty"
      version = "3.30.9"
    }
  }
  required_version = ">= 1.3"
}

module "lambda" {
  source = "../../modules/lambda"

  kms_key_arn = "key-arn-goes-here"
  environment_variables = {
    LOG_LEVEL                          = "INFO"
    MAX_CLASSIFICATION                 = "yellow"
    EXPORTER_PAGERDUTY_ENABLED         = "true"
    EXPORTER_PAGERDUTY_INTEGRATION_KEY = module.pagerduty_integration.pagerduty_integration_key
    IP_PROVIDER_SES_REGIONS            = "eu-central-1,eu-west-1"
  }
}

module "cron" {
  source               = "../../modules/cron"
  lambda_arn           = module.lambda.lambda_arn
  lambda_function_name = module.lambda.lambda_function_name
}

data "pagerduty_service" "this" {
  name = "trutre-test"
}

module "pagerduty_integration" {
  source     = "../../modules/pagerduty_integration"
  service_id = data.pagerduty_service.this.id
}
