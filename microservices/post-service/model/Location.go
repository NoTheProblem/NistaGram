package model

import (
	"github.com/google/uuid"
)


type Location struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"locationName"`
	LocationPosts  []Post `json:"posts" gorm:"foreignKey:PostID"`
}
