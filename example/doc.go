// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package main provides a comprehensive example demonstrating the features
// of github.com/bborbe/argument/v2 library for command-line argument parsing.
//
// This example showcases:
//   - Basic types (string, int, float64, bool) with defaults
//   - Pointer types for optional values
//   - Slice types with custom separators
//   - Custom types with validation via HasValidation interface
//   - Custom types with parsing via encoding.TextUnmarshaler
//   - Time types (time.Duration, time.Time, libtime.Duration, etc.)
//   - Environment variable fallbacks
//   - Display control (hidden fields, length-only display)
//
// To run this example:
//
//	cd example
//	go run main.go -username=alice -port=8080
//
// Or with environment variables:
//
//	BROKER=kafka:9092 go run main.go
//
// To see all available flags:
//
//	go run main.go -help
package main
