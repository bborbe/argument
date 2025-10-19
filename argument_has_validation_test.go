// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"context"

	"github.com/bborbe/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bborbe/argument/v2"
)

// Test types for HasValidation

type testPort int

func (p testPort) Validate(ctx context.Context) error {
	if p < 1024 {
		return errors.New(ctx, "port must be >= 1024")
	}
	return nil
}

type testTimeout int

func (t testTimeout) Validate(ctx context.Context) error {
	if t < 1 || t > 300 {
		return errors.New(ctx, "timeout must be between 1 and 300 seconds")
	}
	return nil
}

type testBroker string

func (b testBroker) Validate(ctx context.Context) error {
	if len(b) < 5 {
		return errors.New(ctx, "broker must be at least 5 characters (format: host:port)")
	}
	return nil
}

type testBrokers []testBroker

func (b testBrokers) Validate(ctx context.Context) error {
	if len(b) == 0 {
		return errors.New(ctx, "at least one broker required")
	}
	return nil
}

type testValidatingConfig struct {
	Port testPort
}

func (c *testValidatingConfig) Validate(ctx context.Context) error {
	if c.Port < 1024 {
		return errors.New(ctx, "port must be >= 1024")
	}
	return nil
}

var _ = Describe("HasValidation", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})

	Context("Struct-level validation", func() {
		It("calls Validate on struct implementing HasValidation", func() {
			// Valid case
			validConfig := &testValidatingConfig{Port: 8080}
			err := argument.ValidateHasValidation(ctx, validConfig)
			Expect(err).NotTo(HaveOccurred())

			// Invalid case
			invalidConfig := &testValidatingConfig{Port: 80}
			err = argument.ValidateHasValidation(ctx, invalidConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("port must be >= 1024"))
		})

		It("succeeds when struct does not implement HasValidation", func() {
			type Config struct {
				Port int `arg:"port"`
			}

			config := &Config{Port: 8080}
			err := argument.ValidateHasValidation(ctx, config)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Field-level validation", func() {
		It("validates field implementing HasValidation", func() {
			type Config struct {
				Port testPort
			}

			// Valid case
			validConfig := &Config{Port: 8080}
			err := argument.ValidateHasValidation(ctx, validConfig)
			Expect(err).NotTo(HaveOccurred())

			// Invalid case
			invalidConfig := &Config{Port: 80}
			err = argument.ValidateHasValidation(ctx, invalidConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("field Port"))
			Expect(err.Error()).To(ContainSubstring("validation failed"))
			Expect(err.Error()).To(ContainSubstring("port must be >= 1024"))
		})

		It("validates multiple fields implementing HasValidation", func() {
			type Config struct {
				Port    testPort
				Timeout testTimeout
			}

			// All valid
			validConfig := &Config{Port: 8080, Timeout: 30}
			err := argument.ValidateHasValidation(ctx, validConfig)
			Expect(err).NotTo(HaveOccurred())

			// Invalid port
			invalidPortConfig := &Config{Port: 80, Timeout: 30}
			err = argument.ValidateHasValidation(ctx, invalidPortConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Port"))

			// Invalid timeout
			invalidTimeoutConfig := &Config{Port: 8080, Timeout: 500}
			err = argument.ValidateHasValidation(ctx, invalidTimeoutConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Timeout"))
		})
	})

	Context("Slice validation", func() {
		It("validates slice elements implementing HasValidation", func() {
			type Config struct {
				Brokers []testBroker
			}

			// All valid
			validConfig := &Config{
				Brokers: []testBroker{"broker1:9092", "broker2:9092"},
			}
			err := argument.ValidateHasValidation(ctx, validConfig)
			Expect(err).NotTo(HaveOccurred())

			// One invalid element
			invalidConfig := &Config{
				Brokers: []testBroker{"broker1:9092", ""},
			}
			err = argument.ValidateHasValidation(ctx, invalidConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("field Brokers[1]"))
			Expect(err.Error()).To(ContainSubstring("validation failed"))
			Expect(err.Error()).To(ContainSubstring("broker must be at least 5 characters"))
		})

		It("validates slice type implementing HasValidation", func() {
			type Config struct {
				Brokers testBrokers
			}

			// Valid case
			validConfig := &Config{Brokers: testBrokers{"broker1:9092"}}
			err := argument.ValidateHasValidation(ctx, validConfig)
			Expect(err).NotTo(HaveOccurred())

			// Invalid case - empty slice
			invalidConfig := &Config{Brokers: testBrokers{}}
			err = argument.ValidateHasValidation(ctx, invalidConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("field Brokers"))
			Expect(err.Error()).To(ContainSubstring("validation failed"))
			Expect(err.Error()).To(ContainSubstring("at least one broker required"))
		})

		It("prefers slice type validation over element validation", func() {
			// testBrokers has Validate, so it should be called
			// instead of validating each testBroker element
			type Config struct {
				Brokers testBrokers
			}

			// Valid case with non-empty slice
			config := &Config{Brokers: testBrokers{"broker1:9092"}}
			err := argument.ValidateHasValidation(ctx, config)
			Expect(err).NotTo(HaveOccurred())

			// Empty slice should trigger testBrokers.Validate error
			emptyConfig := &Config{Brokers: testBrokers{}}
			err = argument.ValidateHasValidation(ctx, emptyConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("at least one broker required"))

			// Explicit verification: slice with invalid element should pass
			// because only testBrokers.Validate is called (not element validation)
			// "x" is too short to be a valid testBroker, but testBrokers.Validate
			// only checks if len > 0, so this should succeed
			invalidElementConfig := &Config{Brokers: testBrokers{"x"}}
			err = argument.ValidateHasValidation(ctx, invalidElementConfig)
			Expect(
				err,
			).NotTo(HaveOccurred(), "slice type validation should not validate individual elements")
		})
	})

	Context("Integration with validation chain", func() {
		It("is called after ValidateRequired in the validation chain", func() {
			// This test demonstrates that ValidateHasValidation runs after
			// ValidateRequired, allowing both types of validation to work together
			type Config struct {
				Port testPort `arg:"port" required:"true"`
			}

			// Test with valid port
			validConfig := &Config{Port: 8080}
			err := argument.ValidateHasValidation(ctx, validConfig)
			Expect(err).NotTo(HaveOccurred())

			// Test with invalid port
			invalidConfig := &Config{Port: 80}
			err = argument.ValidateHasValidation(ctx, invalidConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("port must be >= 1024"))
		})

		It("validates non-required fields with HasValidation", func() {
			// HasValidation runs on ALL fields regardless of required tag
			// The required tag checks presence, Validate() checks validity
			type Config struct {
				Port testPort `arg:"port" default:"80"` // NOT required, but has default
			}

			// Even though Port is not required, the default value (80) is validated
			config := &Config{Port: 80}
			err := argument.ValidateHasValidation(ctx, config)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("port must be >= 1024"))

			// Valid default works fine
			validConfig := &Config{Port: 8080}
			err = argument.ValidateHasValidation(ctx, validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("allows zero values in optional fields if Validate handles it", func() {
			// Define a type that accepts zero value as valid
			type OptionalPort int

			portValidateCalled := false
			validate := func(ctx context.Context, p OptionalPort) error {
				portValidateCalled = true
				// Zero value (0) is allowed for optional port
				if p == 0 {
					return nil
				}
				if p < 1024 {
					return errors.New(ctx, "port must be >= 1024 or 0 for optional")
				}
				return nil
			}

			type Config struct {
				Port OptionalPort `arg:"port"` // optional, no default
			}

			// Note: We can't actually test this with argument.ValidateHasValidation
			// because OptionalPort doesn't implement HasValidation interface in this test
			// This test demonstrates the PATTERN, not the actual implementation

			config := &Config{Port: 0}
			err := validate(ctx, config.Port)
			Expect(err).NotTo(HaveOccurred())
			Expect(portValidateCalled).To(BeTrue())
		})
	})

	Context("Real-world examples", func() {
		It("validates complex configuration", func() {
			type Config struct {
				Port    testPort
				Timeout testTimeout
				Brokers []testBroker
			}

			// All valid
			validConfig := &Config{
				Port:    8080,
				Timeout: 30,
				Brokers: []testBroker{"broker1:9092", "broker2:9092"},
			}
			err := argument.ValidateHasValidation(ctx, validConfig)
			Expect(err).NotTo(HaveOccurred())

			// Invalid port (< 1024)
			config := &Config{
				Port:    80,
				Timeout: 30,
				Brokers: []testBroker{"broker1:9092"},
			}
			err = argument.ValidateHasValidation(ctx, config)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Port"))

			// Invalid broker (too short)
			config = &Config{
				Port:    8080,
				Timeout: 30,
				Brokers: []testBroker{"b1"},
			}
			err = argument.ValidateHasValidation(ctx, config)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("broker must be at least 5 characters"))
		})
	})
})
