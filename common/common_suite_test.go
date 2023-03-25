package common_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/multimodule-ginkgo/common"
)

func TestCommon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Common Suite")
}

var _ = It("says COMMON", func() {
	Expect(common.Common()).To(Equal("COMMON"))
})
