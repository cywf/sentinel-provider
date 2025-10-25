# Contributing to Sentinel Provider

Thank you for your interest in contributing to the Sentinel Provider! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Process](#development-process)
- [Pull Request Process](#pull-request-process)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Documentation](#documentation)

## Code of Conduct

This project adheres to a code of conduct that we expect all contributors to follow:

- Be respectful and inclusive
- Welcome newcomers and help them get started
- Focus on what's best for the community
- Show empathy towards other community members

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Terraform 1.0 or higher
- Git
- A GitHub account

### Fork and Clone

1. Fork the repository on GitHub
2. Clone your fork locally:

```bash
git clone https://github.com/YOUR_USERNAME/sentinel-provider.git
cd sentinel-provider
```

3. Add the upstream repository:

```bash
git remote add upstream https://github.com/cywf/sentinel-provider.git
```

### Set Up Development Environment

1. Install dependencies:

```bash
go mod download
```

2. Build the provider:

```bash
go build
```

3. Run tests:

```bash
go test ./...
```

## Development Process

### Branching Strategy

- `main` - Production-ready code
- `develop` - Integration branch for features
- `feature/*` - New features
- `bugfix/*` - Bug fixes
- `docs/*` - Documentation updates

### Making Changes

1. Create a new branch from `main`:

```bash
git checkout -b feature/your-feature-name
```

2. Make your changes following our [coding standards](#coding-standards)

3. Write or update tests for your changes

4. Update documentation as needed

5. Commit your changes with clear, descriptive messages:

```bash
git commit -m "Add feature: description of what you added"
```

### Commit Message Guidelines

Follow the conventional commits specification:

- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation changes
- `test:` - Adding or updating tests
- `refactor:` - Code refactoring
- `style:` - Code style changes (formatting, etc.)
- `chore:` - Maintenance tasks

Examples:
```
feat: add support for custom monitoring intervals
fix: correct sector assignment for Mercury sentry
docs: update API reference with new parameters
test: add integration tests for Apollo resource
```

## Pull Request Process

1. **Update your branch** with the latest from upstream:

```bash
git fetch upstream
git rebase upstream/main
```

2. **Push your changes** to your fork:

```bash
git push origin feature/your-feature-name
```

3. **Create a Pull Request** on GitHub with:
   - Clear title describing the change
   - Description of what changed and why
   - Reference to any related issues
   - Screenshots if applicable

4. **Address review feedback** by making additional commits

5. **Wait for approval** from maintainers

### PR Review Criteria

Your PR will be reviewed for:

- Code quality and adherence to standards
- Test coverage
- Documentation completeness
- Backward compatibility
- Performance implications

## Coding Standards

### Go Code Style

Follow standard Go conventions:

- Use `gofmt` for formatting
- Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use meaningful variable and function names
- Add comments for exported functions and types

### Provider-Specific Guidelines

1. **Resource Implementation**:
   - Use the common schema from `internal/resources/schema.go`
   - Implement all CRUD operations
   - Add proper error handling
   - Use structured logging with `tflog`

2. **Error Handling**:
```go
if err != nil {
    resp.Diagnostics.AddError(
        "Error Creating Resource",
        fmt.Sprintf("Could not create resource: %s", err),
    )
    return
}
```

3. **Logging**:
```go
tflog.Info(ctx, "Creating resource", map[string]interface{}{
    "id": plan.ID.ValueString(),
})
```

### Testing Standards

- Write unit tests for all new functions
- Maintain or increase code coverage
- Use table-driven tests where appropriate
- Mock external dependencies

Example test:
```go
func TestResourceCreate(t *testing.T) {
    tests := []struct {
        name    string
        input   SentryResourceModel
        wantErr bool
    }{
        {
            name: "valid resource",
            input: SentryResourceModel{
                Name: types.StringValue("test"),
            },
            wantErr: false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package
go test ./provider

# Run with race detection
go test -race ./...
```

### Acceptance Tests

For acceptance tests that interact with a real API:

```bash
TF_ACC=1 go test ./... -v -timeout 120m
```

## Documentation

### Documentation Requirements

All contributions should include:

1. **Code Comments**: Document exported functions and types
2. **API Reference**: Update `docs/api_reference.md` for new resources
3. **User Guide**: Add examples to `docs/user_guide.md`
4. **Developer Guide**: Update `docs/developer_guide.md` for architectural changes

### Documentation Style

- Use clear, concise language
- Include code examples
- Provide context and rationale
- Link to related documentation

## Adding New Resources

To add a new sentry resource:

1. Create `internal/resources/resource_<name>.go`
2. Implement all required methods
3. Register in `provider/provider.go`
4. Add tests in `internal/resources/resource_<name>_test.go`
5. Document in `docs/api_reference.md`
6. Add example usage

See [Developer Guide](docs/developer_guide.md) for detailed instructions.

## Questions?

If you have questions about contributing:

- Open an issue with the `question` label
- Start a discussion on GitHub Discussions
- Reach out to the maintainers

## Recognition

Contributors will be recognized in:

- Release notes
- CONTRIBUTORS file (coming soon)
- Project README

Thank you for contributing to Sentinel Provider!
