// Package main provides an advanced example of using the Sentinel Provider
package main

// This file demonstrates advanced usage patterns with the Sentinel Provider.
// To use this example:
//
// 1. Create a main.tf file with the following content:
//
//    terraform {
//      required_providers {
//        sentinel = {
//          source  = "cywf/sentinel-provider"
//          version = "~> 0.1.0"
//        }
//      }
//
//      backend "s3" {
//        bucket = "sentinel-terraform-state"
//        key    = "sentries/prod/terraform.tfstate"
//        region = "us-east-1"
//      }
//    }
//
//    provider "sentinel" {
//      endpoint = var.sentinel_endpoint
//      api_key  = var.sentinel_api_key
//    }
//
//    # Local values for reusable configuration
//    locals {
//      common_tags = {
//        managed_by  = "terraform"
//        project     = "sentinel"
//        environment = var.environment
//      }
//
//      critical_sectors = ["healthcare", "energy", "finance"]
//    }
//
//    # Healthcare sector sentry with advanced configuration
//    resource "sentinel_apollo" "hospital_network" {
//      for_each = toset(var.hospital_regions)
//
//      name        = "hospital-${each.key}-sentry"
//      description = "Monitors healthcare infrastructure in ${each.key}"
//      enabled     = true
//
//      config = {
//        threat_level        = "high"
//        monitoring_interval = "60"
//        alert_email         = "security-${each.key}@hospital.example.com"
//        region              = each.key
//      }
//
//      tags = merge(local.common_tags, {
//        sector      = "healthcare"
//        region      = each.key
//        criticality = "high"
//      })
//    }
//
//    # Financial sector sentry with compliance configuration
//    resource "sentinel_tyche" "banking_systems" {
//      count = var.enable_financial_monitoring ? 1 : 0
//
//      name        = "core-banking-sentry"
//      description = "Monitors core banking systems with compliance"
//      enabled     = true
//
//      config = {
//        transaction_monitoring = "enabled"
//        fraud_detection        = "advanced"
//        compliance_mode        = "pci-dss"
//        audit_logging          = "enabled"
//      }
//
//      tags = merge(local.common_tags, {
//        sector      = "finance"
//        compliance  = "pci-dss"
//        criticality = "critical"
//      })
//    }
//
//    # Energy sector sentry with dynamic configuration
//    resource "sentinel_ra" "power_grids" {
//      for_each = var.energy_facilities
//
//      name        = "energy-${each.key}-sentry"
//      description = "Monitors ${each.value.type} in ${each.value.location}"
//      enabled     = each.value.monitoring_enabled
//
//      config = {
//        facility_type   = each.value.type
//        location        = each.value.location
//        monitoring_mode = each.value.mode
//        capacity        = tostring(each.value.capacity)
//      }
//
//      tags = merge(local.common_tags, {
//        sector       = "energy"
//        facility     = each.key
//        location     = each.value.location
//        critical     = each.value.is_critical ? "true" : "false"
//      })
//    }
//
//    # Multi-sector deployment for coordinated monitoring
//    module "critical_infrastructure" {
//      source = "./modules/critical-infrastructure"
//
//      for_each = toset(local.critical_sectors)
//
//      sector      = each.key
//      environment = var.environment
//      tags        = local.common_tags
//    }
//
//    # Data sources for existing resources
//    data "terraform_remote_state" "network" {
//      backend = "s3"
//      config = {
//        bucket = "sentinel-terraform-state"
//        key    = "network/terraform.tfstate"
//        region = "us-east-1"
//      }
//    }
//
//    # Outputs with complex data structures
//    output "sentry_inventory" {
//      description = "Complete inventory of deployed sentries"
//      value = {
//        healthcare = {
//          for k, v in sentinel_apollo.hospital_network : k => {
//            id     = v.id
//            status = v.status
//            region = k
//          }
//        }
//        energy = {
//          for k, v in sentinel_ra.power_grids : k => {
//            id       = v.id
//            status   = v.status
//            facility = k
//          }
//        }
//        finance = var.enable_financial_monitoring ? {
//          banking = {
//            id     = sentinel_tyche.banking_systems[0].id
//            status = sentinel_tyche.banking_systems[0].status
//          }
//        } : {}
//      }
//    }
//
//    output "monitoring_endpoints" {
//      description = "Monitoring endpoints for each sentry"
//      value = {
//        for resource_key, resource in merge(
//          { for k, v in sentinel_apollo.hospital_network : "apollo-${k}" => v },
//          { for k, v in sentinel_ra.power_grids : "ra-${k}" => v }
//        ) : resource_key => {
//          id     = resource.id
//          sector = resource.sector
//          status = resource.status
//        }
//      }
//    }
//
// 2. Create a variables.tf file:
//
//    variable "environment" {
//      type        = string
//      description = "Environment name (dev, staging, prod)"
//    }
//
//    variable "sentinel_endpoint" {
//      type        = string
//      description = "Sentinel API endpoint"
//    }
//
//    variable "sentinel_api_key" {
//      type        = string
//      sensitive   = true
//      description = "API key for Sentinel Provider"
//    }
//
//    variable "hospital_regions" {
//      type        = list(string)
//      description = "Regions where hospitals are located"
//      default     = ["east", "west", "central"]
//    }
//
//    variable "enable_financial_monitoring" {
//      type        = bool
//      description = "Enable financial sector monitoring"
//      default     = true
//    }
//
//    variable "energy_facilities" {
//      type = map(object({
//        type                = string
//        location            = string
//        monitoring_enabled  = bool
//        mode                = string
//        capacity            = number
//        is_critical         = bool
//      }))
//      description = "Configuration for energy facilities"
//      default = {
//        grid_alpha = {
//          type                = "power_grid"
//          location            = "northeast"
//          monitoring_enabled  = true
//          mode                = "continuous"
//          capacity            = 5000
//          is_critical         = true
//        }
//        grid_beta = {
//          type                = "power_grid"
//          location            = "southwest"
//          monitoring_enabled  = true
//          mode                = "periodic"
//          capacity            = 3000
//          is_critical         = false
//        }
//      }
//    }
//
// This advanced example demonstrates:
// - Remote state backend configuration
// - Dynamic resource creation with for_each
// - Conditional resource creation with count
// - Local values for DRY configuration
// - Complex configuration objects
// - Module usage
// - Advanced output structures
// - Data source integration
