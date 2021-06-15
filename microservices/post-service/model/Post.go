package model

import "github.com/google/uuid"

type Post struct {
	ID uuid.UUID `json:"id"`
	Description string `json:"description"`
	NumberOfLikes int `json:"NumberOfLikes"`
	NumberOfDislikes int `json:"NumberOfDislikes"`
	IsAdd bool `json:"isAdd"`
	IsAlbum bool `json:"isAlbum"`
	NumberOfReaches int `json:"NumberOfReaches"`
	ListOfComments []Comment
	PostLocation Location
	ListOfTags []Tag
	ListOfPaths []string
}
