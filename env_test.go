// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"time"

	"github.com/bborbe/argument"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseEnv", func() {
	It("parse string", func() {
		var args struct {
			Username string `env:"user"`
		}
		err := argument.ParseEnv(&args, []string{"user=Ben"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse int", func() {
		var args struct {
			Age int `env:"age"`
		}
		err := argument.ParseEnv(&args, []string{"age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(29))
	})
	It("return error if parse int fails", func() {
		var args struct {
			Age int `env:"age"`
		}
		err := argument.ParseEnv(&args, []string{"age=abc"})
		Expect(err).To(HaveOccurred())
	})
	It("return error if parse int64 fails", func() {
		var args struct {
			Age int64 `env:"age"`
		}
		err := argument.ParseEnv(&args, []string{"age=abc"})
		Expect(err).To(HaveOccurred())
	})
	It("parse int64", func() {
		var args struct {
			Age int64 `env:"age"`
		}
		err := argument.ParseEnv(&args, []string{"age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(int64(29)))
	})
	It("parse bool true", func() {
		var args struct {
			Confirm bool `env:"confirm"`
		}
		err := argument.ParseEnv(&args, []string{"confirm=true"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeTrue())
	})
	It("parse bool false", func() {
		var args struct {
			Confirm bool `env:"confirm"`
		}
		err := argument.ParseEnv(&args, []string{"confirm=false"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeFalse())
	})
	It("returns an error if parse bool fails", func() {
		var args struct {
			Confirm bool `env:"confirm"`
		}
		err := argument.ParseEnv(&args, []string{"confirm=hello"})
		Expect(err).To(HaveOccurred())
	})
	It("parse duration", func() {
		var args struct {
			Wait time.Duration `env:"wait"`
		}
		err := argument.ParseEnv(&args, []string{"wait=1m"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Wait).To(Equal(time.Minute))
	})
	It("return an error if parse duration fails", func() {
		var args struct {
			Wait time.Duration `env:"wait"`
		}
		err := argument.ParseEnv(&args, []string{"wait=hello"})
		Expect(err).To(HaveOccurred())
	})
	It("parse float64", func() {
		var args struct {
			Age float64 `env:"age"`
		}
		err := argument.ParseEnv(&args, []string{"age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(float64(29)))
	})
	It("return error if parse float64 fails", func() {
		var args struct {
			Age float64 `env:"age"`
		}
		err := argument.ParseEnv(&args, []string{"age=abc"})
		Expect(err).To(HaveOccurred())
	})
	It("parse uint", func() {
		var args struct {
			Age uint `env:"age"`
		}
		err := argument.ParseEnv(&args, []string{"age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(uint(29)))
	})
	It("return error if parse uint fails", func() {
		var args struct {
			Age uint `env:"age"`
		}
		err := argument.ParseEnv(&args, []string{"age=abc"})
		Expect(err).To(HaveOccurred())
	})
	It("parse uint64", func() {
		var args struct {
			Age uint64 `env:"age"`
		}
		err := argument.ParseEnv(&args, []string{"age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(uint64(29)))
	})
	It("return error if parse uint64 fails", func() {
		var args struct {
			Age uint64 `env:"age"`
		}
		err := argument.ParseEnv(&args, []string{"age=abc"})
		Expect(err).To(HaveOccurred())
	})
	It("skip fields without tag", func() {
		var args struct {
			Age int
		}
		err := argument.ParseEnv(&args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(0))
	})
	It("returns an error if type is not supported", func() {
		var args struct {
			Age interface{} `env:"age" default:"29"`
		}
		err := argument.ParseEnv(&args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
})
