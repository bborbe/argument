// Copyright (c) 2025 Benjamin Borbe All rights reserved.
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

var _ = Describe("DefaultValues", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	It("default string", func() {
		var args struct {
			Username string `default:"user"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(data).To(HaveKeyWithValue("Username", "user"))
	})
	It("default int", func() {
		var args struct {
			Age int `default:"29"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		Expect(data).To(HaveKeyWithValue("Age", 29))
	})
	It("return error if parse int fails", func() {
		var args struct {
			Age int `default:"age"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).To(HaveOccurred())
		Expect(data).To(BeNil())
	})
	It("default int64", func() {
		var args struct {
			Age int64 `default:"29"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		value, ok := data["Age"]
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal(int64(29)))
	})
	It("return error if parse int64 fails", func() {
		var args struct {
			Age int64 `default:"age"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).To(HaveOccurred())
		Expect(data).To(BeNil())
	})
	It("default uint", func() {
		var args struct {
			Age uint `default:"29"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		value, ok := data["Age"]
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal(uint64(29)))
	})
	It("return error if parse uint fails", func() {
		var args struct {
			Age uint `default:"age"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).To(HaveOccurred())
		Expect(data).To(BeNil())
	})
	It("default uint64", func() {
		var args struct {
			Age uint64 `default:"29"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		value, ok := data["Age"]
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal(uint64(29)))
	})
	It("return error if parse uint64 fails", func() {
		var args struct {
			Age uint64 `default:"age"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).To(HaveOccurred())
		Expect(data).To(BeNil())
	})
	It("default int32", func() {
		var args struct {
			Age int32 `default:"29"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		value, ok := data["Age"]
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal(int32(29)))
	})
	It("return error if parse int32 fails", func() {
		var args struct {
			Age int32 `default:"age"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).To(HaveOccurred())
		Expect(data).To(BeNil())
	})
	It("default duration", func() {
		var args struct {
			Age time.Duration `default:"1h"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		value, ok := data["Age"]
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal(time.Hour))
	})
	It("default duration day", func() {
		var args struct {
			Age time.Duration `default:"7d"`
		}
		data, err := argument.DefaultValues(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
		value, ok := data["Age"]
		Expect(ok).To(BeTrue())
		Expect(value).To(Equal(7 * 24 * time.Hour))
	})

	Context("Error handling and edge cases", func() {
		It("returns error for unsupported types", func() {
			var args struct {
				Complex interface{} `default:"invalid"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(data).To(BeNil())
		})

		It("handles empty default values", func() {
			var args struct {
				Username string `default:""`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("Username", ""))
		})

		It("returns error for malformed float64 default", func() {
			var args struct {
				Amount float64 `default:"not-a-number"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(data).To(BeNil())
		})

		It("handles float64 default values", func() {
			var args struct {
				Amount float64 `default:"123.45"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("Amount", 123.45))
		})

		It("returns error for malformed bool default", func() {
			var args struct {
				Active bool `default:"maybe"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(data).To(BeNil())
		})

		It("handles bool true default", func() {
			var args struct {
				Active bool `default:"true"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("Active", true))
		})

		It("handles bool false default", func() {
			var args struct {
				Active bool `default:"false"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("Active", false))
		})

		It("returns error for malformed duration default", func() {
			var args struct {
				Timeout time.Duration `default:"invalid-duration"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(data).To(BeNil())
		})

		It("handles duration with weeks default", func() {
			var args struct {
				Period time.Duration `default:"2w"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			value, ok := data["Period"]
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal(2 * 7 * 24 * time.Hour))
		})

		It("handles complex duration default", func() {
			var args struct {
				Period time.Duration `default:"1d2h30m"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			value, ok := data["Period"]
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal(24*time.Hour + 2*time.Hour + 30*time.Minute))
		})

		It("handles negative number defaults", func() {
			var args struct {
				Count  int     `default:"-42"`
				Amount float64 `default:"-3.14"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("Count", -42))
			Expect(data).To(HaveKeyWithValue("Amount", -3.14))
		})

		It("handles boundary values in defaults", func() {
			var args struct {
				MaxInt   int     `default:"2147483647"`
				MaxFloat float64 `default:"1.7976931348623157e+308"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("MaxInt", 2147483647))
			Expect(data).To(HaveKeyWithValue("MaxFloat", 1.7976931348623157e+308))
		})

		It("handles scientific notation in float64 defaults", func() {
			var args struct {
				Scientific float64 `default:"1.23e-10"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("Scientific", 1.23e-10))
		})

		It("handles unicode strings in defaults", func() {
			var args struct {
				Message string `default:"Hello 世界 🌍"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("Message", "Hello 世界 🌍"))
		})

		It("skips fields without default tag", func() {
			var args struct {
				WithDefault    string `default:"value"`
				WithoutDefault string
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("WithDefault", "value"))
			Expect(data).NotTo(HaveKey("WithoutDefault"))
		})

		It("handles *float64 pointer types with defaults", func() {
			var args struct {
				Amount *float64 `default:"123.45"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			value, ok := data["Amount"]
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal(123.45))
		})

		It("returns error for overflow in integer defaults", func() {
			var args struct {
				SmallInt int32 `default:"9999999999999999999"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(data).To(BeNil())
		})

		It("returns error for malformed *float64 pointer default", func() {
			var args struct {
				Amount *float64 `default:"not-a-number"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(data).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("parse field Amount as *float64 failed"))
		})

		It("handles *float64 pointer defaults correctly", func() {
			var args struct {
				Amount *float64 `default:"123.45"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			value, ok := data["Amount"]
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal(123.45))
		})
	})

	Context("Custom types support", func() {
		type Username string
		type Port int
		type IsActive bool
		type Rate float64

		It("handles custom string type default", func() {
			var args struct {
				Username Username `default:"user"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("Username", "user"))
		})

		It("handles custom int type default", func() {
			var args struct {
				Port Port `default:"8080"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("Port", 8080))
		})

		It("handles custom bool type default", func() {
			var args struct {
				IsActive IsActive `default:"true"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("IsActive", true))
		})

		It("handles custom float64 type default", func() {
			var args struct {
				Rate Rate `default:"3.14"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(HaveKeyWithValue("Rate", 3.14))
		})

		It("returns error for malformed custom int type default", func() {
			var args struct {
				Port Port `default:"not-a-number"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(data).To(BeNil())
		})

		It("returns error for malformed custom bool type default", func() {
			var args struct {
				IsActive IsActive `default:"maybe"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(data).To(BeNil())
		})

		It("returns error for malformed custom float64 type default", func() {
			var args struct {
				Rate Rate `default:"not-a-number"`
			}
			data, err := argument.DefaultValues(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(data).To(BeNil())
		})
	})
})
