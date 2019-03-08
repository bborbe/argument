// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"github.com/bborbe/argument"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Print", func() {
	It("print without error", func() {
		args := struct {
			Username string
			Password string `display:"length"`
			Debug    bool   `display:"hidden"`
		}{
			Username: "Ben",
		}
		err := argument.Print(&args)
		Expect(err).NotTo(HaveOccurred())
	})
})
