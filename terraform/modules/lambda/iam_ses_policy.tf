data "aws_iam_policy_document" "ses" {
  statement {
    effect = "Allow"
    actions = [
      "ses:GetDedicatedIps"
    ]
    #tfsec:ignore:aws-iam-no-policy-wildcards
    resources = [
      "*"
    ]
  }
}

resource "aws_iam_role_policy" "ses" {
  name   = "ses"
  policy = data.aws_iam_policy_document.ses.json
  role   = aws_iam_role.this.name
}
