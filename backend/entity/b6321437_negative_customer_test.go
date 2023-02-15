package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	
)

func TestCustomervalidator(t *testing.T) {
	g := NewGomegaWithT(t)

	u := Customer{
		Name:       "chutinan", 
		Email:      "bxngearnx@gmail.com",
		CustomerID: "6321437", //ผิด
	}
	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("รหัสลูกค้าขึ้นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจำนวน 7 ตัว"))
}
