# Example Terraform configuration for basic Sentinel Provider usage

terraform {
  required_version = ">= 1.0"
  
  required_providers {
    sentinel = {
      source  = "cywf/sentinel-provider"
      version = "~> 0.1.0"
    }
  }
}

# Configure the Sentinel Provider
provider "sentinel" {
  endpoint = "https://api.sentinel-project.io"
  api_key  = var.sentinel_api_key
}

# Variables
variable "sentinel_api_key" {
  type        = string
  sensitive   = true
  description = "API key for Sentinel Provider authentication"
}

variable "environment" {
  type        = string
  default     = "development"
  description = "Environment name"
}

# Create a healthcare monitoring sentry (Apollo)
resource "sentinel_apollo" "hospital" {
  name        = "main-hospital-sentry"
  description = "Monitors critical healthcare infrastructure"
  enabled     = true

  config = {
    threat_level = "high"
    region       = "us-east-1"
  }

  tags = {
    environment = var.environment
    team        = "security"
    sector      = "healthcare"
  }
}

# Create an energy monitoring sentry (Ra)
resource "sentinel_ra" "power_grid" {
  name        = "power-grid-sentry"
  description = "Monitors energy infrastructure"
  enabled     = true

  config = {
    monitoring_mode = "continuous"
    alert_threshold = "medium"
  }

  tags = {
    environment = var.environment
    critical    = "true"
    sector      = "energy"
  }
}

# Create a financial services sentry (Tyche)
resource "sentinel_tyche" "banking" {
  name        = "banking-system-sentry"
  description = "Monitors banking and financial services"
  enabled     = true

  config = {
    compliance_mode        = "pci-dss"
    transaction_monitoring = "enabled"
  }

  tags = {
    environment = var.environment
    compliance  = "required"
    sector      = "finance"
  }
}

# Outputs
output "apollo_sentry_id" {
  description = "ID of the Apollo healthcare sentry"
  value       = sentinel_apollo.hospital.id
}

output "apollo_sentry_status" {
  description = "Status of the Apollo healthcare sentry"
  value       = sentinel_apollo.hospital.status
}

output "ra_sentry_id" {
  description = "ID of the Ra energy sentry"
  value       = sentinel_ra.power_grid.id
}

output "tyche_sentry_id" {
  description = "ID of the Tyche finance sentry"
  value       = sentinel_tyche.banking.id
}

output "all_sentries" {
  description = "Summary of all deployed sentries"
  value = {
    healthcare = {
      id     = sentinel_apollo.hospital.id
      name   = sentinel_apollo.hospital.name
      status = sentinel_apollo.hospital.status
      sector = sentinel_apollo.hospital.sector
    }
    energy = {
      id     = sentinel_ra.power_grid.id
      name   = sentinel_ra.power_grid.name
      status = sentinel_ra.power_grid.status
      sector = sentinel_ra.power_grid.sector
    }
    finance = {
      id     = sentinel_tyche.banking.id
      name   = sentinel_tyche.banking.name
      status = sentinel_tyche.banking.status
      sector = sentinel_tyche.banking.sector
    }
  }
}
