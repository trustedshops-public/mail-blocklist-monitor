lambda module
===

Create lambda function with proper permissions in AWS.

Requirements for execution:
- make
- golang

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
#### Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3 |
| <a name="requirement_archive"></a> [archive](#requirement\_archive) | ~> 2.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |
| <a name="requirement_null"></a> [null](#requirement\_null) | ~> 3.0 |

#### Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cloudwatch_alarm_name"></a> [cloudwatch\_alarm\_name](#input\_cloudwatch\_alarm\_name) | Custom cloudwatch alarm name | `string` | `null` | no |
| <a name="input_cloudwatch_alarm_severity"></a> [cloudwatch\_alarm\_severity](#input\_cloudwatch\_alarm\_severity) | Severity of alarm | `string` | `"WARN"` | no |
| <a name="input_create_cloudwatch_alarm"></a> [create\_cloudwatch\_alarm](#input\_create\_cloudwatch\_alarm) | Create cloudwatch metric alarm for failed invocations | `bool` | `false` | no |
| <a name="input_environment_variables"></a> [environment\_variables](#input\_environment\_variables) | Map with environment variables to configure the lambda | `map(string)` | <pre>{<br>  "LOG_LEVEL": "warn"<br>}</pre> | no |
| <a name="input_function_name"></a> [function\_name](#input\_function\_name) | Name of the lambda function | `string` | `"mail-blocklist-monitor"` | no |
| <a name="input_iam_role_name"></a> [iam\_role\_name](#input\_iam\_role\_name) | Name for the IAM role for lambda execution | `string` | `"mail-blocklist-monitor"` | no |
| <a name="input_kms_key_arn"></a> [kms\_key\_arn](#input\_kms\_key\_arn) | KMS Key to use for lambda and CloudWatch log group | `string` | n/a | yes |
| <a name="input_run_make_build"></a> [run\_make\_build](#input\_run\_make\_build) | Run make build on every apply. By default the prebuilt binary is used, which is suficient for most use cases. | `bool` | `false` | no |

#### Outputs

| Name | Description |
|------|-------------|
| <a name="output_cloudwatch_alarm_failed_invocations_arn"></a> [cloudwatch\_alarm\_failed\_invocations\_arn](#output\_cloudwatch\_alarm\_failed\_invocations\_arn) | ARN of the cloudwatch alarm for failed invocations |
| <a name="output_cloudwatch_alarm_failed_invocations_name"></a> [cloudwatch\_alarm\_failed\_invocations\_name](#output\_cloudwatch\_alarm\_failed\_invocations\_name) | Name of the cloudwatch alarm for failed invocations |
| <a name="output_iam_role_arn"></a> [iam\_role\_arn](#output\_iam\_role\_arn) | ARN of the created IAM role |
| <a name="output_iam_role_name"></a> [iam\_role\_name](#output\_iam\_role\_name) | Name of the created IAM role |
| <a name="output_lambda_arn"></a> [lambda\_arn](#output\_lambda\_arn) | ARN of the created lambda |
| <a name="output_lambda_function_name"></a> [lambda\_function\_name](#output\_lambda\_function\_name) | Function Name of the created lambda |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
