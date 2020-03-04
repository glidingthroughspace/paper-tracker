package collections

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Collections", func() {
	Context("MedianOf", func() {
		It("Calculates Median for an odd amount of values", func() {
			Expect(MedianOf(10, 25, 1)).Should(BeNumerically("==", 10.0))
			Expect(MedianOf(1, 10, 25)).Should(BeNumerically("==", 10.0))
			Expect(MedianOf(-1, 10, 25)).Should(BeNumerically("==", 10.0))
			Expect(MedianOf(-99, -10, -25)).Should(BeNumerically("==", -25.0))
		})

		It("Calculates Median for an even amount of values", func() {
			Expect(MedianOf(10, 25, 1, 12)).Should(BeNumerically("==", 11.0))
			Expect(MedianOf(1, 12, 11, 25)).Should(BeNumerically("==", 11.5))
			Expect(MedianOf(-1, 10, -10, 25)).Should(BeNumerically("==", 4.5))
			Expect(MedianOf(-99, -10, -25, -12)).Should(BeNumerically("==", -18.5))
		})
	})

	Context("MeanOf", func() {
		It("Calculates the correct Mean", func() {
			Expect(MeanOf(10, 25, 1)).Should(BeNumerically("==", 12.0))
			Expect(MeanOf(1, 10, 25)).Should(BeNumerically("==", 12.0))
			Expect(MeanOf(-1, 10, 25)).Should(BeNumerically("==", 34.0/3.0))
			Expect(MeanOf(-99, -10, -25)).Should(BeNumerically("==", -134.0/3.0))
		})
	})

	Context("MinOf", func() {
		It("Returns the number if only one is given", func() {
			Expect(MinOf(1)).To(Equal(1))
		})

		It("Returns the minumum of multiple numbers", func() {
			Expect(MinOf(1, 3, 4, -23, 21)).To(Equal(-23))
		})
	})

	Context("MaxOf", func() {
		It("Returns the number if only one is given", func() {
			Expect(MaxOf(1)).To(Equal(1))
		})

		It("Returns the maximum of multiple numbers", func() {
			Expect(MaxOf(1, 3, 4, -23, 21)).To(Equal(21))
		})
	})

	Context("IsOddAmountOfValues", func() {
		It("Returns false for 0 numbers", func() {
			Expect(IsOddAmountOfValues()).To(Equal(false))
		})

		It("Returns true for 1 number", func() {
			Expect(IsOddAmountOfValues(1)).To(Equal(true))
		})

		It("Returns false for 2 number", func() {
			Expect(IsOddAmountOfValues(2, 3)).To(Equal(false))
		})
	})
})
