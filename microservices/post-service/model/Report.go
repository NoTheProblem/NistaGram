package model

import (
	"github.com/google/uuid"
	"time"
)

type Report struct {
	Id uuid.UUID `json:"id"`
	PostId uuid.UUID `json:"postId"`
	IsAnswered bool `json:"isAnswered"`
	Post Post `json:"post"`
	DateReported time.Time `json:"date"`
	Penalty Penalty `json:"penalty"`
}


type Penalty int

const(
	RemoveContent Penalty = iota
	DeleteProfile
	DeclineReport
)
