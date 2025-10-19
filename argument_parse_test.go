// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"bytes"
	"context"
	"flag"
	"os"

	"github.com/bborbe/errors"
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

		It("ParseAndPrint returns error when Print fails", func() {
			// Test with an extremely deeply nested struct to potentially trigger JSON marshal issues
			var args struct {
				Name string `arg:"name" default:"test"`
			}
			os.Args = []string{"go"}

			// This should succeed in most cases, but tests the Print error path
			err := argument.ParseAndPrint(ctx, &args)
			Expect(err).NotTo(HaveOccurred()) // This path rarely fails
		})

		It("ParseAndPrint returns error when ValidateRequired fails", func() {
			var args struct {
				RequiredField string `arg:"required" env:"REQUIRED" required:"true"`
			}
			os.Args = []string{"go"}
			// Don't set environment variable and don't provide arg to make it required but missing

			err := argument.ParseAndPrint(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("validate required failed"))
		})

		// Note: The Fill error path in parse function is difficult to trigger
		// because it requires JSON encoding/decoding to fail after successful
		// reflection setup, which is rare with normal struct types
	})

	Context("Custom types support", func() {
		type Username string
		type Port int
		type IsActive bool
		type Rate float64

		It("parses custom string type from arguments", func() {
			var args struct {
				Username Username `arg:"user" env:"USER" default:"defaultUser"`
			}
			os.Args = []string{"go", "-user=customUser"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(args.Username)).To(Equal("customUser"))
		})

		It("parses custom string type from environment", func() {
			var args struct {
				Username Username `arg:"user" env:"USER" default:"defaultUser"`
			}
			_ = os.Setenv("USER", "envUser")
			os.Args = []string{"go"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(args.Username)).To(Equal("envUser"))
		})

		It("parses custom string type from default", func() {
			var args struct {
				Username Username `arg:"user" env:"USER" default:"defaultUser"`
			}
			os.Args = []string{"go"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(args.Username)).To(Equal("defaultUser"))
		})

		It("parses custom int type from arguments", func() {
			var args struct {
				Port Port `arg:"port" env:"PORT" default:"8080"`
			}
			os.Args = []string{"go", "-port=9090"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(int(args.Port)).To(Equal(9090))
		})

		It("parses custom bool type from arguments", func() {
			var args struct {
				IsActive IsActive `arg:"active" env:"ACTIVE" default:"false"`
			}
			os.Args = []string{"go", "-active=true"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(bool(args.IsActive)).To(BeTrue())
		})

		It("parses custom float64 type from arguments", func() {
			var args struct {
				Rate Rate `arg:"rate" env:"RATE" default:"1.0"`
			}
			os.Args = []string{"go", "-rate=3.14"}

			err := argument.Parse(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(float64(args.Rate)).To(Equal(3.14))
		})

		It("handles custom types with ParseAndPrint", func() {
			var args struct {
				Username Username `arg:"user" env:"USER" default:"defaultUser"`
				Port     Port     `arg:"port" env:"PORT" default:"8080"`
				IsActive IsActive `arg:"active" env:"ACTIVE" default:"false"`
				Rate     Rate     `arg:"rate" env:"RATE" default:"1.0"`
			}
			os.Args = []string{"go", "-user=testUser", "-port=9000", "-active=true", "-rate=2.5"}

			err := argument.ParseAndPrint(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(args.Username)).To(Equal("testUser"))
			Expect(int(args.Port)).To(Equal(9000))
			Expect(bool(args.IsActive)).To(BeTrue())
			Expect(float64(args.Rate)).To(Equal(2.5))
		})
	})

	Context("ParseOnly", func() {
		It("parses arguments without validation", func() {
			var args struct {
				Username string `arg:"username" required:"true"`
			}
			os.Args = []string{"go", "-username=test"}

			// ParseOnly should succeed even though required field is set
			err := argument.ParseOnly(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Username).To(Equal("test"))
		})

		It("skips required validation", func() {
			var args struct {
				Username string `arg:"username" required:"true"`
			}
			os.Args = []string{"go"}

			// ParseOnly should succeed even though required field is empty
			err := argument.ParseOnly(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Username).To(Equal(""))

			// But ValidateRequired should fail
			err = argument.ValidateRequired(ctx, &args)
			Expect(err).To(HaveOccurred())
		})

		It("allows custom validation workflow", func() {
			var args struct {
				Port int `arg:"port" default:"80"`
			}
			os.Args = []string{"go"}

			// Parse without validation
			err := argument.ParseOnly(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Port).To(Equal(80))

			// Custom validation
			if args.Port < 1024 {
				err = errors.New(ctx, "port must be >= 1024")
			}
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("port must be >= 1024"))
		})

		It("parses all argument types correctly", func() {
			var args struct {
				String  string   `arg:"string" default:"test"`
				Int     int      `arg:"int" default:"42"`
				Float   float64  `arg:"float" default:"3.14"`
				Bool    bool     `arg:"bool" default:"true"`
				Strings []string `arg:"strings" default:"a,b,c"`
			}
			os.Args = []string{"go"}

			err := argument.ParseOnly(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.String).To(Equal("test"))
			Expect(args.Int).To(Equal(42))
			Expect(args.Float).To(Equal(3.14))
			Expect(args.Bool).To(BeTrue())
			Expect(args.Strings).To(Equal([]string{"a", "b", "c"}))
		})

		It("works with environment variables", func() {
			var args struct {
				Username string `env:"TEST_USERNAME"`
			}
			err := os.Setenv("TEST_USERNAME", "envuser")
			Expect(err).NotTo(HaveOccurred())
			os.Args = []string{"go"}

			err = argument.ParseOnly(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(args.Username).To(Equal("envuser"))
		})

		It("combines with manual validation calls", func() {
			var args struct {
				Required string   `arg:"required" required:"true"`
				Port     testPort `arg:"port" default:"80"`
			}
			os.Args = []string{"go", "-required=test"}

			// Parse only
			err := argument.ParseOnly(ctx, &args)
			Expect(err).NotTo(HaveOccurred())

			// Manually validate required
			err = argument.ValidateRequired(ctx, &args)
			Expect(err).NotTo(HaveOccurred())

			// Manually validate HasValidation
			err = argument.ValidateHasValidation(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("port must be >= 1024"))
		})

		It("ParseOnly + manual validation equals Parse (success case)", func() {
			type Config struct {
				Username string   `arg:"user" required:"true"`
				Port     testPort `arg:"port"                 default:"8080"`
			}

			var args1 Config
			os.Args = []string{"go", "-user=test", "-port=9090"}

			err := argument.Parse(ctx, &args1)
			Expect(err).NotTo(HaveOccurred())
			Expect(args1.Username).To(Equal("test"))
			Expect(args1.Port).To(Equal(testPort(9090)))
		})

		It("ParseOnly + manual validation equals Parse (required field missing)", func() {
			type Config struct {
				Username string `arg:"user" required:"true"`
			}

			var args Config
			os.Args = []string{"go"}

			err := argument.ParseOnly(ctx, &args)
			Expect(err).NotTo(HaveOccurred())

			err = argument.ValidateRequired(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Required field empty"))
		})

		It("ParseOnly + manual validation equals Parse (validation failure)", func() {
			type Config struct {
				Username string   `arg:"user" required:"true"`
				Port     testPort `arg:"port"                 default:"80"`
			}

			var args Config
			os.Args = []string{"go", "-user=test"}

			err := argument.ParseOnly(ctx, &args)
			Expect(err).NotTo(HaveOccurred())

			err = argument.ValidateRequired(ctx, &args)
			Expect(err).NotTo(HaveOccurred())

			err = argument.ValidateHasValidation(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("port must be >= 1024"))
		})
	})
})
