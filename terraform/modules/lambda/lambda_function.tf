# tfsec:ignore:aws-lambda-enable-tracing
resource "aws_lambda_function" "this" {
  function_name    = var.function_name
  description      = "Monitor if the SES dedicated IPs in this account are listed on a blocklist and notify if it is the case."
  handler          = "monitor"
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 600
  role             = aws_iam_role.this.arn
  filename         = data.archive_file.this.output_path
  source_code_hash = data.archive_file.this.output_base64sha256

  environment {
    variables = merge(
      var.environment_variables, {
        ENVIRONMENT                     = "aws_lambda"
        EXPORTER_SUMMARY_STDOUT_ENABLED = "true"
      }
    )
  }

  tracing_config {
    mode = "PassThrough"
  }
}

data "archive_file" "this" {
  output_path = ".terraform/build/lambda.zip"
  source_dir  = "${path.module}/../../../monitor/dist"
  type        = "zip"
  depends_on  = [null_resource.this]
}
