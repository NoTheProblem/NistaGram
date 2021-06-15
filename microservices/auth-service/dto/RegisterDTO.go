package dto

type RegisterDTO struct {
	Email string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Username string `json:"username" gorm:"not null"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}
