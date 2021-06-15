package service

import (
	"auth-service/dto"
	"auth-service/model"
	"auth-service/repository"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

func (service *AuthService) RegisterUser (dto dto.RegisterDTO) error {
	user := model.User{ Email: dto.Email, UserRole: 0, Name: dto.Name, Surname: dto.Surname, Password: dto.Password, Username: dto.Username}
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
