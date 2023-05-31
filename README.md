mail-blocklist-monitor
===
[![GitHub License](https://img.shields.io/badge/license-MIT-lightgrey.svg)](https://github.com/trustedshops-public/mail-blocklist-monitor/blob/main/LICENSE)
[![pre-commit](https://img.shields.io/badge/%E2%9A%93%20%20pre--commit-enabled-success)](https://pre-commit.com/)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/trustedshops-public/mail-blocklist-monitor/tree/main.svg?style=shield)](https://dl.circleci.com/status-badge/redirect/gh/trustedshops-public/mail-blocklist-monitor/tree/main)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

Monitor the reputation of your mail server IPs with ease.

## Features

- Monitor your server IP reputation on common block lists
- Specify the IPs with AWS SES or env variable
- Export findings to PagerDuty alerts and more
- Easily extendable plugin architecture
- Ready to use terraform modules for AWS

## Usage

### Terraform

```hcl
# Create lambda resources
module "monitor" {
  source = "github.com/trustedshops-public/mail-blocklist-monitor//terraform/modules/lambda?ref=<version>"

  environment_variables = {
    LOG_LEVEL                          = "INFO"
    MAX_CLASSIFICATION                 = "yellow"
    EXPORTER_PAGERDUTY_ENABLED         = "true"
    EXPORTER_PAGERDUTY_INTEGRATION_KEY = local.pagerduty_integration_key
    # just pull ips from ses in the example
    IP_PROVIDER_SES_REGIONS            = "eu-central-1,eu-west-1"
  }

  create_cloudwatch_alarm = true
}

# Execute lambda every day
module "cron" {
  source               = "github.com/trustedshops-public/mail-blocklist-monitor//terraform/modules/cron?ref=<version>"
  lambda_arn           = module.monitor.lambda_arn
  lambda_function_name = module.monitor.lambda_function_name
}
```

> See  [monitor/README.md](./monitor/README.md#configuration) for all configuration environment variables

### Standalone

1. Download the latest binary for your supported platform
   from [GitHub releases](https://github.com/trustedshops-public/mail-blocklist-monitor/releases)
2. Configure the monitor using environment variables (see [monitor/README.md](./monitor/README.md#configuration)) for
   more details
    1. Make sure `ENVIRONMENT` is set to `standalone`
3. Execute the binary with crontab, manually etc

## Motivation

There are some paid tools out there or even some board resources for some providers. But none satisfied our needs so we
built a small customizable solution to make monitoring easy and effortless. It integrates well with PagerDuty and
requires no maintenance.

## Components

This project consists of two parts:

- [monitor](monitor): Go application (to be run in AWS Lambda or standalone) with plugin architecture that checks again
  block lists, getting IPs from sources and exporting to targets (e. g. PagerDuty incidents).
- [terraform modules](terraform): Terraform modules to easily deploy the monitor using AWS Serverless ecosystem.

The monitor is built using a plugin architecture. So if you need additional sources, targets and exporters feel free to
create a PR or fork to customize to match your demands.

## Development

### Prerequisites

#### Required Tools

| Name       | Version |
|:-----------|:--------|
| Go         | 1.20+   |
| terraform  | 1.3+    |
| pre-commit | *       |
