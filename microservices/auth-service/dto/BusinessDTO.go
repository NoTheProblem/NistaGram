package dto

type BusinessDTO struct {
	Email string `json:"email" gorm:"not null"`
	Web string `json:"web" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Username string `json:"username" gorm:"not null"`

}



