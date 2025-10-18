// Copyright (c) 2025 Benjamin Borbe All rights reserved.
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

	"github.com/bborbe/argument/v2"
)

var _ = Describe("Slice parsing", func() {
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.Background()
		flag.CommandLine.SetOutput(&bytes.Buffer{})
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	})

	Context("[]string", func() {
		It("parses comma-separated values", func() {
			var args struct {
				Names []string `arg:"names"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-names=alice,bob,charlie"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{"alice", "bob", "charlie"}))
		})

		It("trims whitespace from each element", func() {
			var args struct {
				Names []string `arg:"names"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-names=alice, bob , charlie"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{"alice", "bob", "charlie"}))
		})

		It("handles empty string", func() {
			var args struct {
				Names []string `arg:"names"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-names="})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{}))
			Expect(len(args.Names)).To(Equal(0))
		})

		It("skips empty elements after trim", func() {
			var args struct {
				Names []string `arg:"names"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-names=alice,,bob,  ,charlie"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{"alice", "bob", "charlie"}))
		})

		It("uses custom separator", func() {
			var args struct {
				Names []string `arg:"names" separator:":"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-names=alice:bob:charlie"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{"alice", "bob", "charlie"}))
		})

		It("uses default value", func() {
			var args struct {
				Names []string `arg:"names" default:"alice,bob"`
			}
			err := argument.ParseArgs(ctx, &args, []string{})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{"alice", "bob"}))
		})

		It("overrides default with argument", func() {
			var args struct {
				Names []string `arg:"names" default:"alice,bob"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-names=charlie,dave"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{"charlie", "dave"}))
		})
	})

	Context("[]int", func() {
		It("parses comma-separated integers", func() {
			var args struct {
				Ports []int `arg:"ports"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-ports=8080,8081,8082"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Ports).To(Equal([]int{8080, 8081, 8082}))
		})

		It("trims whitespace before parsing", func() {
			var args struct {
				Ports []int `arg:"ports"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-ports=8080, 8081 , 8082"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Ports).To(Equal([]int{8080, 8081, 8082}))
		})

		It("returns error for invalid integer", func() {
			var args struct {
				Ports []int `arg:"ports"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-ports=8080,invalid,8082"})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("invalid"))
		})

		It("handles empty string", func() {
			var args struct {
				Ports []int `arg:"ports"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-ports="})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Ports).To(Equal([]int{}))
			Expect(len(args.Ports)).To(Equal(0))
		})
	})

	Context("[]int64", func() {
		It("parses comma-separated int64 values", func() {
			var args struct {
				Numbers []int64 `arg:"numbers"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-numbers=1000,2000,3000"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Numbers).To(Equal([]int64{1000, 2000, 3000}))
		})
	})

	Context("[]uint", func() {
		It("parses comma-separated uint values", func() {
			var args struct {
				IDs []uint `arg:"ids"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-ids=1,2,3"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.IDs).To(Equal([]uint{1, 2, 3}))
		})
	})

	Context("[]uint64", func() {
		It("parses comma-separated uint64 values", func() {
			var args struct {
				IDs []uint64 `arg:"ids"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-ids=1,2,3"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.IDs).To(Equal([]uint64{1, 2, 3}))
		})
	})

	Context("[]float64", func() {
		It("parses comma-separated float64 values", func() {
			var args struct {
				Prices []float64 `arg:"prices"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-prices=1.5,2.75,3.99"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Prices).To(Equal([]float64{1.5, 2.75, 3.99}))
		})
	})

	Context("[]bool", func() {
		It("parses comma-separated bool values", func() {
			var args struct {
				Flags []bool `arg:"flags"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-flags=true,false,true"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Flags).To(Equal([]bool{true, false, true}))
		})
	})

	Context("custom type slices", func() {
		It("parses []Username (custom string type)", func() {
			type Username string
			var args struct {
				Users []Username `arg:"users"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-users=alice,bob"})
			Expect(err).NotTo(HaveOccurred())
			Expect(len(args.Users)).To(Equal(2))
			Expect(string(args.Users[0])).To(Equal("alice"))
			Expect(string(args.Users[1])).To(Equal("bob"))
		})

		It("parses []Username with whitespace trimming", func() {
			type Username string
			var args struct {
				Users []Username `arg:"users"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-users= alice , bob "})
			Expect(err).NotTo(HaveOccurred())
			Expect(len(args.Users)).To(Equal(2))
			Expect(string(args.Users[0])).To(Equal("alice"))
			Expect(string(args.Users[1])).To(Equal("bob"))
		})
	})

	Context("environment variables", func() {
		It("parses []string from env", func() {
			var args struct {
				Names []string `env:"NAMES"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"NAMES=alice,bob,charlie"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{"alice", "bob", "charlie"}))
		})

		It("parses []int from env", func() {
			var args struct {
				Ports []int `env:"PORTS"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"PORTS=8080,8081,8082"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Ports).To(Equal([]int{8080, 8081, 8082}))
		})

		It("uses custom separator from env", func() {
			var args struct {
				Names []string `env:"NAMES" separator:":"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"NAMES=alice:bob:charlie"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{"alice", "bob", "charlie"}))
		})

		It("trims whitespace from env values", func() {
			var args struct {
				Names []string `env:"NAMES"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"NAMES=alice, bob , charlie"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{"alice", "bob", "charlie"}))
		})

		It("handles empty env value", func() {
			var args struct {
				Names []string `env:"NAMES"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"NAMES="})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{}))
		})
	})

	Context("combined args and env", func() {
		It("args override env", func() {
			var args struct {
				Names []string `arg:"names" env:"NAMES"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"NAMES=alice,bob"})
			Expect(err).NotTo(HaveOccurred())
			err = argument.ParseArgs(ctx, &args, []string{"-names=charlie,dave"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Names).To(Equal([]string{"charlie", "dave"}))
		})
	})
})
