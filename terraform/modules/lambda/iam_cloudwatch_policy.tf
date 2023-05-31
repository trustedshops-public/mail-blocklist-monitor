data "aws_region" "this" {}
data "aws_caller_identity" "this" {}
data "aws_partition" "this" {}

locals {
  log_group_arn  = format("arn:%s:logs:%s:%s:log-group:%s", data.aws_partition.this.id, data.aws_region.this.name, data.aws_caller_identity.this.id, local.log_group_name)
  log_group_name = "/aws/lambda/${var.function_name}"
}

data "aws_iam_policy_document" "cloudwatch" {
  # tfsec:ignore:aws-iam-no-policy-wildcards
  statement {
    sid = "AllowCloudWatchLogsUsage"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = [
      "${local.log_group_arn}:*"
    ]
  }
}

resource "aws_iam_role_policy" "cloudwatch" {
  name   = "cloudwatch"
  policy = data.aws_iam_policy_document.cloudwatch.json
  role   = aws_iam_role.this.name
}
