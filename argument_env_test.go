// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bborbe/argument/v2"
)

var _ = Describe("ParseEnv", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	It("parse string", func() {
		var args struct {
			Username string `env:"user"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"user=Ben"})
		Expect(err).To(BeNil())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse int", func() {
		var args struct {
			Age int `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(29))
	})
	It("return error if parse int fails", func() {
		var args struct {
			Age int `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("return error if parse int64 fails", func() {
		var args struct {
			Age int64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("parse int64", func() {
		var args struct {
			Age int64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(int64(29)))
	})
	It("parse bool true", func() {
		var args struct {
			Confirm bool `env:"confirm"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"confirm=true"})
		Expect(err).To(BeNil())
		Expect(args.Confirm).To(BeTrue())
	})
	It("parse bool false", func() {
		var args struct {
			Confirm bool `env:"confirm"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"confirm=false"})
		Expect(err).To(BeNil())
		Expect(args.Confirm).To(BeFalse())
	})
	It("returns an error if parse bool fails", func() {
		var args struct {
			Confirm bool `env:"confirm"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"confirm=hello"})
		Expect(err).NotTo(BeNil())
	})
	It("parse duration", func() {
		var args struct {
			Wait time.Duration `env:"wait"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"wait=1m"})
		Expect(err).To(BeNil())
		Expect(args.Wait).To(Equal(time.Minute))
	})
	It("parse duration days", func() {
		var args struct {
			Wait time.Duration `env:"wait"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"wait=1d"})
		Expect(err).To(BeNil())
		Expect(args.Wait).To(Equal(24 * time.Hour))
	})
	It("return an error if parse duration fails", func() {
		var args struct {
			Wait time.Duration `env:"wait"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"wait=hello"})
		Expect(err).NotTo(BeNil())
	})
	It("parse float64", func() {
		var args struct {
			Age float64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(float64(29)))
	})
	It("return error if parse float64 fails", func() {
		var args struct {
			Age float64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("parse uint", func() {
		var args struct {
			Age uint `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(uint(29)))
	})
	It("return error if parse uint fails", func() {
		var args struct {
			Age uint `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("parse uint64", func() {
		var args struct {
			Age uint64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(uint64(29)))
	})
	It("return error if parse uint64 fails", func() {
		var args struct {
			Age uint64 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("skip fields without tag", func() {
		var args struct {
			Age int
		}
		err := argument.ParseEnv(ctx, &args, []string{})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(0))
	})
	It("returns an error if type is not supported", func() {
		var args struct {
			Age interface{} `env:"age" default:"29"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})
	It("parse int32", func() {
		var args struct {
			Age int32 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=29"})
		Expect(err).To(BeNil())
		Expect(args.Age).To(Equal(int32(29)))
	})
	It("return error if parse int32 fails", func() {
		var args struct {
			Age int32 `env:"age"`
		}
		err := argument.ParseEnv(ctx, &args, []string{"age=abc"})
		Expect(err).NotTo(BeNil())
	})

	Context("Edge cases and error handling", func() {
		It("handles malformed environment format - missing key", func() {
			var args struct {
				Username string `env:"user"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"=value"})
			Expect(err).NotTo(HaveOccurred()) // Library skips malformed entries
			Expect(args.Username).To(Equal(""))
		})

		It("handles malformed environment format - no equals sign", func() {
			var args struct {
				Username string `env:"user"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"uservalue"})
			Expect(err).NotTo(HaveOccurred()) // Library skips malformed entries
			Expect(args.Username).To(Equal(""))
		})

		It("handles duplicate environment variables - last one wins", func() {
			var args struct {
				Username string `env:"user"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"user=first", "user=second"})
			Expect(err).To(BeNil())
			Expect(args.Username).To(Equal("second"))
		})

		It("handles empty environment variable value", func() {
			var args struct {
				Username string `env:"user"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"user="})
			Expect(err).To(BeNil())
			Expect(args.Username).To(Equal(""))
		})

		It("handles environment variable with equals in value", func() {
			var args struct {
				Config string `env:"config"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"config=key=value"})
			Expect(err).To(BeNil())
			Expect(args.Config).To(Equal("key=value"))
		})

		It("handles negative numbers in environment variables", func() {
			var args struct {
				Temperature int     `env:"temp"`
				Balance     float64 `env:"balance"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"temp=-10", "balance=-99.99"})
			Expect(err).To(BeNil())
			Expect(args.Temperature).To(Equal(-10))
			Expect(args.Balance).To(Equal(-99.99))
		})

		It("handles boundary values for numeric types", func() {
			var args struct {
				MaxInt   int     `env:"maxint"`
				MaxUint  uint    `env:"maxuint"`
				MaxFloat float64 `env:"maxfloat"`
			}
			err := argument.ParseEnv(ctx, &args, []string{
				"maxint=2147483647",
				"maxuint=4294967295",
				"maxfloat=1.7976931348623157e+308",
			})
			Expect(err).To(BeNil())
			Expect(args.MaxInt).To(Equal(2147483647))
			Expect(args.MaxUint).To(Equal(uint(4294967295)))
			Expect(args.MaxFloat).To(Equal(1.7976931348623157e+308))
		})

		It("handles scientific notation in float64", func() {
			var args struct {
				Scientific float64 `env:"sci"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"sci=1.23e-10"})
			Expect(err).To(BeNil())
			Expect(args.Scientific).To(Equal(1.23e-10))
		})

		It("handles unicode strings in environment variables", func() {
			var args struct {
				Message string `env:"message"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"message=Hello 世界 🌍"})
			Expect(err).To(BeNil())
			Expect(args.Message).To(Equal("Hello 世界 🌍"))
		})

		It("handles duration with weeks", func() {
			var args struct {
				Period time.Duration `env:"period"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"period=2w"})
			Expect(err).To(BeNil())
			Expect(args.Period).To(Equal(2 * 7 * 24 * time.Hour))
		})

		It("handles complex duration combinations", func() {
			var args struct {
				Period time.Duration `env:"period"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"period=1d2h30m"})
			Expect(err).To(BeNil())
			Expect(args.Period).To(Equal(24*time.Hour + 2*time.Hour + 30*time.Minute))
		})

		It("returns error for overflow in integer parsing", func() {
			var args struct {
				SmallInt int32 `env:"small"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"small=9999999999999999999"})
			Expect(err).NotTo(BeNil())
		})

		It("handles case sensitivity in environment variable names", func() {
			var args struct {
				Username string `env:"USER"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"user=lowercase", "USER=uppercase"})
			Expect(err).To(BeNil())
			Expect(args.Username).To(Equal("uppercase"))
		})

		It("returns error for unsupported *float64 pointer type", func() {
			var args struct {
				Amount *float64 `env:"amount"`
			}
			err := argument.ParseEnv(ctx, &args, []string{"amount=123.45"})
			Expect(err).To(HaveOccurred()) // *float64 not supported in env parsing
			Expect(err.Error()).To(ContainSubstring("unsupported"))
		})

		// Note: The Fill error path in ParseEnv is difficult to trigger
		// because it requires JSON encoding/decoding to fail after successful
		// reflection setup, which is rare with normal struct types
	})
})
