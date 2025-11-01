// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"bytes"
	"context"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/argument/v2"
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
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
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
		err := argument.Print(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("foo", func() {
		err := argument.Print(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(Equal(`Argument: Username 'Ben'
Argument: Password length 6
Argument: Float64 '13.37'
Argument: Float64Ptr '4.2'
Argument: Float64PtrNil <nil>
`))
	})
})

var _ = Describe("Print with slices", func() {
	type appWithSlices struct {
		EmptySlice    []string
		SingleString  []string
		MultiStrings  []string
		Numbers       []int
		Floats        []float64
		Bools         []bool
		NonSliceField string
	}
	var buf *bytes.Buffer
	var args appWithSlices
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		buf = &bytes.Buffer{}
		log.SetOutput(buf)
		log.SetFlags(0)
		args = appWithSlices{
			EmptySlice:    []string{},
			SingleString:  []string{"alone"},
			MultiStrings:  []string{"alice", "bob", "charlie"},
			Numbers:       []int{1, 2, 3, 4},
			Floats:        []float64{1.5, 2.5, 3.5},
			Bools:         []bool{true, false, true},
			NonSliceField: "regular",
		}
	})
	It("prints empty slice correctly", func() {
		err := argument.Print(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(ContainSubstring("Argument: EmptySlice []"))
	})
	It("prints single element slice correctly", func() {
		err := argument.Print(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(ContainSubstring("Argument: SingleString [1]: alone"))
	})
	It("prints multiple string slice with comma separation", func() {
		err := argument.Print(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(ContainSubstring("Argument: MultiStrings [3]: alice, bob, charlie"))
	})
	It("prints integer slice correctly", func() {
		err := argument.Print(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(ContainSubstring("Argument: Numbers [4]: 1, 2, 3, 4"))
	})
	It("prints float slice correctly", func() {
		err := argument.Print(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(ContainSubstring("Argument: Floats [3]: 1.5, 2.5, 3.5"))
	})
	It("prints bool slice correctly", func() {
		err := argument.Print(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(ContainSubstring("Argument: Bools [3]: true, false, true"))
	})
	It("prints non-slice field with quotes", func() {
		err := argument.Print(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(ContainSubstring("Argument: NonSliceField 'regular'"))
	})
	It("prints complete output", func() {
		err := argument.Print(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(buf.String()).To(Equal(`Argument: EmptySlice []
Argument: SingleString [1]: alone
Argument: MultiStrings [3]: alice, bob, charlie
Argument: Numbers [4]: 1, 2, 3, 4
Argument: Floats [3]: 1.5, 2.5, 3.5
Argument: Bools [3]: true, false, true
Argument: NonSliceField 'regular'
`))
	})
})

func float64Ptr(value float64) *float64 {
	return &value
}
