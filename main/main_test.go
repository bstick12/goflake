package main_test

import (
	"encoding/base64"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"os/exec"
)

var pathToSprocketCLI string

var _ = Describe("Main", func() {

	BeforeSuite(func() {
		var err error
		pathToSprocketCLI, err = gexec.Build("github.com/bstick12/goflake/main")
		Ω(err).ShouldNot(HaveOccurred())
	})

	Describe("Start main", func() {
		It("Assert generated UUID", func() {
			command := exec.Command(pathToSprocketCLI)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			bytes := session.Wait().Out.Contents()
			byteArray, err := base64.RawURLEncoding.DecodeString(string(bytes[:]))
			Ω(err).Should(BeNil())
			Ω(byteArray).Should(HaveLen(15))
			Eventually(session).Should(gexec.Exit())
		})
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

})
