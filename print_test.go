// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"bytes"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bborbe/argument"
)

var _ = Describe("Print", func() {
	type app struct {
		Username      string
		Password      string `display:"length"`
		Debug         bool   `display:"hidden"`
		Float64       float64
		Float64Ptr    *float64
		Float64PtrNil *float64
	}
	var buf *bytes.Buffer
	var args app
	BeforeEach(func() {
		buf = &bytes.Buffer{}
		log.SetOutput(buf)
		log.SetFlags(0)
		args = app{
			Username:      "Ben",
			Password:      "S3CR3T",
			Debug:         true,
			Float64:       13.37,
			Float64Ptr:    float64Ptr(4.2),
			Float64PtrNil: nil,
		}
	})
	It("print without error", func() {
		err := argument.Print(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("foo", func() {
		err := argument.Print(&args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(Equal(`Argument: Username 'Ben'
Argument: Password length 6
Argument: Float64 '13.37'
Argument: Float64Ptr '4.2'
Argument: Float64PtrNil <nil>
`))
	})
})

func float64Ptr(value float64) *float64 {
	return &value
}
