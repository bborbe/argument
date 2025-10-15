// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	libtime "github.com/bborbe/time"

	"github.com/bborbe/argument/v2"
)

type Username string

type Password string

type Active bool

func main() {
	ctx := context.Background()
	var data struct {
		Username          Username         `arg:"username" default:"ben"`
		Password          Password         `arg:"password" display:"length"`
		Active            *Active          `arg:"active"`
		Url               string           `arg:"url"`
		DefaultWithoutArg string           `arg:"defaultWithoutArg" default:"hello world"`
		DefaultWithArg    string           `arg:"defaultWithArg" default:"hello world"`
		Int               int              `arg:"int"`
		Float64           float64          `arg:"float64"`
		Float64Ptr        *float64         `arg:"float64Ptr"`
		StdDuration       time.Duration    `arg:"std-duration"`
		StdTime           time.Time        `arg:"std-time"`
		Duration          libtime.Duration `arg:"duration"`
		DateTime          libtime.DateTime `arg:"datetime"`
		Date              libtime.Date     `arg:"date"`
		UnixTime          libtime.UnixTime `arg:"unixtime"`
	}
	if err := argument.Parse(ctx, &data); err != nil {
		log.Fatalf("parse args failed: %v", err)
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		log.Fatalf("encode data failed: %v", err)
	}
}
