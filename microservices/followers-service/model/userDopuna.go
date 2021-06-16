package model
type UserDopuna struct {
	IsVerified bool `json:"isVerified"`
	IsPublic bool `json:"isPublic"`
	IsAvailableForMessaging bool `json:"isAvailableForMessaging"`
	IsTagable bool `json:"isTagable"`
}