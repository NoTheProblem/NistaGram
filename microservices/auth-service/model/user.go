package model

type User struct {
	Id int `json:"id"`
	Email string `json:"email"  gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Username string `json:"username" gorm:"unique;not null"`
	UserRole Role `json:"role"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	UserGender Gender `json:"gender"`
}
