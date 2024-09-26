package model

import (
	"gorm.io/gorm"
)

type Dude struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Job  string `json:"job"  form:"job"`
}
