package mathutils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"home/mathutils"
)

var _ = Describe("mathutils" , func(){
	Context("when adding positive numbers" , func(){
		It("adds correctly " , func(){
			Expect(mathutils.Add(2,3)).To(Equal(5));
		})
	})

	Context("when adding negative numbers" , func(){
		It("adds correctly" ,func(){
			Expect(mathutils.Add(-2,-3)).To(Equal(-5));
		})
	})

	Context("when multiplying pos / neg  numbers " , func(){
		It("Muls correctly" , func(){
			Expect(mathutils.Mul(2,3)).To(Equal(6));
			Expect(mathutils.Mul(-20,-30)).To(Equal(600));
			Expect(mathutils.Mul(200,300)).To(Equal(60000));
		})
	})
	
	Context("When zero is used" , func(){
		It("should return 0" , func(){
			Expect(mathutils.Mul(0,5)).To(Equal(0))
		})
	})

	Context("when div 0 numbers" , func(){
		It("should return err correctly", func(){
			res , err := mathutils.Div(10,0)
			Expect(res).To(Equal(-1))
			Expect(err).To(HaveOccurred())
			// err returns as an object so convert it into same data type
			Expect(err.Error()).To(Equal("can not divide num with 0"))
		})
	})

	Context("contains ele or not" , func(){
		It("should give true if exists" , func(){
			res ,_ := mathutils.Arr([]int{1,2,3,4,5} , 3);
			Expect(res).To(BeTrue())

			res , err := mathutils.Arr([]int{} , 10);
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("arr is empty"))

		})
	})


})