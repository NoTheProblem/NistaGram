package dto

type PasswordChangerDTO struct {
	PasswordOld string `json:"passwordOld"`
	PasswordNew string `json:"passwordNew"`
}
