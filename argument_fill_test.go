// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bborbe/argument/v2"
)

var _ = Describe("Fill", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	It("fills map to struct", func() {
		var args struct {
			Username string
		}
		data := map[string]interface{}{
			"Username": "Ben",
		}
		err := argument.Fill(ctx, &args, data)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("returns error if decode json fails", func() {
		data := map[string]interface{}{
			"Username": "Ben",
		}
		err := argument.Fill(ctx, "", data)
		Expect(err).To(HaveOccurred())
	})

	It("returns error if encode json fails", func() {
		var args struct {
			Username string
		}
		// Create data that cannot be JSON encoded
		data := map[string]interface{}{
			"Username": make(chan int), // channels cannot be JSON encoded
		}
		err := argument.Fill(ctx, &args, data)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("encode json failed"))
	})

	It("handles complex data structures", func() {
		var args struct {
			Config map[string]interface{}
		}
		data := map[string]interface{}{
			"Config": map[string]interface{}{
				"key": "value",
				"nested": map[string]interface{}{
					"inner": "data",
				},
			},
		}
		err := argument.Fill(ctx, &args, data)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Config["key"]).To(Equal("value"))
	})
})
