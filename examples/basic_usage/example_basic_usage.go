// Package main provides a basic example of using the Sentinel Provider
package main

// This file demonstrates basic usage of the Sentinel Provider in a Terraform configuration.
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
//    }
//
//    provider "sentinel" {
//      endpoint = "https://api.sentinel-project.io"
//      api_key  = var.sentinel_api_key
//    }
//
//    # Create a simple healthcare sentry
//    resource "sentinel_apollo" "hospital" {
//      name        = "main-hospital-sentry"
//      description = "Monitors critical healthcare infrastructure"
//      enabled     = true
//
//      tags = {
//        environment = "production"
//        team        = "security"
//      }
//    }
//
//    # Create an energy sector sentry
//    resource "sentinel_ra" "power_grid" {
//      name        = "power-grid-sentry"
//      description = "Monitors energy infrastructure"
//      enabled     = true
//
//      config = {
//        monitoring_mode = "continuous"
//      }
//
//      tags = {
//        critical = "true"
//      }
//    }
//
//    # Output sentry information
//    output "apollo_sentry_id" {
//      value = sentinel_apollo.hospital.id
//    }
//
//    output "ra_sentry_id" {
//      value = sentinel_ra.power_grid.id
//    }
//
// 2. Create a variables.tf file:
//
//    variable "sentinel_api_key" {
//      type      = string
//      sensitive = true
//      description = "API key for Sentinel Provider authentication"
//    }
//
// 3. Run Terraform commands:
//
//    terraform init
//    terraform plan
//    terraform apply
//
// This basic example shows:
// - Provider configuration
// - Creating sentry resources
// - Using tags for organization
// - Outputting resource information
