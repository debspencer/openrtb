package openrtb

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metric", func() {
	It("should validate", func() {
		Expect((&Metric{}).Validate()).To(Equal(ErrInvalidMetric))
		Expect((&Metric{Type: "viewability", Value: 0.45}).Validate()).NotTo(HaveOccurred())
	})
})
