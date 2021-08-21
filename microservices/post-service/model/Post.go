package model

import (
	"time"
)

type Post struct {

	Description string `json:"description"`
	NumberOfLikes int `json:"NumberOfLikes"`
	NumberOfDislikes int `json:"NumberOfDislikes"`
	IsAdd bool `json:"isAdd"`
	IsAlbum bool `json:"isAlbum"`
	NumberOfReaches int `json:"NumberOfReaches"`
	PostComments []Comment
	Location string
	Tags []string
	Path  string
	Owner string
	Date  time.Time

}
