# Initialize a new Go module
go mod init <module-name>

# Download dependencies (automatically updates go.sum)
go mod tidy

# Add a dependency explicitly
go get <module-path>@<version>

# Upgrade all dependencies to latest minor/patch
go get -u ./...

# Run a Go file or package
go run main.go
go run ./cmd/server

# Build a binary
go build             # builds current package
go build ./cmd/app   # builds a specific package

# Output binary with a custom name
go build -o myapp main.go

# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run a specific test function
go test -run <TestName>

# Run benchmarks
go test -bench=.

# Run tests and show code coverage
go test -cover ./...

# Format code
go fmt ./...

# Check for issues and unused imports
go vet ./...

# Print dependencies
go list -m all

# Upgrade the Go toolchain
go install golang.org/dl/go1.23@latest

# Remove build cache
go clean -cache

# Remove module download cache
go clean -modcache

# Remove all test binaries
go clean -testcache

# Run a simple REPL (with third-party tool)
go install github.com/x-motemen/gore/cmd/gore@latest

# Cross-compile for another OS/arch
GOOS=linux GOARCH=amd64 go build -o app-linux

