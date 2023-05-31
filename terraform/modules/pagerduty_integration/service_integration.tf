resource "pagerduty_service_integration" "this" {
  vendor  = data.pagerduty_vendor.custom_event_transformer.id
  name    = var.integration_name
  service = var.service_id
}
