// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bborbe/argument/v2"
)

var _ = Describe("Required validation examples", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})

	Context("Basic types with required", func() {
		It("validates required string field", func() {
			type Config struct {
				Username string `required:"true" arg:"username" env:"USERNAME"`
			}

			// Empty string should fail
			emptyConfig := Config{Username: ""}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			// Non-empty string should pass
			validConfig := Config{Username: "admin"}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required int field", func() {
			type Config struct {
				Port int `required:"true" arg:"port"`
			}

			// Zero value should fail
			emptyConfig := Config{Port: 0}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			// Non-zero value should pass
			validConfig := Config{Port: 8080}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required int64 field", func() {
			type Config struct {
				ID int64 `required:"true" arg:"id"`
			}

			emptyConfig := Config{ID: 0}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{ID: 12345}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required int32 field", func() {
			type Config struct {
				Count int32 `required:"true" arg:"count"`
			}

			emptyConfig := Config{Count: 0}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{Count: 100}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required uint field", func() {
			type Config struct {
				Workers uint `required:"true" arg:"workers"`
			}

			emptyConfig := Config{Workers: 0}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{Workers: 4}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required uint64 field", func() {
			type Config struct {
				MaxSize uint64 `required:"true" arg:"max-size"`
			}

			emptyConfig := Config{MaxSize: 0}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{MaxSize: 1024}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required float64 field", func() {
			type Config struct {
				Threshold float64 `required:"true" arg:"threshold"`
			}

			emptyConfig := Config{Threshold: 0.0}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{Threshold: 0.95}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required bool field - always passes", func() {
			type Config struct {
				Enabled bool `required:"true" arg:"enabled"`
			}

			// Bool is never considered empty, both true and false are valid
			falseConfig := Config{Enabled: false}
			err := argument.ValidateRequired(ctx, &falseConfig)
			Expect(err).NotTo(HaveOccurred())

			trueConfig := Config{Enabled: true}
			err = argument.ValidateRequired(ctx, &trueConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required *float64 pointer field", func() {
			type Config struct {
				Rate *float64 `required:"true" arg:"rate"`
			}

			// Nil pointer should fail
			emptyConfig := Config{Rate: nil}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			// Non-nil pointer should pass
			rate := 3.14
			validConfig := Config{Rate: &rate}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required time.Duration field", func() {
			type Config struct {
				Timeout time.Duration `required:"true" arg:"timeout"`
			}

			emptyConfig := Config{Timeout: 0}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{Timeout: 30 * time.Second}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Slice types with required", func() {
		It("validates required []string slice", func() {
			type Config struct {
				Hosts []string `required:"true" arg:"hosts"`
			}

			// Empty slice should fail
			emptyConfig := Config{Hosts: []string{}}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			// Nil slice should fail
			nilConfig := Config{Hosts: nil}
			err = argument.ValidateRequired(ctx, &nilConfig)
			Expect(err).To(HaveOccurred())

			// Non-empty slice should pass
			validConfig := Config{Hosts: []string{"host1", "host2"}}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required []int slice", func() {
			type Config struct {
				Ports []int `required:"true" arg:"ports"`
			}

			emptyConfig := Config{Ports: []int{}}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{Ports: []int{8080, 9090}}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required []float64 slice", func() {
			type Config struct {
				Thresholds []float64 `required:"true" arg:"thresholds"`
			}

			emptyConfig := Config{Thresholds: []float64{}}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{Thresholds: []float64{0.5, 0.75, 0.9}}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required []bool slice", func() {
			type Config struct {
				Features []bool `required:"true" arg:"features"`
			}

			emptyConfig := Config{Features: []bool{}}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{Features: []bool{true, false, true}}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Custom types with required", func() {
		type Username string
		type Port int
		type Rate float64
		type Enabled bool

		It("validates required custom string type", func() {
			type Config struct {
				User Username `required:"true" arg:"user"`
			}

			emptyConfig := Config{User: ""}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{User: "admin"}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required custom int type", func() {
			type Config struct {
				Port Port `required:"true" arg:"port"`
			}

			emptyConfig := Config{Port: 0}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{Port: 8080}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required custom float64 type", func() {
			type Config struct {
				Rate Rate `required:"true" arg:"rate"`
			}

			emptyConfig := Config{Rate: 0.0}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{Rate: 1.5}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required custom bool type - always passes", func() {
			type Config struct {
				Enabled Enabled `required:"true" arg:"enabled"`
			}

			// Custom bool types are never considered empty
			falseConfig := Config{Enabled: false}
			err := argument.ValidateRequired(ctx, &falseConfig)
			Expect(err).NotTo(HaveOccurred())

			trueConfig := Config{Enabled: true}
			err = argument.ValidateRequired(ctx, &trueConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates required custom type slice", func() {
			type Broker string
			type Config struct {
				Brokers []Broker `required:"true" arg:"brokers"`
			}

			emptyConfig := Config{Brokers: []Broker{}}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())

			validConfig := Config{Brokers: []Broker{"broker1:9092", "broker2:9092"}}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Real-world examples", func() {
		It("validates Kafka brokers configuration like core-report-controller", func() {
			type KafkaConfig struct {
				KafkaBrokers []string `required:"true" arg:"kafka-brokers" env:"KAFKA_BROKERS"`
			}

			// Should fail with empty slice
			emptyConfig := KafkaConfig{
				KafkaBrokers: []string{},
			}
			err := argument.ValidateRequired(ctx, &emptyConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Required field empty"))

			// Should succeed with values
			validConfig := KafkaConfig{
				KafkaBrokers: []string{"my-cluster-kafka-bootstrap.strimzi.svc.cluster.local:9092"},
			}
			err = argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates multiple required slices in application config", func() {
			type AppConfig struct {
				Servers  []string  `required:"true"  arg:"servers"`
				Ports    []int     `required:"true"  arg:"ports"`
				Timeouts []float64 `required:"false" arg:"timeouts"`
				Features []string  // not required
			}

			// All required fields populated
			config := AppConfig{
				Servers:  []string{"server1", "server2"},
				Ports:    []int{8080, 9090},
				Timeouts: []float64{}, // not required, can be empty
				Features: nil,         // not required, can be nil
			}
			err := argument.ValidateRequired(ctx, &config)
			Expect(err).NotTo(HaveOccurred())

			// Missing required Servers
			invalidConfig := AppConfig{
				Servers: []string{}, // empty but required!
				Ports:   []int{8080},
			}
			err = argument.ValidateRequired(ctx, &invalidConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("servers"))
		})

		It("handles mixed required types including slices", func() {
			type MixedConfig struct {
				AppName      string   `required:"true"  arg:"app-name"`
				Port         int      `required:"true"  arg:"port"`
				Brokers      []string `required:"true"  arg:"brokers"`
				OptionalTags []string `required:"false" arg:"tags"`
				Debug        bool     `required:"false" arg:"debug"`
			}

			validConfig := MixedConfig{
				AppName:      "my-service",
				Port:         8080,
				Brokers:      []string{"broker1:9092"},
				OptionalTags: []string{}, // ok to be empty
				Debug:        false,
			}
			err := argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())

			// Missing Brokers
			invalidConfig := MixedConfig{
				AppName: "my-service",
				Port:    8080,
				Brokers: nil, // required but nil!
			}
			err = argument.ValidateRequired(ctx, &invalidConfig)
			Expect(err).To(HaveOccurred())
		})

		It("validates custom type slices with required", func() {
			type Broker string
			type Stage string

			type TradingConfig struct {
				Brokers []Broker `required:"true"  env:"KAFKA_BROKERS"`
				Stages  []Stage  `required:"false" env:"STAGES"`
			}

			validConfig := TradingConfig{
				Brokers: []Broker{"broker1:9092", "broker2:9092"},
				Stages:  []Stage{}, // not required
			}
			err := argument.ValidateRequired(ctx, &validConfig)
			Expect(err).NotTo(HaveOccurred())

			invalidConfig := TradingConfig{
				Brokers: []Broker{}, // required but empty!
			}
			err = argument.ValidateRequired(ctx, &invalidConfig)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("KAFKA_BROKERS"))
		})
	})
})
