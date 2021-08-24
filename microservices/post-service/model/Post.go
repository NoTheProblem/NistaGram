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
	PostComments []Comment `json:"comments"`
	IsPublic bool `json:"isPublic"`
	Location string `json:"location"`
	Tags []string `json:"tags"`
	Path  []string `json:"path"`
	Owner string `json:"owner"`
	Date  time.Time `json:"date"`
	Images []PostImages `json:"images"`

}

type PostImages struct {
	Image []byte
}
