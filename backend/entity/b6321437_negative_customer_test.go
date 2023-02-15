package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be Blank"` //ต้องไม่เป็นว่าง
	Email      string
	CustomerID string //รหัสลูกค้าขึ้นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจำนวน 7 ตัว
}
func TestNamevalidator(t *testing.T) {
	g := NewGomegaWithT(t)

	u := Customer{
		Name:       "", //ผิด
		Email:      "bxngearnx@gmail.com",
		CustomerID: "M6321437",
	}
	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name cannot be Blank"))
}
