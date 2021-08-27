package dto

type UserPrivacyDTO struct {
	Id string `json:"id"`
	IsPrivate bool `json:"isPrivate"`
	ReceiveMessages bool `json:"receiveMessages"`
	Taggable bool `json:"taggable"`
}

