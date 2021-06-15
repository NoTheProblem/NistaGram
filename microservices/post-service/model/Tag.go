package model

import (
	"github.com/google/uuid"
)


type Tag struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"tagName"`
	TaggedPosts  []Post `json:"posts" gorm:"foreignKey:PostID"`
}
