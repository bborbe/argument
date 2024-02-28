// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bborbe/argument"
)

var _ = Describe("DefaultValues", func() {
	It("default string", func() {
		var args struct {
			Username string `default:"user"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).NotTo(HaveOccurred())
		Expect(data).To(HaveKeyWithValue("Username", "user"))
	})
	It("default int", func() {
		var args struct {
			Age int `default:"29"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).NotTo(HaveOccurred())
		Expect(data).To(HaveKeyWithValue("Age", 29))
	})
	It("return error if parse int fails", func() {
		var args struct {
			Age int `default:"age"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).To(HaveOccurred())
		Expect(data).To(BeNil())
	})
	It("default int64", func() {
		var args struct {
			Age int64 `default:"29"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).NotTo(HaveOccurred())
		value, ok := data["Age"]
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal(int64(29)))
	})
	It("return error if parse int64 fails", func() {
		var args struct {
			Age int64 `default:"age"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).To(HaveOccurred())
		Expect(data).To(BeNil())
	})
	It("default uint", func() {
		var args struct {
			Age uint `default:"29"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).NotTo(HaveOccurred())
		value, ok := data["Age"]
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal(uint64(29)))
	})
	It("return error if parse uint fails", func() {
		var args struct {
			Age uint `default:"age"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).To(HaveOccurred())
		Expect(data).To(BeNil())
	})
	It("default uint64", func() {
		var args struct {
			Age uint64 `default:"29"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).NotTo(HaveOccurred())
		value, ok := data["Age"]
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal(uint64(29)))
	})
	It("return error if parse uint64 fails", func() {
		var args struct {
			Age uint64 `default:"age"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).To(HaveOccurred())
		Expect(data).To(BeNil())
	})
	It("default int32", func() {
		var args struct {
			Age int32 `default:"29"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).NotTo(HaveOccurred())
		value, ok := data["Age"]
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal(int32(29)))
	})
	It("return error if parse int32 fails", func() {
		var args struct {
			Age int32 `default:"age"`
		}
		data, err := argument.DefaultValues(&args)
		Expect(err).To(HaveOccurred())
		Expect(data).To(BeNil())
	})
})
