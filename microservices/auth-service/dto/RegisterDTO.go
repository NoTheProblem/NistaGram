package dto

type RegisterDTO struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}
