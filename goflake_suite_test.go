package goflake_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/reporters"
	"testing"
)

func TestGoflake(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("goflake_junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "goflake Suite", []Reporter{junitReporter})

}
