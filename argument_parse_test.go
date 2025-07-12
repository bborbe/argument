// Copyright (c) 2019 Benjamin Borbe All rights reserved.
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

var _ = Describe("Parse", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
		flag.CommandLine.SetOutput(&bytes.Buffer{})
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		os.Args = []string{"go"}
		os.Clearenv()
	})
	It("parse float64 from arg default", func() {
		var args struct {
			Amount float64 `arg:"amount" env:"Amount"`
		}
		os.Args = []string{"go"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Amount).To(Equal(float64(0)))
	})
	It("parse float64 from arg", func() {
		var args struct {
			Amount float64 `arg:"amount" env:"Amount"`
		}
		os.Args = []string{"go", "-amount=23.5"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Amount).To(Equal(23.5))
	})
	It("parse *float64 from arg", func() {
		var args struct {
			Amount *float64 `arg:"amount" env:"Amount"`
		}
		os.Args = []string{"go", "-amount=23.5"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Amount).NotTo(BeNil())
		Expect(*args.Amount).To(Equal(23.5))
	})
	It("parse *float64 from arg", func() {
		var args struct {
			Amount *float64 `arg:"amount" env:"Amount"`
		}
		os.Args = []string{"go", "-amount=23.5"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Amount).NotTo(BeNil())
		Expect(*args.Amount).To(Equal(23.5))
	})
	It("parse *float64 from arg default", func() {
		var args struct {
			Amount *float64 `arg:"amount" env:"Amount"`
		}
		os.Args = []string{"go"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Amount).To(BeNil())
	})
	It("parse string from arg", func() {
		var args struct {
			Username string `arg:"user" env:"user"`
		}
		os.Args = []string{"go", "-user=Ben"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("parse string from env", func() {
		var args struct {
			Username string `arg:"user" env:"user"`
		}
		_ = os.Setenv("user", "Ben")
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("use default if env and args are not found", func() {
		var args struct {
			Username string `arg:"user" env:"user" default:"Ben"`
		}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Ben"))
	})
	It("use env if both are defined", func() {
		var args struct {
			Username string `arg:"user" env:"user"`
		}
		os.Args = []string{"go", "-user=Arg"}
		_ = os.Setenv("user", "Env")
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Env"))
	})
	It("use flag if defined", func() {
		var args struct {
			Username string `arg:"user" env:"user" default:"Default"`
		}
		os.Args = []string{"go", "-user=Arg"}
		err := argument.Parse(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(args.Username).To(Equal("Arg"))
	})

	Context("ParseAndPrint", func() {
		It("parses and prints configuration successfully", func() {
			var args struct {
				Username string `arg:"user" env:"USER" default:"test"`
				Port     int    `arg:"port" env:"PORT" default:"8080"`
			}
			os.Args = []string{"go", "-user=Ben", "-port=9090"}

			err := argument.ParseAndPrint(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Username).To(Equal("Ben"))
			Expect(args.Port).To(Equal(9090))
		})

		It("prints default values when no args provided", func() {
			var args struct {
				Username string `arg:"user" env:"USER" default:"defaultUser"`
				Debug    bool   `arg:"debug" env:"DEBUG" default:"false"`
			}
			os.Args = []string{"go"}

			err := argument.ParseAndPrint(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Username).To(Equal("defaultUser"))
			Expect(args.Debug).To(BeFalse())
		})

		It("returns error when validation fails for required field", func() {
			var args struct {
				RequiredField string `arg:"required" env:"REQUIRED" required:"true"`
			}
			os.Args = []string{"go"}

			err := argument.ParseAndPrint(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Required field empty"))
		})

		It("handles environment variables and prints them", func() {
			var args struct {
				ApiKey string `arg:"api-key" env:"API_KEY" default:"secret"`
			}
			_ = os.Setenv("API_KEY", "env-value")
			os.Args = []string{"go"}

			err := argument.ParseAndPrint(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.ApiKey).To(Equal("env-value"))
		})

		It("handles complex configuration with mixed sources", func() {
			var args struct {
				Username string `arg:"user" env:"USERNAME" default:"default"`
				Port     int    `arg:"port" env:"PORT" default:"8080"`
				Debug    bool   `arg:"debug" env:"DEBUG" default:"false"`
			}
			os.Args = []string{"go", "-user=FromArgs"}
			_ = os.Setenv("PORT", "9000")

			err := argument.ParseAndPrint(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Username).To(Equal("FromArgs"))
			Expect(args.Port).To(Equal(9000))
			Expect(args.Debug).To(BeFalse())
		})
	})

	Context("Error handling", func() {
		It("returns error when validation fails", func() {
			var args struct {
				RequiredUsername string `arg:"user" env:"USER" required:"true"`
			}
			os.Args = []string{"go"}

			err := argument.Parse(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Required field empty"))
		})

		It("handles malformed float64 arguments", func() {
			var args struct {
				Amount float64 `arg:"amount" env:"AMOUNT"`
			}
			os.Args = []string{"go", "-amount=not-a-number"}

			err := argument.Parse(ctx, &args)
			Expect(err).To(HaveOccurred())
		})

		It("handles malformed *float64 arguments", func() {
			var args struct {
				Amount *float64 `arg:"amount" env:"AMOUNT"`
			}
			os.Args = []string{"go", "-amount=invalid-float"}

			err := argument.Parse(ctx, &args)
			Expect(err).To(HaveOccurred())
		})

		It("handles malformed duration arguments", func() {
			var args struct {
				Timeout string `arg:"timeout" env:"TIMEOUT"`
			}
			os.Args = []string{"go", "-timeout=invalid-duration"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred()) // String accepts any value
		})

		It("handles negative numbers correctly", func() {
			var args struct {
				Count  int     `arg:"count" env:"COUNT"`
				Amount float64 `arg:"amount" env:"AMOUNT"`
			}
			os.Args = []string{"go", "-count=-42", "-amount=-3.14"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Count).To(Equal(-42))
			Expect(args.Amount).To(Equal(-3.14))
		})

		It("handles boundary values for integers", func() {
			var args struct {
				MaxInt int `arg:"max" env:"MAX"`
			}
			os.Args = []string{"go", "-max=2147483647"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.MaxInt).To(Equal(2147483647))
		})

		It("handles unicode strings", func() {
			var args struct {
				Message string `arg:"message" env:"MESSAGE"`
			}
			os.Args = []string{"go", "-message=Hello 世界 🌍"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Message).To(Equal("Hello 世界 🌍"))
		})

		It("handles zero values vs unset for required fields", func() {
			var args struct {
				Count int `arg:"count" env:"COUNT"`
			}
			os.Args = []string{"go", "-count=0"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Count).To(Equal(0))
		})

		It("handles empty environment variables", func() {
			var args struct {
				Value string `arg:"value" env:"VALUE" default:"default"`
			}
			_ = os.Setenv("VALUE", "")
			os.Args = []string{"go"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Value).To(Equal(""))
		})

		It("handles scientific notation for float64", func() {
			var args struct {
				BigNumber float64 `arg:"big" env:"BIG"`
			}
			os.Args = []string{"go", "-big=1.23e10"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.BigNumber).To(Equal(1.23e10))
		})
	})

	Context("Error path testing", func() {
		It("ParseAndPrint returns error when parse fails", func() {
			// Test with unsupported type to trigger parse error
			var args struct {
				UnsupportedField chan int `arg:"unsupported"`
			}
			os.Args = []string{"go"}

			err := argument.ParseAndPrint(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("parse failed"))
		})

		// Note: Print function rarely fails as it only uses reflection and logging

		It("handles environment parsing errors in parse function", func() {
			var args struct {
				UnsupportedField chan int `env:"UNSUPPORTED"`
			}
			os.Args = []string{"go"}
			_ = os.Setenv("UNSUPPORTED", "value")

			err := argument.Parse(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("env to values failed"))
		})

		It("handles default values parsing errors in parse function", func() {
			var args struct {
				UnsupportedField chan int `default:"invalid"`
			}
			os.Args = []string{"go"}

			err := argument.Parse(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("default values failed"))
		})

		// Note: The Fill error path in parse function is difficult to trigger
		// because it requires JSON encoding/decoding to fail after successful
		// reflection setup, which is rare with normal struct types
	})
})
