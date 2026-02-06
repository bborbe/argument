# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v2.12.4

- Update Go to 1.25.7
- Update github.com/bborbe/errors to v1.5.2
- Update github.com/bborbe/time to v1.22.0
- Update testing dependencies (ginkgo v2.28.1, gomega v1.39.1)
- Update various indirect dependencies

## v2.12.3

- update golang to 1.25.6
- update github.com/bborbe/errors to v1.5.1
- update github.com/bborbe/time to v1.21.1
- update golangci-lint to v1.64.8
- add CI workflow

## v2.12.2

- Fix nil pointer panic in ValidateHasValidation when validating fields with nil pointer values
- Add generic pointer nil check in ValidateRequired to handle all pointer types consistently
- Refactor ValidateRequired into helper function to reduce cyclomatic complexity
- Add comprehensive test coverage for nil pointer handling (required and optional fields)

## v2.12.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update bborbe dependencies (errors v1.5.0, time v1.21.0, collection v1.20.0, parse v1.9.1, run v1.8.3)
- Update testing dependencies (ginkgo v2.27.5, gomega v1.39.0)
- Update security tools (gosec v2.22.11, osv-scanner v2.3.1)
- Update indirect dependencies

## v2.12.0

- update go and deps

## v2.11.4

- Improve Makefile organization with .PHONY declarations for all targets
- Reorganize check target to include lint alongside other quality checks
- Remove TODO comment and enable lint in check target

## v2.11.3

- Update github.com/google/osv-scanner from v2.2.4 to v2.3.0
- Update github.com/incu6us/goimports-reviser from v3.10.0 to v3.11.0
- Update multiple indirect dependencies

## v2.11.2

- Update Go version from 1.25.3 to 1.25.4 in CI workflow
- Update github.com/shoenig/go-modtool from v0.4.0 to v0.5.0
- Update github.com/containerd/containerd from v1.7.27 to v1.7.29
- Update github.com/cyphar/filepath-securejoin from v0.5.0 to v0.6.0
- Update github.com/opencontainers/selinux from v1.12.0 to v1.13.0
- Add cyphar.com/go-pathrs v0.2.1 dependency

## v2.11.1

- Enhance documentation with comprehensive priority examples (arg > env > default)
- Add zero-value documentation for fields without defaults
- Improve golangci-lint configuration with additional linters (gofmt, goimports, errname, bodyclose, asasalint, prealloc)
- Add io/ioutil deprecation rule to depguard
- Add nolint comments for existing technical debt (duplication and complexity)
- Integrate lint into precommit workflow
- Fix naming conventions (Url → URL, ApiKey → APIKey)
- Remove duplicate test file (example_required_test.go)
- Remove example/doc.go (documentation belongs in main package)
- Streamline README by removing Development and Contributing sections
- Improve example Makefile with priority demonstration comments

## v2.11.0

**BREAKING BEHAVIOR CHANGE**: Fixed argument precedence to match documentation and standard CLI tool behavior

- Fix precedence bug: command-line arguments now correctly override environment variables
- Previous (incorrect) behavior: env vars overrode args when both were provided
- New (correct) behavior: args > env > defaults (matches Docker, Kubernetes, most CLI tools)
- Add `argsToValuesExplicit()` internal function to filter only explicitly-set command-line flags
- Update `ParseOnly()` to use explicit args filtering and correct merge order
- Update test expectations to validate correct precedence behavior
- Upgrade test framework from Ginkgo v1 to Ginkgo v2 (all 11 test files)
- Add counterfeiter mock generation for `HasValidation` interface
- Add UTC timezone configuration to test suite to prevent timezone-dependent failures
- Enable race detection in test suite for concurrent safety validation
- Add `github.com/shoenig/go-modtool` for go.mod formatting
- Update README priority order documentation to match correct behavior
- No API changes: all public function signatures remain identical
- Migration: verify configuration sources if code relied on env-override-args behavior

## v2.10.0

- Add `HasValidation` interface for custom field validation with `Validate(context.Context) error` method
- Add `ParseOnly()` function for validation-free parsing (enables custom validation workflows)
- Add `ValidateHasValidation()` for manual validation of types implementing `HasValidation`
- Update validation chain: `Parse()` now calls `ParseOnly()` → `ValidateRequired()` → `ValidateHasValidation()`
- Add comprehensive test coverage for `HasValidation` interface across struct, field, and slice validation
- Add example package documentation demonstrating all library features

## v2.9.0

- Add slice support for required field validation
- Improve array/slice output formatting in Print with count and comma-separated values
- Add comprehensive test coverage for required validation across all supported types
- Add example_required_test.go with 23 real-world validation examples

## v2.8.0

- Add support for slice types implementing TextUnmarshaler on the slice itself (e.g., kafka.Brokers)
- Enhanced Fill function to handle encoding.TextMarshaler types for proper JSON round-tripping
- Reordered type checks to prioritize TextUnmarshaler before generic slice handling
- Add MarshalText support requirement for slice types using TextUnmarshaler
- Consolidate test files following one implementation file per test file pattern
- Add comprehensive Context-based test structure for TextUnmarshaler functionality

