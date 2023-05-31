pagerduty-integration module
===

Create PagerDuty integration to be used by monitor to create alerts

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
#### Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3 |
| <a name="requirement_pagerduty"></a> [pagerduty](#requirement\_pagerduty) | 2.15.0 |

#### Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_integration_name"></a> [integration\_name](#input\_integration\_name) | Name to set for the integration | `string` | `"Mail-Blocklist-Notifications"` | no |
| <a name="input_service_id"></a> [service\_id](#input\_service\_id) | ID of the PagerDuty service | `string` | n/a | yes |

#### Outputs

| Name | Description |
|------|-------------|
| <a name="output_pagerduty_integration_key"></a> [pagerduty\_integration\_key](#output\_pagerduty\_integration\_key) | Integration Key for the new service integration |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
