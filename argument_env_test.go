// Copyright (c) 2019 Benjamin Borbe All rights reserved.
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

var _ = Describe("ParseEnv", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	It("parse string", func() {
		var args struct {
			Username string `env:"user"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"user=Ben"})
		Expect(err).To(BeNil())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse int", func() {
		var args struct {
			Age int `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(29))
	})
	It("return error if parse int fails", func() {
		var args struct {
			Age int `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("return error if parse int64 fails", func() {
		var args struct {
			Age int64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("parse int64", func() {
		var args struct {
			Age int64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(int64(29)))
	})
	It("parse bool true", func() {
		var args struct {
			Confirm bool `env:"confirm"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"confirm=true"})
		Expect(err).To(BeNil())
		Expect(args.Confirm).To(BeTrue())
	})
	It("parse bool false", func() {
		var args struct {
			Confirm bool `env:"confirm"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"confirm=false"})
		Expect(err).To(BeNil())
		Expect(args.Confirm).To(BeFalse())
	})
	It("returns an error if parse bool fails", func() {
		var args struct {
			Confirm bool `env:"confirm"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"confirm=hello"})
		Expect(err).NotTo(BeNil())
	})
	It("parse duration", func() {
		var args struct {
			Wait time.Duration `env:"wait"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"wait=1m"})
		Expect(err).To(BeNil())
		Expect(args.Wait).To(Equal(time.Minute))
	})
	It("parse duration days", func() {
		var args struct {
			Wait time.Duration `env:"wait"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"wait=1d"})
		Expect(err).To(BeNil())
		Expect(args.Wait).To(Equal(24 * time.Hour))
	})
	It("return an error if parse duration fails", func() {
		var args struct {
			Wait time.Duration `env:"wait"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"wait=hello"})
		Expect(err).NotTo(BeNil())
	})
	It("parse float64", func() {
		var args struct {
			Age float64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(float64(29)))
	})
	It("return error if parse float64 fails", func() {
		var args struct {
			Age float64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("parse uint", func() {
		var args struct {
			Age uint `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(uint(29)))
	})
	It("return error if parse uint fails", func() {
		var args struct {
			Age uint `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("parse uint64", func() {
		var args struct {
			Age uint64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(uint64(29)))
	})
	It("return error if parse uint64 fails", func() {
		var args struct {
			Age uint64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("skip fields without tag", func() {
		var args struct {
			Age int
		}
		err := argument.ParseEnv(ctx, &args, []string{})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(0))
	})
	It("returns an error if type is not supported", func() {
		var args struct {
			Age interface{} `env:"age" default:"29"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("parse int32", func() {
		var args struct {
			Age int32 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(int32(29)))
	})
	It("return error if parse int32 fails", func() {
		var args struct {
			Age int32 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
})
