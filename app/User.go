package app

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model  `json:"-"`
	Name        string     `json:"name"`
	PhoneNumber int64      `json:"phone_number"`
	CountryCode string     `json:"country_code"`
	//CreatedAt   time.Date `json:"created_at"`
	//UpdatedAt   mysql.Date `json:"updated_at"`
}
