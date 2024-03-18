// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"bytes"
	"context"
	"flag"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	argument "github.com/bborbe/argument/v2"
)

var _ = Describe("Parse", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		flag.CommandLine.SetOutput(&bytes.Buffer{})
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		os.Args = []string{"go"}
		os.Clearenv()
	})
	It("parse float64 from arg default", func() {
		var args struct {
			Amount float64 `arg:"amount" env:"Amount"`
		}
		os.Args = []string{"go"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Amount).To(Equal(float64(0)))
	})
	It("parse float64 from arg", func() {
		var args struct {
			Amount float64 `arg:"amount" env:"Amount"`
		}
		os.Args = []string{"go", "-amount=23.5"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Amount).To(Equal(23.5))
	})
	It("parse *float64 from arg", func() {
		var args struct {
			Amount *float64 `arg:"amount" env:"Amount"`
		}
		os.Args = []string{"go", "-amount=23.5"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Amount).NotTo(BeNil())
		Expect(*args.Amount).To(Equal(23.5))
	})
	It("parse *float64 from arg", func() {
		var args struct {
			Amount *float64 `arg:"amount" env:"Amount"`
		}
		os.Args = []string{"go", "-amount=23.5"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Amount).NotTo(BeNil())
		Expect(*args.Amount).To(Equal(23.5))
	})
	It("parse *float64 from arg default", func() {
		var args struct {
			Amount *float64 `arg:"amount" env:"Amount"`
		}
		os.Args = []string{"go"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Amount).To(BeNil())
	})
	It("parse string from arg", func() {
		var args struct {
			Username string `arg:"user" env:"user"`
		}
		os.Args = []string{"go", "-user=Ben"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse string from env", func() {
		var args struct {
			Username string `arg:"user" env:"user"`
		}
		_ = os.Setenv("user", "Ben")
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("use default if env and args are not found", func() {
		var args struct {
			Username string `arg:"user" env:"user" default:"Ben"`
		}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("use env if both are defined", func() {
		var args struct {
			Username string `arg:"user" env:"user"`
		}
		os.Args = []string{"go", "-user=Arg"}
		_ = os.Setenv("user", "Env")
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Env"))
	})
	It("use flag if defined", func() {
		var args struct {
			Username string `arg:"user" env:"user" default:"Default"`
		}
		os.Args = []string{"go", "-user=Arg"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Arg"))
	})
})
