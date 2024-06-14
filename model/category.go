package model

import "gorm.io/gorm"

type Category struct {
	Name        string `gorm:"index:idx_name;unique;type:varchar(256);not null" json:"name"`
	Description string `gorm:"type:varchar(256)" json:"description"`
	BaseModel   gorm.Model
}
