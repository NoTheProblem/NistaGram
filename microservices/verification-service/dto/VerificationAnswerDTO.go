package dto


type VerificationAnswerDTO struct {
	Id string `json:"id"`
	Answer string `json:"answer"`
	VerificationAnswer bool `json:"verificationAnswer"`
}

