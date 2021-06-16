package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Description string `json:"description"`
	NumberOfLikes int `json:"NumberOfLikes"`
	NumberOfDislikes int `json:"NumberOfDislikes"`
	IsAdd bool `json:"isAdd"`
	IsAlbum bool `json:"isAlbum"`
	NumberOfReaches int `json:"NumberOfReaches"`
	PostComments []Comment
	LocationID uint
	Tags []Tag `gorm:"many2many:post_tags;"`

}
