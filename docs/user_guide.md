# Sentinel Provider User Guide

Welcome to the Sentinel Provider user guide. This guide will help you get started with using the Sentinel Provider to manage AI Sentry resources for critical infrastructure protection.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Managing Sentries](#managing-sentries)
- [Examples](#examples)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)

---

## Introduction

The Sentinel Provider is a Terraform provider that allows you to manage AI Sentry resources as infrastructure as code. Each sentry is specialized to protect a specific critical infrastructure sector.

### Available Sentries

| Sentry   | Sector                                    | Resource Type          |
|----------|-------------------------------------------|------------------------|
| Apollo   | Healthcare                                | `sentinel_apollo`      |
| Ares     | Defense Industrial Base                   | `sentinel_ares`        |
| Athena   | Community-Based Governmental Organizations| `sentinel_athena`      |
| Demeter  | Food & Agriculture                        | `sentinel_demeter`     |
| Fenrir   | Information Technology                    | `sentinel_fenrir`      |
| Hermes   | Transportation                            | `sentinel_hermes`      |
| Jupiter  | Government                                | `sentinel_jupiter`     |
| Lir      | Water                                     | `sentinel_lir`         |
| Lugh     | Postal & Shipping                         | `sentinel_lugh`        |
| Mercury  | Commercial Facilities                     | `sentinel_mercury`     |
| Morrigan | Chemical                                  | `sentinel_morrigan`    |
| Osiris   | Emergency Services                        | `sentinel_osiris`      |
| Ptah     | Critical Manufacturing                    | `sentinel_ptah`        |
| Ra       | Energy                                    | `sentinel_ra`          |
| Shiva    | Nuclear Reactors, Materials, and Waste    | `sentinel_shiva`       |
| Sobek    | Dams                                      | `sentinel_sobek`       |
| Thoth    | Telecommunications                        | `sentinel_thoth`       |
| Tyche    | Banking & Finance                         | `sentinel_tyche`       |

---

## Installation

### Terraform 0.13+

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    sentinel = {
      source  = "cywf/sentinel-provider"
      version = "~> 0.1.0"
    }
  }
}
```

### Initialize Terraform

Run `terraform init` to download and install the provider:

```bash
terraform init
```

---

## Quick Start

Here's a simple example to get you started:

```hcl
# Configure the Sentinel Provider
provider "sentinel" {
  endpoint = "https://api.sentinel-project.io"
  api_key  = var.sentinel_api_key
}

# Create an Apollo sentry for healthcare monitoring
resource "sentinel_apollo" "hospital" {
  name        = "main-hospital-sentry"
  description = "Monitors critical healthcare infrastructure"
  enabled     = true

  config = {
    threat_level = "high"
    region       = "us-east-1"
  }

  tags = {
    environment = "production"
    team        = "security"
  }
}

# Output the sentry ID
output "sentry_id" {
  value = sentinel_apollo.hospital.id
}
```

Save this as `main.tf` and run:

```bash
terraform init
terraform plan
terraform apply
```

---

## Configuration

### Provider Configuration

The provider accepts the following configuration options:

```hcl
provider "sentinel" {
  endpoint = "https://api.sentinel-project.io"  # Optional
  api_key  = "your-api-key"                     # Optional, sensitive
}
```

### Using Environment Variables

Instead of hardcoding credentials, use environment variables:

```bash
export SENTINEL_ENDPOINT="https://api.sentinel-project.io"
export SENTINEL_API_KEY="your-api-key"
```

Then configure the provider without explicit credentials:

```hcl
provider "sentinel" {}
```

### Using Terraform Variables

For better security, use Terraform variables:

```hcl
variable "sentinel_api_key" {
  type      = string
  sensitive = true
}

provider "sentinel" {
  endpoint = "https://api.sentinel-project.io"
  api_key  = var.sentinel_api_key
}
```

---

## Managing Sentries

### Creating a Sentry

```hcl
resource "sentinel_ra" "energy_monitor" {
  name        = "national-grid-sentry"
  description = "Monitors the national energy grid"
  enabled     = true

  config = {
    monitoring_mode = "continuous"
    alert_threshold = "medium"
  }

  tags = {
    critical    = "true"
    environment = "production"
  }
}
```

### Updating a Sentry

Modify the resource in your Terraform configuration and apply:

```hcl
resource "sentinel_ra" "energy_monitor" {
  name        = "national-grid-sentry"
  description = "Updated description"
  enabled     = true

  config = {
    monitoring_mode = "continuous"
    alert_threshold = "high"  # Changed from medium
  }

  tags = {
    critical    = "true"
    environment = "production"
  }
}
```

```bash
terraform apply
```

### Disabling a Sentry

Set `enabled = false` to disable a sentry without destroying it:

```hcl
resource "sentinel_ra" "energy_monitor" {
  name    = "national-grid-sentry"
  enabled = false
}
```

### Deleting a Sentry

Remove the resource from your configuration and apply:

```bash
terraform destroy -target=sentinel_ra.energy_monitor
```

### Importing Existing Sentries

Import an existing sentry into Terraform:

```bash
terraform import sentinel_apollo.hospital apollo-hospital-1234567890
```

---

## Examples

### Example 1: Healthcare Infrastructure

```hcl
resource "sentinel_apollo" "hospital_network" {
  name        = "hospital-network-sentry"
  description = "Monitors hospital network infrastructure"
  enabled     = true

  config = {
    threat_level        = "high"
    monitoring_interval = "60"
    alert_email         = "security@hospital.example.com"
  }

  tags = {
    department  = "it-security"
    criticality = "high"
  }
}
```

### Example 2: Financial Services

```hcl
resource "sentinel_tyche" "banking_system" {
  name        = "core-banking-sentry"
  description = "Monitors core banking systems"
  enabled     = true

  config = {
    transaction_monitoring = "enabled"
    fraud_detection        = "advanced"
    compliance_mode        = "pci-dss"
  }

  tags = {
    compliance  = "required"
    environment = "production"
    team        = "fintech-security"
  }
}
```

### Example 3: Multi-Sector Deployment

```hcl
# Energy sector
resource "sentinel_ra" "power_grid" {
  name    = "power-grid-sentry"
  enabled = true
  
  config = {
    region = "northeast"
  }
}

