provider "aws" {
  region = "eu-central-1"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
  }
  required_version = ">= 1.3"
}


module "lambda" {
  source = "../../modules/lambda"

  kms_key_arn = null
  environment_variables = {
    LOG_LEVEL                     = "INFO"
    MAX_CLASSIFICATION            = "yellow"
    EXPORTER_TABLE_STDOUT_ENABLED = "true"
    IP_PROVIDER_SES_REGIONS       = "eu-central-1"
  }

  function_name = "mail-blocklist-monitor-${var.suffix}"
  iam_role_name = "mail-blocklist-monitor-${var.suffix}"
}

module "cron" {
  source               = "../../modules/cron"
  lambda_arn           = module.lambda.lambda_arn
  lambda_function_name = module.lambda.lambda_function_name
}

variable "suffix" {
  type        = string
  description = "Suffix to append to resources"
}
