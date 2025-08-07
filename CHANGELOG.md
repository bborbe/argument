# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

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
