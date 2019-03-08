// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"github.com/bborbe/argument"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fill", func() {
	It("fills map to struct", func() {
		var args struct {
			Username string
		}
		data := map[string]interface{}{
			"Username": "Ben",
		}
		err := argument.Fill(&args, data)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("returns error if decode json fails", func() {
		data := map[string]interface{}{
			"Username": "Ben",
		}
		err := argument.Fill("", data)
		Expect(err).To(HaveOccurred())
	})
})
