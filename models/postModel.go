package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	StoryTitle string
	Body       string
}
