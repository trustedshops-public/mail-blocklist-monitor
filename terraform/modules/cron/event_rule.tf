resource "aws_cloudwatch_event_rule" "this" {
  name                = var.rule_name
  description         = "Run blacklist-monitor lambda on schedule"
  schedule_expression = var.rule_schedule
}