## v2.7.0

- Add support for encoding.TextUnmarshaler interface for custom type parsing
- Allow types to define custom unmarshaling logic via UnmarshalText method
- Support TextUnmarshaler for both single values and slice elements
- Works with command-line arguments, environment variables, and default values
- Add comprehensive test suite with TestBroker and TestURL examples
- Update documentation with TextUnmarshaler examples

## v2.6.0

- Add support for slice types with comma-separated values
- Support []string, []int, []int64, []uint, []uint64, []float64, []bool
- Support custom type slices (e.g., []Username where type Username string)
- Add configurable separator via `separator:` struct tag (default: ",")
- Automatically trim whitespace from slice elements
- Support slice parsing from both command-line arguments and environment variables
- Add slice support to DefaultValues function for default tag handling
- Add comprehensive test suite for slice functionality
- Enhance example with demonstrations of all slice types and separators

## v2.5.1

- Fix error wrapping pattern: replace errors.Wrapf with errors.Wrap where no format placeholders are used
- Add package-level documentation in doc.go
- Add comprehensive documentation for ParseAndPrint function
- Improve Fill function documentation with detailed parameter descriptions
- Upgrade to Go 1.25.3
- Enhance README with CI and Go Report Card badges
- Expand Development section in README with all available Makefile commands
- Update Contributing section with step-by-step instructions
- Update copyright years from 2019 to 2025

## v2.5.0

- Add support for time.Time and *time.Time types with RFC3339 format parsing
- Add support for *time.Duration pointer type for optional duration values
- Add support for libtime.Duration and *libtime.Duration with extended format (weeks, days)
- Add support for libtime.DateTime and *libtime.DateTime for timestamp parsing
- Add support for libtime.Date and *libtime.Date for date-only parsing
- Add support for libtime.UnixTime and *libtime.UnixTime for Unix timestamp parsing
- Add comprehensive documentation for Parse, ParseArgs, and ParseEnv functions
- Add 44 new tests covering all time types and pointer variants
- Fix silent error swallowing in default value parsing - now returns descriptive errors
- Improve test coverage from 86.9% to include all time type scenarios

## v2.4.3

- Upgrade to Go 1.25.2
- Add golangci-lint configuration for enhanced code quality checks
- Add security scanning tools (osv-scanner, gosec, trivy) to build pipeline
- Enhance CI workflow with Trivy installation step
- Improve Makefile with additional quality and security check targets
- Update dependencies and development tools

## v2.4.2

- Improve code formatting with better function parameter layout
- Add golines dependency for automated line length formatting
- Update build tools and dependencies for enhanced code quality
- Enhance test coverage and code structure improvements

## v2.4.1

- Fix integration test timeouts for CI environment compatibility
- Improve gexec test robustness with 30-second timeout configuration

## v2.4.0

- Add support for custom types with underlying primitive types (e.g., `type Username string`)
- Support custom types in all parsing modes: command-line arguments, environment variables, and defaults
- Add custom type support to ValidateRequired function for proper required field validation
- Enhance test coverage with 42 new tests across all parsing methods and validation
- Add integration tests using gexec to validate end-to-end executable behavior
- Update README documentation with custom types examples and usage patterns
- Improve code coverage from 82.2% to 85.4%

## v2.3.2

- Enhance README with comprehensive documentation, examples, and badges
- Add additional test coverage for error paths and edge cases
- Add GitHub Actions workflows for CI/CD automation
- Improve test completeness reaching 98.4% coverage

## v2.3.1

- add tests
- go mod update

## v2.3.0

- remove print by default
- add ParseAndPrint
- go mod update

## v2.2.0

- remove vendor
- go mod update

## v2.1.5

- go mod update

## v2.1.4

- go mod update

## v2.1.3

- update license
- go mod update

## v2.1.2

- go mod update

## v2.1.1

- go mod update

## v2.1.0

- fix libtime.ParseDuration
- go mod update

## v2.0.5

- go mod update

## v2.0.4

- go mod update

## v2.0.3

- Fix parseDuration

## v2.0.2

- Fix parseDuration

## v2.0.1

- fix v2

## v2.0.0

- add context to all methods
- allow parse duration with unit d=day and w=week

## v1.3.2

- fix int32

## v1.3.1

- fix int32

## v1.3.0

- go mod update
- allow int32 type

## v1.2.4

- add vulncheck
- go mod update

## v1.2.3

- fix parse of empty *float64

## v1.2.2

- update deps

## v1.2.1

- fix print of *float64

## v1.2.0

- change go version to 1.20
- add *float64

## v1.1.0

- Go mod upgrade

## v1.0.1

- Fix validation of multiple fields

## v1.0.0

- Initial Version
