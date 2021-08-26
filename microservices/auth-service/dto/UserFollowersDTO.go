package dto

type UserFollowersDTO struct {
	Username string `json:"username"`
	IsPrivate bool `json:"isPrivate"`
	IsNotifications bool `json:"isNotifications"`
}

