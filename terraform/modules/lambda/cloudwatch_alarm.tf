resource "aws_cloudwatch_metric_alarm" "this" {
  count = var.create_cloudwatch_alarm ? 1 : 0

  alarm_name          = var.cloudwatch_alarm_name == null ? "lambda-${aws_lambda_function.this.function_name}-errrors" : var.cloudwatch_alarm_name
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = 1
  metric_name         = "Errors"
  namespace           = "AWS/Lambda"
  statistic           = "Sum"
  threshold           = "0"
  period              = 120
  treat_missing_data  = "notBreaching"
  alarm_description   = "[${var.cloudwatch_alarm_severity}] Mail blocklist checker has failed invocations"
  dimensions = {
    FunctionName = aws_lambda_function.this.function_name
  }
}
