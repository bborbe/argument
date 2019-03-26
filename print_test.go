// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"bytes"
	"github.com/bborbe/argument"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
)

var _ = Describe("Print", func() {
	type app struct {
		Username string
		Password string `display:"length"`
		Debug    bool   `display:"hidden"`
	}
	var buf *bytes.Buffer
	var args app
	BeforeEach(func() {
		buf = &bytes.Buffer{}
		log.SetOutput(buf)
		log.SetFlags(0)
		args = app{
			Username: "Ben",
			Password: "S3CR3T",
			Debug:    true,
		}
	})
	It("print without error", func() {
		err := argument.Print(&args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("foo", func() {
		err := argument.Print(&args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(Equal("Argument: Username 'Ben'\nArgument: Password length 6\n"))
	})
})
