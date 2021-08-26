package DTO

type UserDTO struct {
	Username string `json:"username"`
	IsPrivate bool `json:"isPrivate"`
	IsNotifications bool `json:"isNotifications"`
}
