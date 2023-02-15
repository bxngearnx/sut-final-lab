package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be Blank"` //ต้องไม่เป็นว่าง
	Email      string
	CustomerID string `valid:"matches(string$(L|M|H) +d$(7))~รหัสลูกค้าขึ้นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจำนวน 7 ตัว"` //รหัสลูกค้าขึ้นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจำนวน 7 ตัว
}
