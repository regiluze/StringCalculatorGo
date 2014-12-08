package src_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSrc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Src Suite")
}
