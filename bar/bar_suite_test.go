package bar_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/multimodule-ginkgo/bar"
)

func TestBar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bar Suite")
}

var _ = It("says BAR", func() {
	Expect(bar.Bar()).To(Equal("BAR"))
})
