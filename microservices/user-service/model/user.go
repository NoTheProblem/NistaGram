package model

import "github.com/google/uuid"

type User struct {
	Id uuid.UUID `json:"id"`
	Email *string `json:"email"  gorm:"unique;default:null"`
	Username string `json:"username" gorm:"unique;not null;default:null"`
	UserRole Role `json:"role"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	PhoneNumber string `json:"phoneNumber"`
	Gender string `json:"gender"`
	DateOfBirth string `json:"birth"`
	WebSite string `json:"webSite"`
	Bio string `json:"bio"`
	NumberOfPosts int `json:"numberOfPosts"`
	NumberOfFollowers int `json:"numberOfFollowers"`
	NumberOfFollowing int `json:"numberOfFollowing"`
	Verified *bool `json:"verified"`
	IsPrivate *bool `json:"isPrivate"`
	ReceiveMessages *bool `json:"receiveMessages"`
	Taggable *bool `json:"taggable"`
	ReceivePostNotifications *bool `json:"receivePostNotifications"`
	ReceiveCommentNotifications *bool `json:"receiveCommentNotifications"`
	ReceiveMessagesNotifications *bool `json:"receiveMessagesNotifications"`
}
