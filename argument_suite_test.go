// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

func TestArgument(t *testing.T) {
	format.TruncatedDiff = false
	RegisterFailHandler(Fail)
	RunSpecs(t, "Argument Suite")
}
