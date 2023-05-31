output "pagerduty_integration_key" {
  value       = pagerduty_service_integration.this.integration_key
  description = "Integration Key for the new service integration"
}
