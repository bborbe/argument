// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	It("runs with default values", func() {
		cmd := exec.Command("go", "run", "main.go")
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, "30s").Should(gexec.Exit(0))
		output := string(session.Out.Contents())
		Expect(output).To(ContainSubstring("Username:ben"))
		Expect(output).To(ContainSubstring("Password:"))
		Expect(output).To(ContainSubstring("Active:"))
		Expect(output).To(ContainSubstring("Url:"))
	})

	It("runs with custom type arguments like in Makefile", func() {
		cmd := exec.Command("go", "run", "main.go",
			"-password=1337",
			"-url=http://example.com",
			"-active=true")
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, "30s").Should(gexec.Exit(0))
		output := string(session.Out.Contents())
		Expect(output).To(ContainSubstring("Username:ben")) // default value
		Expect(output).To(ContainSubstring("Password:1337"))
		Expect(output).To(ContainSubstring("Active:"))
		Expect(output).To(ContainSubstring("Url:http://example.com"))
	})

	It("runs with custom Username type", func() {
		cmd := exec.Command("go", "run", "main.go", "-username=testuser")
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, "30s").Should(gexec.Exit(0))
		output := string(session.Out.Contents())
		Expect(output).To(ContainSubstring("Username:testuser"))
	})

	It("handles all custom types together", func() {
		cmd := exec.Command("go", "run", "main.go",
			"-username=customuser",
			"-password=secret123",
			"-url=https://test.com",
			"-active=false")
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, "30s").Should(gexec.Exit(0))
		output := string(session.Out.Contents())
		Expect(output).To(ContainSubstring("Username:customuser"))
		Expect(output).To(ContainSubstring("Password:secret123"))
		Expect(output).To(ContainSubstring("Url:https://test.com"))
	})

	It("handles boolean flag style for custom bool type", func() {
		cmd := exec.Command("go", "run", "main.go", "-active")
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, "30s").Should(gexec.Exit(0))
		output := string(session.Out.Contents())
		Expect(output).To(ContainSubstring("Active:"))
	})

	It("demonstrates custom types work end-to-end", func() {
		cmd := exec.Command("go", "run", "main.go",
			"-username=demo",
			"-password=pass",
			"-active=true",
			"-url=http://demo.com")
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(session, "30s").Should(gexec.Exit(0))
		// Verify the custom types are working by checking output format
		output := string(session.Out.Contents())
		Expect(output).To(MatchRegexp(`Username:demo`))
		Expect(output).To(MatchRegexp(`Password:pass`))
		Expect(output).To(MatchRegexp(`Url:http://demo.com`))
	})
})
