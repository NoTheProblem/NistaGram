package model

import "github.com/google/uuid"

type UserReaction struct {
	Username string `json:"username"`
	LikedPosts []uuid.UUID `json:"likedPosts"`
	DislikedPosts []uuid.UUID `json:"dislikedPosts"`
}
