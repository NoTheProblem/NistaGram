package dto

type RegisterDTO struct {
	Password string `json:"password" gorm:"not null"`
	Username string `json:"username" gorm:"not null"`
}
