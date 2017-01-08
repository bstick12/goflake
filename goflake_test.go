package goflake_test

import (
	"encoding/base64"
	. "github.com/bstick12/goflake"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goflake", func() {

	Describe("Generator can use unique key instead of MAC address", func() {
		It("UUIDs should be generated", func() {
			instance := GoFlakeInstanceUsingUnique("value1")
			uuid := instance.GetBase64UUID()
			byteArray, err := base64.RawURLEncoding.DecodeString(uuid)
			Ω(err).Should(BeNil())
			Ω(byteArray).Should(HaveLen(15))
		})
	})

	Describe("Generator should create UUID", func() {
		It("UUIDs should be Base64", func() {
			instance := GoFlakeInstanceUsingMacAddress()
			uuid := instance.GetBase64UUID()
			byteArray, err := base64.RawURLEncoding.DecodeString(uuid)
			Ω(err).Should(BeNil())
			Ω(byteArray).Should(HaveLen(15))

		})
		It("UUID should be sortable", func() {
			instance := GoFlakeInstanceUsingMacAddress()
			uuid := instance.GetBase64UUID()
			Ω(uuid < instance.GetBase64UUID()).Should(BeTrue())
		})
	})

	Describe("Test Get First Non Local Interface", func() {
		It("Should not be null", func() {
			iface, _ := GetNonLocalInterface()
			Ω(iface.Name).Should(Not(Equal("")))
		})
	})

	Describe("Generation Benchmarks", func() {
		Measure("Time to generate 1 million should be reasonable", func(b Benchmarker) {
			runtime := b.Time("Generate 1 million", func() {
				instance := GoFlakeInstanceUsingMacAddress()
				for i := 0; i < 1000000; i++ {
					instance.GetBase64UUID()
				}
			})

			Ω(runtime.Seconds()).Should(BeNumerically("<", 10), "GetBase64UUID() should be quick.")

		}, 3)
	})
})
