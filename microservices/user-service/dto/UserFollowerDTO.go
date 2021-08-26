package dto

type UserFollowerDTO struct {
	Username string `json:"username"`
	IsPrivate bool `json:"isPrivate"`
	IsNotifications bool `json:"isNotifications"`
}

