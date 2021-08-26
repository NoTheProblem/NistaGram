package DTO

type FollowRequestDTO struct {
	FollowingUsername string `json:"followedUsername"`
	IsPrivate bool `json:"isPrivate"`
}