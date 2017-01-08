package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/reporters"
	"testing"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("main_junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "goflake Main Suite", []Reporter{junitReporter})
}
