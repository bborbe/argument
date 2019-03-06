// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"os"

	"github.com/bborbe/argument"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse", func() {
	BeforeEach(func() {
		os.Args = []string{"go"}
		os.Clearenv()
	})
	It("parse string from arg", func() {
		var args struct {
			Username string `arg:"user" env:"user"`
		}
		os.Args = []string{"go", "-user=Ben"}
		err := argument.Parse(&args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse string from env", func() {
		var args struct {
			Username string `arg:"user" env:"user"`
		}
		_ = os.Setenv("user", "Ben")
		err := argument.Parse(&args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("use default if env and args are not found", func() {
		var args struct {
			Username string `arg:"user" env:"user" default:"Ben"`
		}
		err := argument.Parse(&args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("use env if both are defined", func() {
		var args struct {
			Username string `arg:"user" env:"user"`
		}
		os.Args = []string{"go", "-user=Arg"}
		_ = os.Setenv("user", "Env")
		err := argument.Parse(&args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Env"))
	})
})
