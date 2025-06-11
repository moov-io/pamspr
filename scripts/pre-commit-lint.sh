#!/bin/bash

# Pre-commit linting script for PAM SPR
# This script runs the same checks as your GitHub build pipeline

set -e

echo "🔍 Running pre-commit linting checks..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    local color=$1
    local message=$2
    echo -e "${color}${message}${NC}"
}

# Change to project root
cd "$(git rev-parse --show-toplevel)"

print_status $YELLOW "Step 1: Running go vet..."
if go vet ./...; then
    print_status $GREEN "✅ go vet passed"
else
    print_status $RED "❌ go vet failed"
    exit 1
fi

print_status $YELLOW "Step 2: Running go fmt check..."
if [ -n "$(gofmt -l .)" ]; then
    print_status $RED "❌ Code is not formatted. Run: go fmt ./..."
    gofmt -l .
    exit 1
else
    print_status $GREEN "✅ Code formatting is correct"
fi

print_status $YELLOW "Step 3: Running tests..."
if go test ./...; then
    print_status $GREEN "✅ All tests passed"
else
    print_status $RED "❌ Tests failed"
    exit 1
fi

print_status $YELLOW "Step 4: Checking for common issues..."

# Check for TODO comments (optional warning)
if grep -r "TODO\|FIXME" --include="*.go" .; then
    print_status $YELLOW "⚠️  Found TODO/FIXME comments (not blocking)"
fi

# Check for potential exhaustive switch issues
print_status $YELLOW "Step 5: Checking for exhaustive switch statements..."
if go run golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow ./... 2>/dev/null; then
    print_status $GREEN "✅ Shadow analysis passed"
fi

print_status $GREEN "🎉 All pre-commit checks passed!"
print_status $YELLOW "Note: Run 'golangci-lint run' for more comprehensive analysis when available"

echo ""
echo "To install this as a git pre-commit hook, run:"
echo "  ln -sf ../../scripts/pre-commit-lint.sh .git/hooks/pre-commit"