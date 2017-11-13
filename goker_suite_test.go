package goker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goker Suite")
}
