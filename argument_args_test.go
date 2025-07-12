// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"bytes"
	"context"
	"flag"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bborbe/argument/v2"
)

var _ = Describe("ParseArgs", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		flag.CommandLine.SetOutput(&bytes.Buffer{})
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	})
	It("parse empty struct", func() {
		var args struct {
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).To(BeNil())
	})
	It("ignore private fields", func() {
		var args struct {
			private string
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).To(BeNil())
	})
	It("parse string from args parameter", func() {
		var args struct {
			Username string `arg:"user"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-user=Ben"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse string from default", func() {
		var args struct {
			Username string `arg:"user" default:"Ben"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse int", func() {
		var args struct {
			Age int `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(29))
	})
	It("returns an error if parse int fails", func() {
		var args struct {
			Age int `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=abc"})
		Expect(err).To(HaveOccurred())
	})
	It("return error if parse fails", func() {
		var args struct {
			Age int `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=abc"})
		Expect(err).To(HaveOccurred())
	})
	It("skip fields without tag", func() {
		var args struct {
			Age int
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(0))
	})
	It("default int", func() {
		var args struct {
			Age int `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(29))
	})
	It("parse int64", func() {
		var args struct {
			Age int64 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(int64(29)))
	})
	It("default int64", func() {
		var args struct {
			Age int64 `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(int64(29)))
	})
	It("parse bool true", func() {
		var args struct {
			Confirm bool `arg:"confirm"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-confirm=true"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeTrue())
	})
	It("parse bool false", func() {
		var args struct {
			Confirm bool `arg:"confirm"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-confirm=false"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeFalse())
	})
	It("default bool true", func() {
		var args struct {
			Confirm bool `arg:"confirm" default:"true"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeTrue())
	})
	It("default bool false", func() {
		var args struct {
			Confirm bool `arg:"confirm" default:"false"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Confirm).To(BeFalse())
	})
	It("returns an error if parse bool fails", func() {
		var args struct {
			Confirm bool `arg:"confirm"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-confirm=banana"})
		Expect(err).To(HaveOccurred())
	})
	It("parse duration", func() {
		var args struct {
			Wait time.Duration `arg:"wait"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-wait=1m"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Wait).To(Equal(time.Minute))
	})
	It("default duration", func() {
		var args struct {
			Wait time.Duration `arg:"wait" default:"1m"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Wait).To(Equal(time.Minute))
	})
	It("parse duration days", func() {
		var args struct {
			Wait time.Duration `arg:"wait"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-wait=7d"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Wait).To(Equal(7 * 24 * time.Hour))
	})
	It("default duration days", func() {
		var args struct {
			Wait time.Duration `arg:"wait" default:"7d"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Wait).To(Equal(7 * 24 * time.Hour))
	})
	It("parse float64", func() {
		var args struct {
			Age float64 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(float64(29)))
	})
	It("parse *float64", func() {
		var args struct {
			Age *float64 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(*args.Age).To(Equal(float64(29)))
	})
	It("default float64", func() {
		var args struct {
			Age float64 `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(float64(29)))
	})
	It("parse *float64", func() {
		var args struct {
			Age *float64 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age="})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(BeNil())
	})
	It("default *float64", func() {
		var args struct {
			Age *float64 `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(*args.Age).To(Equal(float64(29)))
	})

	It("parse uint", func() {
		var args struct {
			Age uint `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(uint(29)))
	})
	It("default uint", func() {
		var args struct {
			Age uint `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(uint(29)))
	})
	It("parse uint64", func() {
		var args struct {
			Age uint64 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(uint64(29)))
	})
	It("default uint64", func() {
		var args struct {
			Age uint64 `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(uint64(29)))
	})
	It("returns an error if type is not supported", func() {
		var args struct {
			Age interface{} `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).To(HaveOccurred())
	})
	It("parse int32", func() {
		var args struct {
			Age int32 `arg:"age"`
		}
		err := argument.ParseArgs(ctx, &args, []string{"-age=29"})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(int32(29)))
	})
	It("default int32", func() {
		var args struct {
			Age int32 `arg:"age" default:"29"`
		}
		err := argument.ParseArgs(ctx, &args, []string{})
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Age).To(Equal(int32(29)))
	})

	Context("Edge cases and error handling", func() {
		It("handles arguments with special characters", func() {
			var args struct {
				Message string `arg:"message"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-message=Hello World! @#$%"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Message).To(Equal("Hello World! @#$%"))
		})

		It("handles arguments with unicode characters", func() {
			var args struct {
				Name string `arg:"name"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-name=世界 🌍"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Name).To(Equal("世界 🌍"))
		})

		It("handles negative numbers in arguments", func() {
			var args struct {
				Temperature int     `arg:"temp"`
				Balance     float64 `arg:"balance"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-temp=-10", "-balance=-99.99"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Temperature).To(Equal(-10))
			Expect(args.Balance).To(Equal(-99.99))
		})

		It("handles boundary values for numeric types", func() {
			var args struct {
				MaxInt   int     `arg:"maxint"`
				MaxFloat float64 `arg:"maxfloat"`
			}
			err := argument.ParseArgs(ctx, &args, []string{
				"-maxint=2147483647",
				"-maxfloat=1.7976931348623157e+308",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.MaxInt).To(Equal(2147483647))
			Expect(args.MaxFloat).To(Equal(1.7976931348623157e+308))
		})

		It("handles scientific notation for float64", func() {
			var args struct {
				Scientific float64 `arg:"sci"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-sci=1.23e-10"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Scientific).To(Equal(1.23e-10))
		})

		It("returns error for overflow in integer parsing", func() {
			var args struct {
				SmallInt int32 `arg:"small"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-small=9999999999999999999"})
			Expect(err).To(HaveOccurred())
		})

		It("handles malformed duration arguments", func() {
			var args struct {
				Timeout time.Duration `arg:"timeout"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-timeout=invalid-duration"})
			Expect(err).To(HaveOccurred())
		})

		It("handles duration with weeks", func() {
			var args struct {
				Period time.Duration `arg:"period"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-period=2w"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Period).To(Equal(2 * 7 * 24 * time.Hour))
		})

		It("handles complex duration combinations", func() {
			var args struct {
				Period time.Duration `arg:"period"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-period=1d2h30m"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Period).To(Equal(24*time.Hour + 2*time.Hour + 30*time.Minute))
		})

		It("handles zero values correctly", func() {
			var args struct {
				Count  int     `arg:"count"`
				Amount float64 `arg:"amount"`
				Active bool    `arg:"active"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-count=0", "-amount=0.0", "-active=false"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Count).To(Equal(0))
			Expect(args.Amount).To(Equal(0.0))
			Expect(args.Active).To(BeFalse())
		})

		It("handles empty string arguments", func() {
			var args struct {
				EmptyString string `arg:"empty"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-empty="})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.EmptyString).To(Equal(""))
		})

		It("handles multiple boolean formats", func() {
			var args struct {
				True1  bool `arg:"true1"`
				True2  bool `arg:"true2"`
				False1 bool `arg:"false1"`
				False2 bool `arg:"false2"`
			}
			err := argument.ParseArgs(ctx, &args, []string{
				"-true1=true",
				"-true2=1",
				"-false1=false",
				"-false2=0",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.True1).To(BeTrue())
			Expect(args.True2).To(BeTrue())
			Expect(args.False1).To(BeFalse())
			Expect(args.False2).To(BeFalse())
		})

		It("returns error for malformed float64 arguments", func() {
			var args struct {
				Amount float64 `arg:"amount"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-amount=not-a-number"})
			Expect(err).To(HaveOccurred())
		})

		It("returns error for malformed *float64 arguments", func() {
			var args struct {
				Amount *float64 `arg:"amount"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-amount=invalid-float"})
			Expect(err).To(HaveOccurred())
		})

		It("handles arguments containing equals signs", func() {
			var args struct {
				Config string `arg:"config"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-config=key=value=another"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Config).To(Equal("key=value=another"))
		})

		It("handles arguments with spaces when quoted properly", func() {
			var args struct {
				Message string `arg:"message"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-message=Hello World"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Message).To(Equal("Hello World"))
		})

		It("leaves *float64 nil when argument is empty", func() {
			var args struct {
				Amount *float64 `arg:"amount"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-amount="})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Amount).To(BeNil())
		})

		It("handles case sensitive argument names", func() {
			var args struct {
				LowerCase string `arg:"lowercase"`
				UpperCase string `arg:"UPPERCASE"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-lowercase=lower", "-UPPERCASE=upper"})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.LowerCase).To(Equal("lower"))
			Expect(args.UpperCase).To(Equal("upper"))
		})

		It("handles empty duration argument value", func() {
			var args struct {
				Timeout time.Duration `arg:"timeout"`
			}
			err := argument.ParseArgs(ctx, &args, []string{"-timeout="})
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Timeout).To(Equal(time.Duration(0)))
		})

		// Note: The Fill error path in ParseArgs is difficult to trigger
		// because it requires JSON encoding/decoding to fail after successful
		// reflection setup, which is rare with normal struct types
	})
})
