package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/reporters"
	"testing"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("server_junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "goflake Server Suite", []Reporter{junitReporter})
}
