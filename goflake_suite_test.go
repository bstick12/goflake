package goflake_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/onsi/ginkgo/reporters"
)

func TestGoflake(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("goflake_junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "goflake Suite", []Reporter{junitReporter})
}
