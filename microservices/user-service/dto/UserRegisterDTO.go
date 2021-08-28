package dto

type UserRegisterDTO struct {
	Username string `json:"username" gorm:"unique;not null;default:null"`
	UserRole int `json:"role"`
}

