terraform
===

Terraform code sample and reusable modules

## Modules

- [lambda](./modules/lambda) - Configure and deploy lambda to run blocklist check
- [cron](./modules/cron) - Configure CloudWatch Event to run the lambda on a regular basis
- [pagerduty_integration](./modules/pagerduty_integration) - Setup PagerDuty Integration

## Examples

- [full](./examples/full) - Full example to copy&pasta'n'adapt
- [minimal](./examples/minimal) - PagerDuty details must be specified/created manually
