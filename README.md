# Argument

[![Go Version](https://img.shields.io/github/go-mod/go-version/bborbe/argument)](https://github.com/bborbe/argument)
[![Go Reference](https://pkg.go.dev/badge/github.com/bborbe/argument/v2.svg)](https://pkg.go.dev/github.com/bborbe/argument/v2)
[![License](https://img.shields.io/github/license/bborbe/argument)](https://github.com/bborbe/argument/blob/master/LICENSE)

A declarative Go library for parsing command-line arguments and environment variables into structs using struct tags. Perfect for building CLI applications with clean configuration management.

## Features

- 🏷️ **Declarative**: Use struct tags to define argument names, environment variables, and defaults
- 🔄 **Multiple Sources**: Supports command-line arguments, environment variables, and default values
- ⚡ **Zero Dependencies**: Minimal external dependencies for core functionality
- ✅ **Type Safe**: Supports all common Go types including pointers for optional values
- 🕐 **Extended Duration**: Enhanced `time.Duration` parsing with support for days and weeks
- 🧪 **Well Tested**: Comprehensive test suite with BDD-style tests

## Installation

```bash
go get github.com/bborbe/argument/v2
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/bborbe/argument/v2"
)

func main() {
    var config struct {
        Username string `arg:"username" env:"USERNAME" default:"guest"`
        Password string `arg:"password" env:"PASSWORD"`
        Port     int    `arg:"port" env:"PORT" default:"8080"`
        Debug    bool   `arg:"debug" env:"DEBUG"`
    }
    
    ctx := context.Background()
    if err := argument.Parse(ctx, &config); err != nil {
        log.Fatalf("Failed to parse arguments: %v", err)
    }
    
    fmt.Printf("Starting server on port %d for user %s\n", config.Port, config.Username)
}
```

## Usage Examples

### Basic Configuration

```go
type Config struct {
    Host     string `arg:"host" env:"HOST" default:"localhost"`
    Port     int    `arg:"port" env:"PORT" default:"8080"`
    LogLevel string `arg:"log-level" env:"LOG_LEVEL" default:"info"`
}

var config Config
err := argument.Parse(context.Background(), &config)
```

Run with: `./app -host=0.0.0.0 -port=9090 -log-level=debug`

### Optional Values with Pointers

```go
type DatabaseConfig struct {
    Host     string  `arg:"db-host" env:"DB_HOST" default:"localhost"`
    Port     int     `arg:"db-port" env:"DB_PORT" default:"5432"`
    Timeout  *int    `arg:"timeout" env:"DB_TIMEOUT"`        // Optional
    MaxConns *int    `arg:"max-conns" env:"DB_MAX_CONNS"`    // Optional
}
```

### Duration with Extended Parsing

```go
type ServerConfig struct {
    ReadTimeout  time.Duration `arg:"read-timeout" env:"READ_TIMEOUT" default:"30s"`
    WriteTimeout time.Duration `arg:"write-timeout" env:"WRITE_TIMEOUT" default:"1m"`
    IdleTimeout  time.Duration `arg:"idle-timeout" env:"IDLE_TIMEOUT" default:"2h"`
    Retention    time.Duration `arg:"retention" env:"RETENTION" default:"30d"`  // 30 days
    BackupFreq   time.Duration `arg:"backup-freq" env:"BACKUP_FREQ" default:"1w"` // 1 week
}
```

Supported duration units: `ns`, `us`, `ms`, `s`, `m`, `h`, `d` (days), `w` (weeks)

### Required Fields

```go
type APIConfig struct {
    APIKey   string `arg:"api-key" env:"API_KEY"`           // Required (no default)
    Endpoint string `arg:"endpoint" env:"ENDPOINT"`         // Required (no default)
    Region   string `arg:"region" env:"REGION" default:"us-east-1"`
}

// This will return an error if APIKey or Endpoint are not provided
err := argument.Parse(context.Background(), &config)
```

## Supported Types

- **Strings**: `string`
- **Integers**: `int`, `int32`, `int64`, `uint`, `uint64`
- **Floats**: `float64`
- **Booleans**: `bool`
- **Durations**: `time.Duration` (with extended parsing)
- **Pointers**: `*string`, `*int`, `*float64`, etc. (for optional values)

## Priority Order

Values are applied in the following priority order (higher priority overwrites lower):

1. **Default values** (from `default:` tag)
2. **Environment variables** (from `env:` tag)  
3. **Command-line arguments** (from `arg:` tag)

## API Reference

### Parse Functions

```go
// Parse parses arguments and environment variables (quiet mode)
func Parse(ctx context.Context, data interface{}) error

// ParseAndPrint parses and prints the final configuration values
func ParseAndPrint(ctx context.Context, data interface{}) error
```

### Validation

```go
// ValidateRequired checks that all required fields (no default value) are set
func ValidateRequired(ctx context.Context, data interface{}) error
```

## Command-Line Usage

Your application will automatically support standard Go flag syntax:

```bash
# Long form with equals
./app -host=localhost -port=8080

# Long form with space  
./app -host localhost -port 8080

# Boolean flags
./app -debug        # sets debug=true
./app -debug=false  # sets debug=false
```

## Environment Variables

Set environment variables to configure your application:

```bash
export HOST=0.0.0.0
export PORT=9090
export DEBUG=true
./app
```

## Error Handling

The library provides detailed error messages for common issues:

```go
err := argument.Parse(ctx, &config)
if err != nil {
    // Errors include context about what failed:
    // - Missing required fields
    // - Type conversion errors  
    // - Invalid duration formats
    log.Fatal(err)
}
```

## Testing

The library is thoroughly tested with BDD-style tests using Ginkgo and Gomega:

```bash
make test        # Run all tests
make precommit   # Run full development workflow
```

## Version 2 Changes

Version 2.3.0 introduced breaking changes for better library behavior:

- `Parse()` no longer prints configuration by default (quieter)
- New `ParseAndPrint()` function when you want to display parsed values
- Focus on library-like behavior vs CLI tool behavior

## Contributing

Contributions are welcome! This project follows standard Go conventions and includes:

- Comprehensive tests with Ginkgo/Gomega
- Code generation with `make generate`
- Linting with `make check`
- Formatting with `make format`

## License

BSD 3-Clause License - see [LICENSE](LICENSE) file for details.