# Water sector
resource "sentinel_lir" "water_system" {
  name    = "water-treatment-sentry"
  enabled = true
  
  config = {
    facility_type = "treatment"
  }
}

# Emergency services
resource "sentinel_osiris" "emergency_dispatch" {
  name    = "emergency-services-sentry"
  enabled = true
  
  config = {
    service_type = "dispatch"
  }
}
```

### Example 4: Conditional Sentry Deployment

```hcl
variable "enable_healthcare_monitoring" {
  type    = bool
  default = true
}

resource "sentinel_apollo" "hospital" {
  count = var.enable_healthcare_monitoring ? 1 : 0
  
  name        = "hospital-sentry"
  description = "Healthcare infrastructure monitoring"
  enabled     = true
}
```

---

## Best Practices

### 1. Use Descriptive Names

Choose names that clearly indicate the sentry's purpose:

```hcl
# Good
resource "sentinel_apollo" "cardiac_center_monitor" {
  name = "cardiac-center-infrastructure-sentry"
}

# Avoid
resource "sentinel_apollo" "sentry1" {
  name = "s1"
}
```

### 2. Tag Resources Appropriately

Use tags for organization, cost allocation, and filtering:

```hcl
resource "sentinel_ra" "energy_monitor" {
  name = "grid-monitor"
  
  tags = {
    environment = "production"
    team        = "infrastructure"
    cost_center = "security-ops"
    region      = "us-east"
  }
}
```

### 3. Use Modules for Reusability

Create modules for common sentry configurations:

```hcl
module "hospital_sentry" {
  source = "./modules/healthcare-sentry"
  
  name          = "hospital-a"
  threat_level  = "high"
  alert_email   = "security@hospital.example.com"
}
```

### 4. Secure Sensitive Data

Never commit API keys or sensitive configuration:

```hcl
# Use variables
variable "api_key" {
  type      = string
  sensitive = true
}

# Or environment variables
# export SENTINEL_API_KEY="..."
```

### 5. Implement State Locking

Use remote state with locking to prevent concurrent modifications:

```hcl
terraform {
  backend "s3" {
    bucket = "sentinel-terraform-state"
    key    = "sentries/terraform.tfstate"
    region = "us-east-1"
    
    dynamodb_table = "terraform-lock"
    encrypt        = true
  }
}
```

### 6. Regular Monitoring

Monitor sentry status and health:

```hcl
output "sentry_status" {
  value = {
    id     = sentinel_apollo.hospital.id
    status = sentinel_apollo.hospital.status
    sector = sentinel_apollo.hospital.sector
  }
}
```

---

## Troubleshooting

### Common Issues

#### 1. Provider Not Found

**Error:**
```
Error: Failed to query available provider packages
```

**Solution:**
```bash
terraform init -upgrade
```

#### 2. Authentication Failed

**Error:**
```
Error: Unable to authenticate with Sentinel API
```

**Solution:**
- Verify your API key is correct
- Check environment variables are set
- Ensure network connectivity to the API endpoint

#### 3. Resource Already Exists

**Error:**
```
Error: Resource already exists
```

**Solution:**
Import the existing resource:
```bash
terraform import sentinel_apollo.hospital apollo-hospital-1234567890
```

#### 4. Invalid Configuration

**Error:**
```
Error: Invalid configuration value
```

**Solution:**
- Check the API reference for valid configuration options
- Verify data types match schema requirements
- Ensure required fields are provided

### Debug Mode

Enable debug logging for troubleshooting:

```bash
TF_LOG=DEBUG terraform apply
```

### Getting Help

- Check the [API Reference](api_reference.md) for detailed resource documentation
- Review the [Developer Guide](developer_guide.md) for technical details
- Open an issue on [GitHub](https://github.com/cywf/sentinel-provider/issues)

---

## Next Steps

- Explore the [examples](../examples/) directory for more complex scenarios
- Read the [API Reference](api_reference.md) for complete resource documentation
- Check the [sentinel-project](https://github.com/cywf/sentinel-project) repository for integration examples

---

## Version History

- **v0.1.0**: Initial release with 18 sentry resources
