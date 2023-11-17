terraform {
  required_providers {
    pagerduty = {
      source  = "PagerDuty/pagerduty"
      version = "3.1.1"
    }
  }
  required_version = ">= 1.3"
}

data "pagerduty_vendor" "custom_event_transformer" {
  name = "Custom Event Transformer"
}
