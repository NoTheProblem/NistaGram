package model


type User struct {
	Email string `json:"email"  gorm:"unique;not null;default:null;"`
	Password string `json:"password" gorm:"not null;default:null"`
	Username string `json:"username" gorm:"unique;not null;default:null"`
	UserRole Role `json:"role"`

}
