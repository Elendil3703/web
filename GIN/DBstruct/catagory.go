package DBstruct

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CategoryID   uint
	CategoryName string
	CanteenID    uint
}
