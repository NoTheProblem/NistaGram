package model


type User struct {
	Password string `json:"password" gorm:"not null;default:null"`
	Username string `json:"username" gorm:"unique;not null;default:null"`
	UserRole Role `json:"role"`
}
