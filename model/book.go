package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title    string
	Category int
	Author   string
}
