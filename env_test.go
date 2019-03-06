// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
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
	It("default string", func() {
		var args struct {
			Username string `env:"user" default:"Ben"`
		}
		err := argument.ParseEnv(&args, []string{})
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
	It("default int", func() {
		var args struct {
			Age int `env:"age" default:"29"`
		}
		err := argument.ParseEnv(&args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(29))
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
	It("default bool true", func() {
		var args struct {
			Confirm bool `env:"confirm" default:"true"`
		}
		err := argument.ParseEnv(&args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeTrue())
	})
	It("default bool false", func() {
		var args struct {
			Confirm bool `env:"confirm" default:"false"`
		}
		err := argument.ParseEnv(&args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeFalse())
	})
})
