package DBstruct

import (
	"github.com/jinzhu/gorm"
)

// Comment 模型
type Comment struct {
	gorm.Model
	UserName       string
	UserID         int
	ProductID      uint
	ProductComment string `gorm:"size:1000"`
	Score          string
}
