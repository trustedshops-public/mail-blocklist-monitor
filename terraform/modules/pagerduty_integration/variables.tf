variable "service_id" {
  description = "ID of the PagerDuty service"
  type        = string
}

variable "integration_name" {
  description = "Name to set for the integration"
  default     = "Mail-Blocklist-Notifications"
  type        = string
}
