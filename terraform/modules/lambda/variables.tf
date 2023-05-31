variable "function_name" {
  description = "Name of the lambda function"
  type        = string
  default     = "mail-blocklist-monitor"
}

variable "environment_variables" {
  description = "Map with environment variables to configure the lambda"
  type        = map(string)
  default = {
    LOG_LEVEL = "warn"
  }
}

variable "iam_role_name" {
  description = "Name for the IAM role for lambda execution"
  type        = string
  default     = "mail-blocklist-monitor"
}

variable "run_make_build" {
  description = "Run make build on every apply. By default the prebuilt binary is used, which is suficient for most use cases."
  type        = bool
  default     = false
}

variable "create_cloudwatch_alarm" {
  description = "Create cloudwatch metric alarm for failed invocations"
  type        = bool
  default     = false
}

variable "cloudwatch_alarm_name" {
  description = "Custom cloudwatch alarm name"
  type        = string
  default     = null
}

variable "cloudwatch_alarm_severity" {
  description = "Severity of alarm"
  type        = string
  default     = "WARN"
}

variable "kms_key_arn" {
  description = "KMS Key to use for lambda and CloudWatch log group"
  type        = string
}
