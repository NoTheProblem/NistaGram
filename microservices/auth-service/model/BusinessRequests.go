package model


type BusinessRequests struct {
	Email string `json:"email" gorm:"not null"`
	Web string `json:"web" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Username string `json:"username" gorm:"not null"`
	Status Status `json:"status"`
}

type Status int

const(
	Pending Status = iota
	Declined
	Accepted
)

