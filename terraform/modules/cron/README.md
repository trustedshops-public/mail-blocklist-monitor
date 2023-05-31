cron module
===

Create EventBridge cron event to run the monitor in a specific interval

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
#### Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

#### Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_lambda_arn"></a> [lambda\_arn](#input\_lambda\_arn) | ARN of the created lambda | `string` | n/a | yes |
| <a name="input_lambda_function_name"></a> [lambda\_function\_name](#input\_lambda\_function\_name) | Function name of the created lambda | `string` | n/a | yes |
| <a name="input_rule_name"></a> [rule\_name](#input\_rule\_name) | Name of the rule | `string` | `"mail-blocklist-monitor-cron"` | no |
| <a name="input_rule_schedule"></a> [rule\_schedule](#input\_rule\_schedule) | Value to pass to aws\_cloudwatch\_event\_rule#schedule\_expression | `string` | `"rate(1 day)"` | no |

#### Outputs

No outputs.
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
