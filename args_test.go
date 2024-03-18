// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"bytes"
	"context"
	"flag"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	argument "github.com/bborbe/argument/v2"
)

var _ = Describe("ParseArgs", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		flag.CommandLine.SetOutput(&bytes.Buffer{})
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	})
	It("parse empty struct", func() {
		var args struct {
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).To(BeNil())
	})
	It("ignore private fields", func() {
		var args struct {
			private string
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).To(BeNil())
	})
	It("parse string from args parameter", func() {
		var args struct {
			Username string `arg:"user"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-user=Ben"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse string from default", func() {
		var args struct {
			Username string `arg:"user" default:"Ben"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse int", func() {
		var args struct {
			Age int `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(29))
	})
	It("returns an error if parse int fails", func() {
		var args struct {
			Age int `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=abc"})
		Expect(err).To(HaveOccurred())
	})
	It("return error if parse fails", func() {
		var args struct {
			Age int `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=abc"})
		Expect(err).To(HaveOccurred())
	})
	It("skip fields without tag", func() {
		var args struct {
			Age int
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(0))
	})
	It("default int", func() {
		var args struct {
			Age int `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(29))
	})
	It("parse int64", func() {
		var args struct {
			Age int64 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(int64(29)))
	})
	It("default int64", func() {
		var args struct {
			Age int64 `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(int64(29)))
	})
	It("parse bool true", func() {
		var args struct {
			Confirm bool `arg:"confirm"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-confirm=true"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeTrue())
	})
	It("parse bool false", func() {
		var args struct {
			Confirm bool `arg:"confirm"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-confirm=false"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeFalse())
	})
	It("default bool true", func() {
		var args struct {
			Confirm bool `arg:"confirm" default:"true"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeTrue())
	})
	It("default bool false", func() {
		var args struct {
			Confirm bool `arg:"confirm" default:"false"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeFalse())
	})
	It("returns an error if parse bool fails", func() {
		var args struct {
			Confirm bool `arg:"confirm"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-confirm=banana"})
		Expect(err).To(HaveOccurred())
	})
	It("parse duration", func() {
		var args struct {
			Wait time.Duration `arg:"wait"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-wait=1m"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Wait).To(Equal(time.Minute))
	})
	It("default duration", func() {
		var args struct {
			Wait time.Duration `arg:"wait" default:"1m"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Wait).To(Equal(time.Minute))
	})
	It("parse duration days", func() {
		var args struct {
			Wait time.Duration `arg:"wait"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-wait=7d"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Wait).To(Equal(7 * 24 * time.Hour))
	})
	It("default duration days", func() {
		var args struct {
			Wait time.Duration `arg:"wait" default:"7d"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Wait).To(Equal(7 * 24 * time.Hour))
	})
	It("parse float64", func() {
		var args struct {
			Age float64 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(float64(29)))
	})
	It("parse *float64", func() {
		var args struct {
			Age *float64 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(*args.Age).To(Equal(float64(29)))
	})
	It("default float64", func() {
		var args struct {
			Age float64 `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(float64(29)))
	})
	It("parse *float64", func() {
		var args struct {
			Age *float64 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age="})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(BeNil())
	})
	It("default *float64", func() {
		var args struct {
			Age *float64 `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(*args.Age).To(Equal(float64(29)))
	})

	It("parse uint", func() {
		var args struct {
			Age uint `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(uint(29)))
	})
	It("default uint", func() {
		var args struct {
			Age uint `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(uint(29)))
	})
	It("parse uint64", func() {
		var args struct {
			Age uint64 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(uint64(29)))
	})
	It("default uint64", func() {
		var args struct {
			Age uint64 `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(uint64(29)))
	})
	It("returns an error if type is not supported", func() {
		var args struct {
			Age interface{} `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).To(HaveOccurred())
	})
	It("parse int32", func() {
		var args struct {
			Age int32 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(int32(29)))
	})
	It("default int32", func() {
		var args struct {
			Age int32 `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(int32(29)))
	})
})
