package handlers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPlayTesthttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PlayTesthttp Suite")
}
