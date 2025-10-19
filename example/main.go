// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bborbe/errors"
	libtime "github.com/bborbe/time"

	"github.com/bborbe/argument/v2"
)

type Username string

type Password string

type Active bool

type Environment string

type Brokers []Broker

func (b Brokers) Validate(ctx context.Context) error {
	if len(b) == 0 {
		return errors.New(ctx, "list of brokers should not be empty")
	}
	for _, broker := range b {
		if err := broker.Validate(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Broker demonstrates encoding.TextUnmarshaler for custom parsing logic.
// It adds a default "plain://" schema if none is provided.
type Broker string

func (b Broker) Validate(ctx context.Context) error {
	if len(b) == 0 {
		return errors.New(ctx, "empty Broker")
	}
	// Basic format check - broker should be in host:port format
	s := string(b)
	if idx := strings.Index(s, "://"); idx != -1 {
		s = s[idx+3:]
	}
	if !strings.Contains(s, ":") {
		return errors.New(ctx, "broker must contain host:port")
	}
	return nil
}

func (b *Broker) UnmarshalText(text []byte) error {
	value := string(text)
	if strings.Contains(value, "://") {
		*b = Broker(value)
		return nil
	}
	// Add default plain:// schema if missing
	*b = Broker("plain://" + value)
	return nil
}

func (b Broker) String() string {
	return string(b)
}

func main() {
	ctx := context.Background()
	var data struct {
		// Basic types
		Username          Username `arg:"username" default:"ben"`
		Password          Password `arg:"password" display:"length"`
		Active            *Active  `arg:"active"`
		Url               string   `arg:"url"`
		DefaultWithoutArg string   `arg:"defaultWithoutArg" default:"hello world"`
		DefaultWithArg    string   `arg:"defaultWithArg" default:"hello world"`
		Int               int      `arg:"int"`
		Float64           float64  `arg:"float64"`
		Float64Ptr        *float64 `arg:"float64Ptr"`

		// Slice types - string slices
		Names        []string      `arg:"names" env:"NAMES" default:"alice,bob"`
		Tags         []string      `arg:"tags" env:"TAGS" default:"prod,api"`
		Environments []Environment `arg:"environments" separator:"|" default:"dev|staging|prod"`

		// Slice types - integer slices
		Ports      []int    `arg:"ports" env:"PORTS" separator:":"`
		IDs        []int64  `arg:"ids" default:"1001,1002,1003"`
		Counters   []uint   `arg:"counters" separator:";"`
		BigNumbers []uint64 `arg:"big-numbers"`

		// Slice types - float and bool slices
		Prices []float64 `arg:"prices" default:"9.99,19.99,29.99"`
		Flags  []bool    `arg:"flags" default:"true,false,true"`

		// Slice types - custom type slices
		AllowedUsers []Username `arg:"allowed_users" env:"ALLOWED_USERS"`

		// TextUnmarshaler types - custom parsing logic
		Broker     Broker   `arg:"broker" env:"BROKER" default:"localhost:9092"`
		BrokerList []Broker `arg:"broker-list" env:"BROKER_LIST"`
		Brokers    Brokers  `arg:"brokers" env:"BROKERS"`

		// Time types
		StdDuration time.Duration    `arg:"std-duration"`
		StdTime     time.Time        `arg:"std-time"`
		Duration    libtime.Duration `arg:"duration"`
		DateTime    libtime.DateTime `arg:"datetime"`
		Date        libtime.Date     `arg:"date"`
		UnixTime    libtime.UnixTime `arg:"unixtime"`
	}
	if err := argument.ParseAndPrint(ctx, &data); err != nil {
		log.Fatalf("parse args failed: %v", err)
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		log.Fatalf("encode data failed: %v", err)
	}

	// Example validation failures (commented out):
	//
	// 1. Broker validation - empty broker:
	//    emptyBroker := Broker("")
	//    err := emptyBroker.Validate(ctx)
	//    // Error: "empty Broker"
	//
	// 2. Broker validation - missing port:
	//    invalidBroker := Broker("localhost")
	//    err := invalidBroker.Validate(ctx)
	//    // Error: "broker must contain host:port"
	//
	// 3. Brokers validation - empty list:
	//    emptyBrokers := Brokers{}
	//    err := emptyBrokers.Validate(ctx)
	//    // Error: "list of brokers should not be empty"
	//
	// 4. Brokers validation - invalid broker in list:
	//    invalidBrokers := Brokers{Broker("localhost:9092"), Broker("")}
	//    err := invalidBrokers.Validate(ctx)
	//    // Error: "empty Broker"
}
