package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRpnc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "rpnc Suite")
}
