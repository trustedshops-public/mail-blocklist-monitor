terraform {
  required_providers {
    pagerduty = {
      source  = "PagerDuty/pagerduty"
      version = "3.16.0"
    }
  }
  required_version = ">= 1.3"
}


module "pagerduty_integration" {
  source     = "../../modules/pagerduty_integration"
  service_id = "my-service-id"
}
