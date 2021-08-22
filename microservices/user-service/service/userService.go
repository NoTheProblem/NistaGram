package service

import (
	"github.com/google/uuid"
	"user-service/dto"
	"user-service/model"
	"user-service/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (service *UserService) RegisterUser (dto dto.UserRegisterDTO) error {
	user := model.User{Id: uuid.New(), Email: dto.Email, UserRole: model.Role(dto.UserRole), Username: dto.Username}
	err := service.UserRepository.RegisterUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateProfileInfo(profileDTO dto.UserEditDTO, username string) error {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}
	user.Name = profileDTO.Name
	user.Surname = profileDTO.Surname
	user.Username = profileDTO.Username
	user.Email = profileDTO.Email
	user.Gender = profileDTO.Gender
	user.PhoneNumber = profileDTO.PhoneNumber
	user.DateOfBirth = profileDTO.DateOfBirth
	user.WebSite = profileDTO.WebSite
	user.Bio = profileDTO.Bio
	err = service.UserRepository.UpdateUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateUserPrivacy(privacyDTO dto.UserPrivacyDTO, username string) error {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}
	user.ProfilePrivacy = privacyDTO.ProfilePrivacy
	user.ReceiveMessages = privacyDTO.ReceiveMessages
	user.Taggable = privacyDTO.Taggable
	err = service.UserRepository.UpdateUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateProfileNotification(notificationDTO dto.UserNotificationDTO, username string) error {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}
	user.ReceiveCommentNotifications = notificationDTO.ReceiveCommentNotifications
	user.ReceiveMessagesNotifications = notificationDTO.ReceiveMessagesNotifications
	user.ReceivePostNotifications = notificationDTO.ReceivePostNotifications
	err = service.UserRepository.UpdateUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}
