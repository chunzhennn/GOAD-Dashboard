# GOAD-Dashboard

A dashboard for displaying GOAD lab status

## Backend

- Get the status of lab instances from Proxmox VE
- Get the status of client connections from pfSense
- Send requests of restoring lab instance snapshots to Proxmox VE
- Log the time of the last reset request

### Configurations

Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| PROXMOX_URL | Proxmox VE API URL (e.g., https://proxmox.example.com:8006) | Yes | - |
| PROXMOX_USERNAME | Proxmox VE API username | Yes | - |
| PROXMOX_REALM | Proxmox VE authentication realm (e.g., pam, pve) | Yes | - |
| PROXMOX_API_TOKEN_NAME | Proxmox VE API token name | Yes | - |
| PROXMOX_API_TOKEN | Proxmox VE API token value | Yes | - |
| PFSENSE_URL | pfSense API URL | Yes | - |
| PFSENSE_USERNAME | pfSense API username | Yes | - |
| PFSENSE_PASSWORD | pfSense API password | Yes | - |
| ENABLE_SWAGGER | Enable Swagger UI documentation (set to "1" to enable) | No | - |


