package service

import (
	"errors"
	"github.com/google/uuid"
	"user-service/dto"
	"user-service/model"
	"user-service/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (service *UserService) RegisterUser (dto dto.UserRegisterDTO) error {
	t := true
	f := false
	user := model.User{Id: uuid.New(), Email: dto.Email, UserRole: model.Role(dto.UserRole), Username: dto.Username,
		Taggable: &t, ReceiveMessages: &t, NumberOfFollowers: 0, NumberOfFollowing: 0, ProfilePrivacy: &f,
		NumberOfPosts: 0, Verified: &f, ReceiveMessagesNotifications: &t, ReceivePostNotifications: &f,
		ReceiveCommentNotifications: &f}
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
	user.ProfilePrivacy = &privacyDTO.ProfilePrivacy
	user.ReceiveMessages = &privacyDTO.ReceiveMessages
	user.Taggable = &privacyDTO.Taggable
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
	user.ReceiveCommentNotifications = &notificationDTO.ReceiveCommentNotifications
	user.ReceiveMessagesNotifications = &notificationDTO.ReceiveMessagesNotifications
	user.ReceivePostNotifications = &notificationDTO.ReceivePostNotifications
	err = service.UserRepository.UpdateUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) VerifyProfile(username string) error {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}
	value := true
	user.Verified = &value
	err = service.UserRepository.UpdateUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) AddFollower(username string) error {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}
	user.NumberOfFollowers = user.NumberOfFollowers + 1
	err = service.UserRepository.UpdateUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) AddFollowing(username string) error {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}
	user.NumberOfFollowing = user.NumberOfFollowing + 1
	err = service.UserRepository.UpdateUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) AddPost(username string) error {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}
	user.NumberOfPosts = user.NumberOfPosts + 1
	err = service.UserRepository.UpdateUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) GetUserProfile(username string, requester string) (*model.User, error) {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if !*user.ProfilePrivacy {
		if requester != "" {
			//TODO videti da li se prate
			return nil, errors.New("private profile, send request to follow")
		}else {
			return nil, errors.New("private profile, log in to send request")
		}
	}

	return user, nil
}

func (service *UserService) DeleteUser(username string) {
	service.UserRepository.Delete(username)

}

func (service *UserService) SearchPublicUsers(username string) interface{} {
	publicUsers, _:= service.UserRepository.GetPublicUsersByUsername(username)
	return publicUsers
	// TODO pagable?
}
