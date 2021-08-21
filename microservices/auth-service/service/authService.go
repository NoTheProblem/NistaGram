package service

import (
	"auth-service/dto"
	"auth-service/model"
	"auth-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

func (service *AuthService) RegisterUser (dto dto.RegisterDTO) error {
	hashPw, _ := HashPassword(dto.Password)
	user := model.User{ Email: dto.Email, UserRole: 0, Name: dto.Name, Surname: dto.Surname, Password: hashPw, Username: dto.Username}
	err := service.AuthRepository.RegisterUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthService) UpdateUser(dto dto.UpdateDTO, username string) error {
	user, err := service.AuthRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}
	user.Name = dto.Name
	user.Surname = dto.Surname
	user.Username = dto.Username
	user.Email = dto.Email
	user.Gender = dto.Gender
	user.PhoneNumber = dto.PhoneNumber
	user.DateOfBirth = dto.DateOfBirth
	user.WebSite = dto.WebSite
	user.Bio = dto.Bio
	err = service.AuthRepository.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthService) FindByUsername (dto dto.LogInDTO) (*model.User, error){
	user, err := service.AuthRepository.FindUserByUsername(dto.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func (service *AuthService) ChangePassword(username string, passwords dto.PasswordChangerDTO) (*model.User, error) {
	user, err := service.AuthRepository.FindUserByUsername(username)
	if CheckPasswordHash(passwords.PasswordOld,user.Password){
		user.Password, _ = HashPassword(passwords.PasswordNew)
		err = service.AuthRepository.UpdateUser(user)
	}
	return user, err
}

func (service *AuthService) Authenticate(username string) (model.Role, error){
	user, err := service.AuthRepository.FindUserByUsername(username)
	if err != nil {
		return model.Role(0), err
	}
	return user.UserRole, nil

}

