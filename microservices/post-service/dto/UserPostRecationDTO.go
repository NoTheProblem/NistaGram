package dto

import "post-service/model"

type UserPostReactionDTO struct {
	Username string `json:"username"`
	LikedPosts []model.Post `json:"likedPosts"`
	DislikedPosts []model.Post `json:"dislikedPosts"`
}

