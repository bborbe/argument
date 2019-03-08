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

var _ = Describe("Required", func() {
	It("returns error message for env", func() {
		args := struct {
			Username string `required:"true" env:"abc"`
		}{
			Username: "",
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Required field empty, define env abc"))
	})
	It("returns error message for arg", func() {
		args := struct {
			Username string `required:"true" arg:"abc"`
		}{
			Username: "",
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Required field empty, define parameter abc"))
	})
	It("returns error message for env and arg", func() {
		args := struct {
			Username string `required:"true" arg:"abc" env:"abc"`
		}{
			Username: "",
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Required field empty, define parameter abc or define env abc"))
	})
	It("returns no error if nothing is required", func() {
		args := struct {
			Username string
		}{
			Username: "Ben",
		}
		err := argument.ValidateRequired(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required string is empty", func() {
		args := struct {
			Username string `required:"true"`
		}{
			Username: "",
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required string is not empty", func() {
		args := struct {
			Username string `required:"true"`
		}{
			Username: "Ben",
		}
		err := argument.ValidateRequired(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required int is empty", func() {
		args := struct {
			Age int `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required int is not empty", func() {
		args := struct {
			Age int `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required int64 is empty", func() {
		args := struct {
			Age int64 `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required int64 is not empty", func() {
		args := struct {
			Age int64 `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required uint is empty", func() {
		args := struct {
			Age uint `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required uint is not empty", func() {
		args := struct {
			Age uint `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required uint64 is empty", func() {
		args := struct {
			Age uint64 `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required uint64 is not empty", func() {
		args := struct {
			Age uint64 `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required float64 is empty", func() {
		args := struct {
			Age float64 `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required float64 is not empty", func() {
		args := struct {
			Age float64 `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required time.Duration is empty", func() {
		args := struct {
			Age time.Duration `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required time.Duration is not empty", func() {
		args := struct {
			Age time.Duration `required:"true"`
		}{
			Age: time.Minute,
		}
		err := argument.ValidateRequired(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns no error if if bool", func() {
		var args struct {
			Confirm bool `required:"true"`
		}
		err := argument.ValidateRequired(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error type is not supported", func() {
		var args struct {
			Banana interface{} `required:"true"`
		}
		err := argument.ValidateRequired(&args)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("field Banana with type <nil> is unsupported"))
	})
})
