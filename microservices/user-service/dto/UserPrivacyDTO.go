package dto

type UserPrivacyDTO struct {
	Id string `json:"id"`
	ProfilePrivacy bool `json:"profilePrivacy"`
	ReceiveMessages bool `json:"receiveMessages"`
	Taggable bool `json:"taggable"`
}

