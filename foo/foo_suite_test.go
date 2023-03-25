package foo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/multimodule-ginkgo/foo"
)

func TestFoo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Foo Suite")
}

var _ = It("says FOO", func() {
	Expect(foo.Foo()).To(Equal("FOO"))
})
