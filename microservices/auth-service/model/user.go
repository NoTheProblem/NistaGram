package model

type User struct {
	Id int `json:"id"`
	// TODO email validation?
	Email string `json:"email"  gorm:"unique;not null;default:null;"`
	Password string `json:"password" gorm:"not null;default:null"`
	Username string `json:"username" gorm:"unique;not null;default:null"`
	UserRole Role `json:"role"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	PhoneNumber string `json:"phoneNumber"`
	Gender string `json:"gender"`
	DateOfBirth string `json:"birth"`
	WebSite string `json:"webSite"`
	Bio string `json:"bio"`

}
