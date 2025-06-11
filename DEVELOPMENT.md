# PAM SPR Development Guide

## Local Linting Setup

This project includes local linting tools to match the GitHub build pipeline and prevent CI/CD failures.

### Quick Start

```bash
# Run pre-commit checks manually
./scripts/pre-commit-lint.sh

# Install git pre-commit hook (runs automatically on commit)
ln -sf ../../scripts/pre-commit-lint.sh .git/hooks/pre-commit
```

### Available Commands

#### Original Moov.io Commands
```bash
# Run comprehensive linting (downloads moov-io lint script)
make check
```

#### Local Development Commands
```bash
# Core testing
go test ./...              # Run all tests
go test -v ./...           # Run tests with verbose output
go test -cover ./...       # Run tests with coverage

# Code quality
go vet ./...               # Run go vet
go fmt ./...               # Format code
./scripts/pre-commit-lint.sh  # Run pre-commit checks

# Build
go build -o bin/pamspr cmd/pamspr/main.go  # Build CLI
go install ./cmd/pamspr    # Install CLI to GOPATH/bin
```

### Pre-Commit Hook

The pre-commit hook runs automatically when you commit and includes:

1. **go vet** - Static analysis for potential issues
2. **go fmt check** - Code formatting validation
3. **go test** - Full test suite execution
4. **TODO/FIXME detection** - Warns about remaining TODOs (non-blocking)
5. **Shadow analysis** - Checks for variable shadowing issues

### Linting Configuration

The project includes `.golangci.yml` with the same linters used in GitHub:

- **exhaustive** - Ensures all enum values are handled in switch statements
- **gosec** - Security vulnerability detection
- **misspell** - Spelling error detection
- **wastedassign** - Detects wasted variable assignments
- **errcheck** - Ensures error return values are checked
- **govet** - Standard Go static analysis
- **staticcheck** - Advanced static analysis
- **unused** - Detects unused code
- **gosimple** - Suggests code simplifications
- **ineffassign** - Detects ineffectual assignments

### Troubleshooting

#### golangci-lint Issues
If you encounter Go version compatibility issues with golangci-lint:

1. The pre-commit script will automatically fall back to basic checks
2. All essential validations still run (vet, fmt, tests)
3. The GitHub pipeline will catch any advanced linting issues

#### Pre-Commit Hook Not Running
```bash
# Verify hook is installed
ls -la .git/hooks/pre-commit

# Reinstall if needed
ln -sf ../../scripts/pre-commit-lint.sh .git/hooks/pre-commit
chmod +x scripts/pre-commit-lint.sh
```

#### Bypassing Hooks (Emergency)
```bash
# Skip pre-commit hook if needed (not recommended)
git commit --no-verify -m "emergency commit"
```

### Best Practices

1. **Always run tests before committing**:
   ```bash
   go test ./...
   ```

2. **Format code before committing**:
   ```bash
   go fmt ./...
   ```

3. **Test the pre-commit script manually**:
   ```bash
   ./scripts/pre-commit-lint.sh
   ```

4. **Keep the CI/CD pipeline green** - the local tools match GitHub exactly

### Integration with GitHub Pipeline

The local tools are designed to catch the same issues as the GitHub build pipeline:

- **Local**: `./scripts/pre-commit-lint.sh` 
- **GitHub**: Uses golangci-lint with exhaustive, gosec, misspell, etc.

This ensures you catch issues locally before they cause CI/CD failures.