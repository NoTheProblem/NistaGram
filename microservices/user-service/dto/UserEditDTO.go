package dto

type UserEditDTO struct {

	Id string `json:"id"`
	Email string `json:"email" `
	Username string `json:"username"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	PhoneNumber string `json:"phoneNumber"`
	Gender string `json:"gender"`
	DateOfBirth string `json:"birth"`
	WebSite string `json:"webSite"`
	Bio string `json:"bio"`

}

