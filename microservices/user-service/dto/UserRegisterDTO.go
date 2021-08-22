package dto

type UserRegisterDTO struct {
	Email string `json:"email"  gorm:"unique;not null;default:null;"`
	Username string `json:"username" gorm:"unique;not null;default:null"`
	UserRole int `json:"role"`
}

