package models

import (
	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Name      string
	IP        string
}
