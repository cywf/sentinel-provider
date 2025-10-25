# Sentinel Provider API Reference

This document provides detailed API reference for the Sentinel Provider, a Terraform provider for managing AI Sentry resources for critical infrastructure protection.

## Table of Contents

- [Provider Configuration](#provider-configuration)
- [Resources](#resources)
  - [Common Resource Schema](#common-resource-schema)
  - [Individual Sentry Resources](#individual-sentry-resources)

---

## Provider Configuration

### Provider Block

```hcl
terraform {
  required_providers {
    sentinel = {
      source  = "cywf/sentinel-provider"
      version = "~> 0.1.0"
    }
  }
}

provider "sentinel" {
  endpoint = "https://api.sentinel-project.io"  # Optional
  api_key  = "your-api-key-here"                # Optional, sensitive
}
```

### Provider Arguments

| Argument   | Type   | Required | Description                                          |
|------------|--------|----------|------------------------------------------------------|
| `endpoint` | string | No       | Sentinel API endpoint URL. Can be set via `SENTINEL_ENDPOINT` environment variable |
| `api_key`  | string | No       | API key for authentication. Can be set via `SENTINEL_API_KEY` environment variable. This value is sensitive. |

---

## Resources

### Common Resource Schema

All sentry resources share a common schema with the following attributes:

#### Arguments

| Argument      | Type         | Required | Description                                                    |
|---------------|--------------|----------|----------------------------------------------------------------|
| `name`        | string       | Yes      | The name of the sentry instance                                |
| `description` | string       | No       | A description of the sentry instance and its purpose          |
| `enabled`     | bool         | No       | Whether the sentry is enabled and actively monitoring (default: `true`) |
| `config`      | map(string)  | No       | Configuration parameters specific to this sentry              |
| `tags`        | map(string)  | No       | A map of tags to assign to the sentry resource                |

#### Attributes

| Attribute      | Type   | Description                                                    |
|----------------|--------|----------------------------------------------------------------|
| `id`           | string | The unique identifier for this sentry resource                 |
| `sector`       | string | The critical infrastructure sector this sentry protects        |
| `status`       | string | The current operational status (e.g., active, inactive)        |
| `last_updated` | string | Timestamp of the last update to this resource (RFC3339 format) |

---

## Individual Sentry Resources

### sentinel_apollo

Manages an Apollo Sentry resource specialized for protecting the Healthcare sector.

**Protected Industries:**
- Hospitals and clinics
- Research and diagnostic labs
- Pharmaceutical companies
- Medical device manufacturers
- Health insurance providers

**Example:**

```hcl
resource "sentinel_apollo" "hospital_monitor" {
  name        = "main-hospital-sentry"
  description = "Monitors critical healthcare infrastructure"
  enabled     = true

  config = {
    threat_level    = "high"
    monitoring_zone = "east-region"
    alert_email     = "security@hospital.example.com"
  }

  tags = {
    environment = "production"
    sector      = "healthcare"
  }
}
```

---

### sentinel_ares

Manages an Ares Sentry resource specialized for protecting the Defense Industrial Base sector.

---

### sentinel_athena

Manages an Athena Sentry resource specialized for protecting Community-Based Governmental Organizations.

---

### sentinel_demeter

Manages a Demeter Sentry resource specialized for protecting the Food & Agriculture sector.

---

### sentinel_fenrir

Manages a Fenrir Sentry resource specialized for protecting the Information Technology sector.

---

### sentinel_hermes

Manages a Hermes Sentry resource specialized for protecting the Transportation sector.

---

### sentinel_jupiter

Manages a Jupiter Sentry resource specialized for protecting the Government sector.

---

### sentinel_lir

Manages a Lir Sentry resource specialized for protecting the Water sector.

---

### sentinel_lugh

Manages a Lugh Sentry resource specialized for protecting the Postal & Shipping sector.

---

### sentinel_mercury

Manages a Mercury Sentry resource specialized for protecting the Commercial Facilities sector.

---

### sentinel_morrigan

Manages a Morrigan Sentry resource specialized for protecting the Chemical sector.

---

### sentinel_osiris

Manages an Osiris Sentry resource specialized for protecting the Emergency Services sector.

---

### sentinel_ptah

Manages a Ptah Sentry resource specialized for protecting the Critical Manufacturing sector.

---

### sentinel_ra

Manages a Ra Sentry resource specialized for protecting the Energy sector.

**Example:**

```hcl
resource "sentinel_ra" "power_grid_monitor" {
  name        = "national-grid-sentry"
  description = "Monitors energy infrastructure and power grid"
  enabled     = true

  config = {
    grid_region     = "northeast"
    monitoring_mode = "continuous"
  }

  tags = {
    critical = "true"
    sector   = "energy"
  }
}
```

---

### sentinel_shiva

Manages a Shiva Sentry resource specialized for protecting the Nuclear Reactors, Materials, and Waste sector.

---

### sentinel_sobek

Manages a Sobek Sentry resource specialized for protecting the Dams sector.

---

### sentinel_thoth

Manages a Thoth Sentry resource specialized for protecting the Telecommunications sector.

---

### sentinel_tyche

Manages a Tyche Sentry resource specialized for protecting the Banking & Finance sector.

**Example:**

```hcl
resource "sentinel_tyche" "bank_monitor" {
  name        = "financial-system-sentry"
  description = "Monitors banking and financial services"
  enabled     = true

  config = {
    transaction_monitoring = "enabled"
    fraud_detection        = "advanced"
  }

  tags = {
    compliance = "pci-dss"
    sector     = "finance"
  }
}
```

---

## Import

All sentry resources support importing using their ID:

```bash
terraform import sentinel_apollo.example apollo-example-1234567890
```

---

## Best Practices

1. **Enable Monitoring**: Keep sentries enabled in production environments
2. **Use Tags**: Tag resources for easier management and cost allocation
3. **Configuration**: Use the `config` map for sentry-specific settings
4. **Naming**: Use descriptive names that indicate the sentry's purpose
5. **Security**: Store API keys securely using environment variables or secret management systems
