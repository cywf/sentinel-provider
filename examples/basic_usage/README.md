# Basic Usage Example

This directory contains a basic example of using the Sentinel Provider to manage AI Sentry resources.

## What This Example Demonstrates

- Provider configuration with API authentication
- Creating multiple sentry resources across different sectors
- Using variables for configuration
- Tagging resources for organization
- Outputting resource information

## Prerequisites

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- Sentinel API credentials

## Usage

1. **Set your API key** as an environment variable:

```bash
export TF_VAR_sentinel_api_key="your-api-key-here"
```

Or create a `terraform.tfvars` file:

```hcl
sentinel_api_key = "your-api-key-here"
environment      = "development"
```

2. **Initialize Terraform**:

```bash
terraform init
```

3. **Review the plan**:

```bash
terraform plan
```

4. **Apply the configuration**:

```bash
terraform apply
```

5. **View outputs**:

```bash
terraform output
```

6. **Clean up** (when done):

```bash
terraform destroy
```

## Resources Created

This example creates three sentry resources:

1. **Apollo Sentry** - Healthcare sector monitoring
2. **Ra Sentry** - Energy sector monitoring
3. **Tyche Sentry** - Financial sector monitoring

## Customization

You can customize the example by:

- Changing the `environment` variable
- Modifying sentry configurations in the `config` blocks
- Adding or removing sentries
- Adjusting tags for better organization

## Next Steps

- Check the [advanced_usage](../advanced_usage/) example for more complex scenarios
- Read the [API Reference](../../docs/api_reference.md) for all available resources
- Explore the [User Guide](../../docs/user_guide.md) for best practices
