package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ThoughtID uint
	Content   string
}
