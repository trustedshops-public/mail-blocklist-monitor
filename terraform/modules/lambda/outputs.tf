output "lambda_arn" {
  value       = aws_lambda_function.this.arn
  description = "ARN of the created lambda"
}

output "lambda_function_name" {
  value       = aws_lambda_function.this.function_name
  description = "Function Name of the created lambda"
}

output "iam_role_arn" {
  value       = aws_iam_role.this.arn
  description = "ARN of the created IAM role"
}

output "iam_role_name" {
  value       = aws_iam_role.this.name
  description = "Name of the created IAM role"
}

output "cloudwatch_alarm_failed_invocations_arn" {
  value       = try(aws_cloudwatch_metric_alarm.this[0].arn, null)
  description = "ARN of the cloudwatch alarm for failed invocations"
}

output "cloudwatch_alarm_failed_invocations_name" {
  value       = try(aws_cloudwatch_metric_alarm.this[0].alarm_name, null)
  description = "Name of the cloudwatch alarm for failed invocations"
}
