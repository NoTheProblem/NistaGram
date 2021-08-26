package model

type User struct {
	Username string `json:"username"`
	IsPrivate string `json:"isPrivate"`
	IsNotifications string `json:"isNotification"`

}