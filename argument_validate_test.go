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

var _ = Describe("Required", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	It("returns error message for env", func() {
		args := struct {
			Username string `required:"true" env:"abc"`
		}{
			Username: "",
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Required field empty, define env abc"))
	})
	It("returns error message for arg", func() {
		args := struct {
			Username string `required:"true" arg:"abc"`
		}{
			Username: "",
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Required field empty, define parameter abc"))
	})
	It("returns error message for env and arg", func() {
		args := struct {
			Username string `required:"true" arg:"abc" env:"abc"`
		}{
			Username: "",
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
		Expect(
			err.Error(),
		).To(Equal("Required field empty, define parameter abc or define env abc"))
	})
	It("returns no error if nothing is required", func() {
		args := struct {
			Username string
		}{
			Username: "Ben",
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required string is empty", func() {
		args := struct {
			Username string `required:"true"`
		}{
			Username: "",
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required string is not empty", func() {
		args := struct {
			Username string `required:"true"`
		}{
			Username: "Ben",
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required int is empty", func() {
		args := struct {
			Age int `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required int is not empty", func() {
		args := struct {
			Age int `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required int64 is empty", func() {
		args := struct {
			Age int64 `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required int64 is not empty", func() {
		args := struct {
			Age int64 `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required uint is empty", func() {
		args := struct {
			Age uint `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required uint is not empty", func() {
		args := struct {
			Age uint `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required uint64 is empty", func() {
		args := struct {
			Age uint64 `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required uint64 is not empty", func() {
		args := struct {
			Age uint64 `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required float64 is empty", func() {
		args := struct {
			Age float64 `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required float64 is not empty", func() {
		args := struct {
			Age float64 `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error if required *float64 is empty", func() {
		args := struct {
			Age *float64 `required:"true"`
		}{
			Age: nil,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required *float64 is not empty", func() {
		v := float64(123)
		args := struct {
			Age *float64 `required:"true"`
		}{
			Age: &v,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})

	It("returns error if required time.Duration is empty", func() {
		args := struct {
			Age time.Duration `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required time.Duration is not empty", func() {
		args := struct {
			Age time.Duration `required:"true"`
		}{
			Age: time.Minute,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns no error if if bool", func() {
		var args struct {
			Confirm bool `required:"true"`
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})
	It("returns error type is not supported", func() {
		var args struct {
			Banana interface{} `required:"true"`
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("field Banana with type <nil> is unsupported"))
	})
	It("returns errors if second field is invalid", func() {
		args := struct {
			Username string `required:"true"`
			Password string `required:"true"`
		}{
			Username: "Ben",
			Password: "",
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(BeNil())
	})
	It("returns error if required int32 is empty", func() {
		args := struct {
			Age int32 `required:"true"`
		}{
			Age: 0,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).To(HaveOccurred())
	})
	It("returns no error if required int32 is not empty", func() {
		args := struct {
			Age int32 `required:"true"`
		}{
			Age: 29,
		}
		err := argument.ValidateRequired(ctx, &args)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Custom types support", func() {
		type Username string
		type Port int
		type IsActive bool
		type Rate float64

		It("returns error if required custom string type is empty", func() {
			args := struct {
				Username Username `required:"true" arg:"user"`
			}{
				Username: "",
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Required field empty, define parameter user"))
		})

		It("returns no error if required custom string type is not empty", func() {
			args := struct {
				Username Username `required:"true"`
			}{
				Username: "testuser",
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns error if required custom int type is empty", func() {
			args := struct {
				Port Port `required:"true" env:"PORT"`
			}{
				Port: 0,
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Required field empty, define env PORT"))
		})

		It("returns no error if required custom int type is not empty", func() {
			args := struct {
				Port Port `required:"true"`
			}{
				Port: 8080,
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns no error if required custom bool type is false", func() {
			args := struct {
				IsActive IsActive `required:"true"`
			}{
				IsActive: false,
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns no error if required custom bool type is true", func() {
			args := struct {
				IsActive IsActive `required:"true"`
			}{
				IsActive: true,
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns error if required custom float64 type is empty", func() {
			args := struct {
				Rate Rate `required:"true" arg:"rate" env:"RATE"`
			}{
				Rate: 0.0,
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(
				err.Error(),
			).To(Equal("Required field empty, define parameter rate or define env RATE"))
		})

		It("returns no error if required custom float64 type is not empty", func() {
			args := struct {
				Rate Rate `required:"true"`
			}{
				Rate: 3.14,
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
		})

		It("handles multiple custom types with mixed requirements", func() {
			args := struct {
				Username Username `required:"true"`
				Port     Port     `required:"true"`
				IsActive IsActive `required:"true"`
				Rate     Rate     // not required
			}{
				Username: "testuser",
				Port:     8080,
				IsActive: false,
				Rate:     0.0, // should be fine since not required
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns error for first empty required custom field", func() {
			args := struct {
				Username Username `required:"true" arg:"user"`
				Port     Port     `required:"true" arg:"port"`
			}{
				Username: "", // empty, should fail here
				Port:     8080,
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Required field empty, define parameter user"))
		})

		It("handles negative values as valid for custom numeric types", func() {
			args := struct {
				Port Port `required:"true"`
				Rate Rate `required:"true"`
			}{
				Port: -1,
				Rate: -3.14,
			}
			err := argument.ValidateRequired(ctx, &args)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
