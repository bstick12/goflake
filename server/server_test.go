package main_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"net/http"
	"os/exec"
)

var pathToSprocketCLI string

type UUIDS []string

var _ = Describe("Goflake Server", func() {

	BeforeSuite(func() {
		var err error
		pathToSprocketCLI, err = gexec.Build("github.com/bstick12/goflake/server")
		Ω(err).ShouldNot(HaveOccurred())
	})

	BeforeEach(func() {
		command := exec.Command(pathToSprocketCLI)
		_, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
	})

	Describe("UUID Server", func() {
		It("Assert get UUID", func() {
			resp, _ := http.Get("http://localhost:8080/ids")
			var uuids UUIDS
			json.NewDecoder(resp.Body).Decode(&uuids)
			Ω(uuids).Should(HaveLen(1))
		})

		It("Assert get n UUIDs", func() {
			resp, _ := http.Get("http://localhost:8080/ids?count=10")
			var uuids UUIDS
			json.NewDecoder(resp.Body).Decode(&uuids)
			Ω(uuids).Should(HaveLen(10))
		})
	})

	AfterEach(func() {
		gexec.Terminate()
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

})
