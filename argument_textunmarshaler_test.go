// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bborbe/argument/v2"
)

// TestBroker is a test type that implements encoding.TextUnmarshaler
// It mimics the kafka.Broker behavior by adding a default schema if missing
type TestBroker string

func (b *TestBroker) UnmarshalText(text []byte) error {
	value := string(text)
	if strings.Contains(value, "://") {
		*b = TestBroker(value)
		return nil
	}
	// Add default plain:// schema if missing
	*b = TestBroker("plain://" + value)
	return nil
}

func (b TestBroker) String() string {
	return string(b)
}

// TestURL is another test type for more complex validation
type TestURL string

func (u *TestURL) UnmarshalText(text []byte) error {
	value := string(text)
	if value == "" {
		return fmt.Errorf("URL cannot be empty")
	}
	if !strings.HasPrefix(value, "http://") && !strings.HasPrefix(value, "https://") {
		return fmt.Errorf("URL must start with http:// or https://")
	}
	*u = TestURL(value)
	return nil
}

func (u TestURL) String() string {
	return string(u)
}

var _ = Describe("TextUnmarshaler", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		flag.CommandLine.SetOutput(&bytes.Buffer{})
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	})

	Context("single values", func() {
		It("parses broker from args with schema", func() {
			var args struct {
				Broker TestBroker `arg:"broker"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-broker=ssl://localhost:9092"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Broker.String()).To(Equal("ssl://localhost:9092"))
		})

		It("parses broker from args without schema (adds default)", func() {
			var args struct {
				Broker TestBroker `arg:"broker"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-broker=localhost:9092"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Broker.String()).To(Equal("plain://localhost:9092"))
		})

		It("parses broker from env variable", func() {
			var args struct {
				Broker TestBroker `env:"BROKER"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"BROKER=ssl://kafka:9092"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Broker.String()).To(Equal("ssl://kafka:9092"))
		})

		It("parses broker from default value", func() {
			var args struct {
				Broker TestBroker `arg:"broker" default:"localhost:9092"`
			}
			err := argument.ParseArgs(ctx, &args, []string{})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Broker.String()).To(Equal("plain://localhost:9092"))
		})

		It("returns error for invalid URL", func() {
			var args struct {
				URL TestURL `arg:"url"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-url=invalid-url"})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("URL must start with http://"))
		})

		It("validates URL from args", func() {
			var args struct {
				URL TestURL `arg:"url"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-url=https://example.com"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.URL.String()).To(Equal("https://example.com"))
		})
	})

	Context("slice values", func() {
		It("parses broker slice from args", func() {
			var args struct {
				Brokers []TestBroker `arg:"brokers"`
			}
			err := argument.ParseArgs(
				ctx,
				&args,
				[]string{"-brokers=localhost:9092,ssl://kafka:9093"},
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Brokers).To(HaveLen(2))
			Expect(args.Brokers[0].String()).To(Equal("plain://localhost:9092"))
			Expect(args.Brokers[1].String()).To(Equal("ssl://kafka:9093"))
		})

		It("parses broker slice from env", func() {
			var args struct {
				Brokers []TestBroker `env:"BROKERS"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"BROKERS=broker1:9092,broker2:9092"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Brokers).To(HaveLen(2))
			Expect(args.Brokers[0].String()).To(Equal("plain://broker1:9092"))
			Expect(args.Brokers[1].String()).To(Equal("plain://broker2:9092"))
		})

		It("parses broker slice from default", func() {
			var args struct {
				Brokers []TestBroker `arg:"brokers" default:"localhost:9092,kafka:9093"`
			}
			err := argument.ParseArgs(ctx, &args, []string{})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Brokers).To(HaveLen(2))
			Expect(args.Brokers[0].String()).To(Equal("plain://localhost:9092"))
			Expect(args.Brokers[1].String()).To(Equal("plain://kafka:9093"))
		})

		It("handles empty broker slice", func() {
			var args struct {
				Brokers []TestBroker `arg:"brokers"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-brokers="})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Brokers).To(HaveLen(0))
		})

		It("returns error for invalid URL in slice", func() {
			var args struct {
				URLs []TestURL `arg:"urls"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-urls=https://valid.com,invalid-url"})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("URL must start with http://"))
		})

		It("parses URL slice with custom separator", func() {
			var args struct {
				URLs []TestURL `arg:"urls" separator:"|"`
			}
			err := argument.ParseArgs(
				ctx,
				&args,
				[]string{"-urls=https://a.com|https://b.com|http://c.com"},
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.URLs).To(HaveLen(3))
			Expect(args.URLs[0].String()).To(Equal("https://a.com"))
			Expect(args.URLs[1].String()).To(Equal("https://b.com"))
			Expect(args.URLs[2].String()).To(Equal("http://c.com"))
		})

		It("trims whitespace from broker slice elements", func() {
			var args struct {
				Brokers []TestBroker `arg:"brokers"`
			}
			err := argument.ParseArgs(
				ctx,
				&args,
				[]string{"-brokers= localhost:9092 , kafka:9093 "},
			)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Brokers).To(HaveLen(2))
			Expect(args.Brokers[0].String()).To(Equal("plain://localhost:9092"))
			Expect(args.Brokers[1].String()).To(Equal("plain://kafka:9093"))
		})
	})

	Context("combined args and env", func() {
		It("parses broker from args and env separately", func() {
			var argsConfig struct {
				Broker TestBroker `arg:"broker"`
			}
			var envConfig struct {
				Broker TestBroker `env:"BROKER"`
			}

			err := argument.ParseArgs(ctx, &argsConfig, []string{"-broker=localhost:9092"})
			Expect(err).NotTo(HaveOccurred())
			Expect(argsConfig.Broker.String()).To(Equal("plain://localhost:9092"))

			err = argument.ParseEnv(ctx, &envConfig, []string{"BROKER=kafka:9093"})
			Expect(err).NotTo(HaveOccurred())
			Expect(envConfig.Broker.String()).To(Equal("plain://kafka:9093"))
		})

		It("overrides default with args", func() {
			var args struct {
				Broker TestBroker `arg:"broker" default:"default:9092"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-broker=override:9093"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Broker.String()).To(Equal("plain://override:9093"))
		})
	})
})
