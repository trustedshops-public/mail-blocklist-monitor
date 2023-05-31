variable "rule_name" {
  description = "Name of the rule"
  type        = string
  default     = "mail-blocklist-monitor-cron"
}

variable "rule_schedule" {
  description = "Value to pass to aws_cloudwatch_event_rule#schedule_expression"
  type        = string
  default     = "rate(1 day)"
}

variable "lambda_arn" {
  description = "ARN of the created lambda"
  type        = string
}

variable "lambda_function_name" {
  description = "Function name of the created lambda"
  type        = string
}
