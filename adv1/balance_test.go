package adv1_test

import (
	"home/adv1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// "fmt"

)

var _ = Describe("Balance", func() {
	
	var bal int
	
	BeforeEach(func(){
		bal = 100
	})

	It("should hav ini val as 100" , func(){
		bal += 100 
		res := adv1.Balance(bal)
		Expect(res).To(Equal(200))
	})

	It("already added 100 but has prev. value" , func(){
		res := adv1.Balance(bal)
		Expect(res).To(Equal(100))
	})

})
