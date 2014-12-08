package src_test

import (

	"strings"
	. "StringCalculatorKataGo/src"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringCalculator", func() {
    var cal StringCalculator

    BeforeEach(func() {
        cal = StringCalculator{
        }
    })

    Describe("Add method", func() {
        Context("with empty string imput", func() {
            It("should return zero", func() {
				Expect(cal.Add("")).To(Equal(0))
            })
        })

        Context("With only one number", func() {
            It("should be the same number", func() {
				Expect(cal.Add("3")).To(Equal(3))
            })
        })
		Context("With two numbers", func() {
            It("should be the add of numbers", func() {
				Expect(cal.Add("3,3")).To(Equal(6))
            })
        })
		Context("With more numbers", func() {
            It("should be the add of numbers", func() {
				Expect(cal.Add("3,3,3")).To(Equal(9))
            })
        })
		Context("With new lines between numbers", func() {
            It("should be the add of numbers", func() {
				Expect(cal.Add("1/n3,3,3")).To(Equal(10))
            })
        })
		Context("With different delimiter starting by //[delimiter]", func() {
            It("should be the add of numbers", func() {
				Expect(cal.Add("//[;]1/n3;3;3")).To(Equal(10))
            })
        })
		Context("With numbers bigger than 1000", func() {
            It("should throw an error", func() {
				_ , err := cal.Add("3,-3")
				Expect(err).To(HaveOccurred())
            })
            It("should throw an error with only one number", func() {
				_ , err := cal.Add("-3")
				Expect(err).To(HaveOccurred())
            })
            It("should throw show in the error message negative numbers", func() {
				_ , err := cal.Add("-3,4,-5")
				Expect(strings.Contains(err.Error(),"-3")).To(Equal(true))
				Expect(strings.Contains(err.Error(),"-5")).To(Equal(true))
            })
        })
	})

})
