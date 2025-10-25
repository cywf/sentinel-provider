# Sentinel Provider Developer Guide

This guide provides information for developers who want to contribute to or extend the Sentinel Provider.

## Table of Contents

- [Getting Started](#getting-started)
- [Development Environment](#development-environment)
- [Project Structure](#project-structure)
- [Building the Provider](#building-the-provider)
- [Testing](#testing)
- [Adding a New Sentry Resource](#adding-a-new-sentry-resource)
- [Contributing](#contributing)

---

## Getting Started

### Prerequisites

- **Go 1.21+**: The provider is written in Go
- **Terraform 1.0+**: For testing the provider
- **Git**: For version control

### Clone the Repository

```bash
git clone https://github.com/cywf/sentinel-provider.git
cd sentinel-provider
```

---

## Development Environment

### Install Dependencies

```bash
go mod download
```

### Environment Variables

For development, you may want to set:

```bash
export SENTINEL_ENDPOINT="http://localhost:8080"
export SENTINEL_API_KEY="dev-key"
```

---

## Project Structure

```
sentinel-provider/
├── main.go                      # Provider entry point
├── provider/
│   ├── provider.go              # Provider implementation
│   └── provider_test.go         # Provider tests
├── internal/
│   └── resources/
│       ├── models.go            # Common data models
│       ├── schema.go            # Shared schema definitions
│       ├── resource_apollo.go   # Apollo sentry resource
│       ├── resource_ares.go     # Ares sentry resource
│       └── ...                  # Other sentry resources
├── examples/
│   ├── basic_usage/             # Basic usage examples
│   └── advanced_usage/          # Advanced usage examples
├── docs/
│   ├── api_reference.md         # API documentation
│   ├── developer_guide.md       # This file
│   └── user_guide.md            # User documentation
├── sentries/                    # Legacy sentry definitions
├── scripts/                     # Utility scripts
├── go.mod                       # Go module definition
└── go.sum                       # Go module checksums
```

---

## Building the Provider

### Development Build

Build the provider for local development:

```bash
go build -o sentinel-provider
```

### Install Locally for Terraform

To use the provider with Terraform locally:

```bash
# Build with version info
go build -ldflags="-X 'main.version=dev'" -o terraform-provider-sentinel-provider

# Create local plugin directory
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/cywf/sentinel-provider/0.1.0/linux_amd64/

# Copy binary to plugin directory
cp terraform-provider-sentinel-provider ~/.terraform.d/plugins/registry.terraform.io/cywf/sentinel-provider/0.1.0/linux_amd64/
```

### Production Build

For production releases:

```bash
go build -ldflags="-X 'main.version=0.1.0'" -o terraform-provider-sentinel-provider
```

---

## Testing

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Run Specific Tests

```bash
# Test specific package
go test ./provider

# Test specific function
go test -run TestProviderSchema ./provider
```

### Acceptance Tests

Acceptance tests run against a real or mock API:

```bash
TF_ACC=1 go test ./... -v -timeout 120m
```

---

## Adding a New Sentry Resource

If you need to add a new sentry resource (e.g., for a new critical infrastructure sector):

### 1. Create Resource File

Create a new file in `internal/resources/resource_<name>.go`:

```go
package resources

import (
"context"
"fmt"
"time"

"github.com/hashicorp/terraform-plugin-framework/path"
"github.com/hashicorp/terraform-plugin-framework/resource"
"github.com/hashicorp/terraform-plugin-framework/types"
"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
_ resource.Resource                = &NewSentryResource{}
_ resource.ResourceWithConfigure   = &NewSentryResource{}
_ resource.ResourceWithImportState = &NewSentryResource{}
)

// NewNewSentryResource is a helper function to simplify the provider implementation.
func NewNewSentryResource() resource.Resource {
return &NewSentryResource{}
}

// NewSentryResource is the resource implementation.
type NewSentryResource struct{}

// Metadata returns the resource type name.
func (r *NewSentryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
resp.TypeName = req.ProviderTypeName + "_newsentry"
}

// Schema defines the schema for the resource.
func (r *NewSentryResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
resp.Schema = GetCommonSentrySchema(
"New Sector",
"Manages a NewSentry Sentry resource for the New Sector.",
)
}

// Configure adds the provider configured client to the resource.
func (r *NewSentryResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
// Configure API client here
}

// Implement Create, Read, Update, Delete, and ImportState methods
// following the pattern from other sentry resources
```

### 2. Register Resource in Provider

Edit `provider/provider.go` and add the new resource to the `Resources()` method:

```go
func (p *SentinelProvider) Resources(ctx context.Context) []func() resource.Resource {
return []func() resource.Resource{
// ... existing resources ...
resources.NewNewSentryResource,
}
}
```

### 3. Add Tests

Create tests in `internal/resources/resource_<name>_test.go`

### 4. Update Documentation

Add documentation in `docs/api_reference.md`

---

## Code Style and Standards

### Go Formatting

Use `gofmt` to format code:

```bash
gofmt -w .
```

### Linting

Use `golangci-lint` for linting:

```bash
golangci-lint run
```

### Conventions

1. **Error Handling**: Always check and properly handle errors
2. **Logging**: Use `tflog` for structured logging
3. **Context**: Pass context through all function calls
4. **Naming**: Use clear, descriptive names for variables and functions
5. **Comments**: Document exported functions and types

---

## Architecture

### Provider Architecture

The Sentinel Provider follows the Terraform Plugin Framework architecture:

1. **main.go**: Entry point that starts the provider server
2. **provider/provider.go**: Provider implementation with schema and configuration
3. **internal/resources**: Individual resource implementations
4. **models.go**: Shared data models
5. **schema.go**: Common schema definitions

### Resource Lifecycle

Each resource implements the following operations:

- **Create**: Initialize a new sentry resource
- **Read**: Retrieve current state
- **Update**: Modify existing resource
- **Delete**: Remove resource
- **ImportState**: Import existing resources into Terraform state

### Data Flow

```
Terraform CLI
    ↓
Provider Server (main.go)
    ↓
Provider Implementation (provider.go)
    ↓
Resource Implementation (resource_*.go)
    ↓
API Client (future implementation)
    ↓
Sentinel API (backend service)
```

---

## Interoperability with sentinel-project

The Sentinel Provider is designed to work seamlessly with the sentinel-project repository:

### Terraform Integration

- **Provider Resources**: Map to Terraform modules in sentinel-project
- **Configuration**: Provider configuration can reference sentinel-project variables
- **State Management**: Terraform state tracks both provider resources and project infrastructure

### Example Integration

```hcl
# In sentinel-project terraform configuration
module "sentry_deployment" {
  source = "./terraform/global_modules/sentry"
  
  sentry_type = "apollo"
  sentry_name = sentinel_apollo.hospital.name
  sentry_id   = sentinel_apollo.hospital.id
}

# Provider resource
resource "sentinel_apollo" "hospital" {
  name        = "hospital-sentry-1"
  description = "Healthcare monitoring sentry"
  enabled     = true
}
```

---

## Debugging

### Enable Debug Logging

Run the provider with debug flag:

```bash
TF_LOG=DEBUG terraform apply
```

### Use Debugger

Run provider with debug support:

```bash
go run -debug main.go
```

Then attach a debugger like `dlv`:

```bash
dlv attach <pid>
```

---

## Release Process

1. Update version in `main.go`
2. Update CHANGELOG.md
3. Create git tag: `git tag v0.1.0`
4. Push tag: `git push origin v0.1.0`
5. Build release binaries
6. Publish to Terraform Registry

---

## Getting Help

- **Issues**: Open an issue on GitHub
- **Discussions**: Join the discussion board
- **Contributing**: See CONTRIBUTING.md

---

## Resources

- [Terraform Plugin Framework Documentation](https://developer.hashicorp.com/terraform/plugin/framework)
- [Terraform Plugin Development](https://developer.hashicorp.com/terraform/plugin)
- [Go Documentation](https://go.dev/doc/)
