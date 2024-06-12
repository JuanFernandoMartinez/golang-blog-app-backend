package models

import "gorm.io/gorm"

type Thought struct {
	gorm.Model
	Title   string
	Content string
	Owner   uint
}
