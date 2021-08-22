package model

import (
	"github.com/google/uuid"
	"time"
)

type VerificationRequest struct {
	Id uuid.UUID `json:"id"`
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Category string `json:"category"`
	Path string `json:"path"`
	DateSubmitted time.Time `json:"dateSubmitted"`
	DateAnswered time.Time `json:"dateAnswered"`
	Answer string `json:"answer"`
	IsAnswered bool `json:"isAnswered"`

}
