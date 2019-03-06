// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flagjson_test

import (
	"github.com/bborbe/flagjson"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flag JSON", func() {
	It("parse string", func() {
		var args struct {
			Username string `arg:"user"`
		}
		err := flagjson.ParseArgs(&args, []string{"-user=Ben"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("default string", func() {
		var args struct {
			Username string `arg:"user" default:"Ben"`
		}
		err := flagjson.ParseArgs(&args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse int", func() {
		var args struct {
			Age int `arg:"age"`
		}
		err := flagjson.ParseArgs(&args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(29))
	})
	It("default int", func() {
		var args struct {
			Age int `arg:"age" default:"29"`
		}
		err := flagjson.ParseArgs(&args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(29))
	})
	It("parse bool true", func() {
		var args struct {
			Confirm bool `arg:"confirm"`
		}
		err := flagjson.ParseArgs(&args, []string{"-confirm=true"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeTrue())
	})
	It("parse bool false", func() {
		var args struct {
			Confirm bool `arg:"confirm"`
		}
		err := flagjson.ParseArgs(&args, []string{"-confirm=false"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeFalse())
	})
	It("default bool true", func() {
		var args struct {
			Confirm bool `arg:"confirm" default:"true"`
		}
		err := flagjson.ParseArgs(&args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeTrue())
	})
	It("default bool false", func() {
		var args struct {
			Confirm bool `arg:"confirm" default:"false"`
		}
		err := flagjson.ParseArgs(&args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeFalse())
	})
})
