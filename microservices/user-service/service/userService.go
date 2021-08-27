package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
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
		Taggable: &t, ReceiveMessages: &t, NumberOfFollowers: 0, NumberOfFollowing: 0, IsPrivate: &f,
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

func (service *UserService) UpdateUserPrivacy(privacyDTO dto.UserPrivacyDTO, username string, token string) error {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}
	if user.IsPrivate != &privacyDTO.IsPrivate{
		var userF dto.UserFollowerDTO
		userF.Username = user.Username
		userF.IsPrivate = privacyDTO.IsPrivate
		userF.IsNotifications = *user.ReceivePostNotifications
		updateUserFollower(userF, token)
		updatePosts(privacyDTO.IsPrivate,token)
	}
	user.IsPrivate = &privacyDTO.IsPrivate
	user.ReceiveMessages = &privacyDTO.ReceiveMessages
	user.Taggable = &privacyDTO.Taggable
	err = service.UserRepository.UpdateUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}


func (service *UserService) UpdateProfileNotification(notificationDTO dto.UserNotificationDTO, username string, token string) error {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return err
	}

	if user.ReceivePostNotifications != &notificationDTO.ReceivePostNotifications{
		var userF dto.UserFollowerDTO
		userF.Username = user.Username
		userF.IsPrivate = *user.IsPrivate
		userF.IsNotifications = notificationDTO.ReceivePostNotifications
		updateUserFollower(userF, token)
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

func (service *UserService) GetUserProfile(username string, requester string, token string) (*model.User, error) {
	user, err := service.UserRepository.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	var relationType = getRelationType(username, token)

	if relationType.Relation == model.Blocked{
		return nil, errors.New("record not found")

	}

	if *user.IsPrivate {
		if requester != "" {
			switch relationType.Relation {
			case model.Blocking:
				return nil, errors.New("user blocked")
			case model.NotAccepted:
				return nil, errors.New("request not accepted")
			case model.NotFollowing:
				return nil, errors.New("private profile, send request")
			case model.Following:
				return user, nil
			}
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


func getRelationType(username string, token string) model.RelationType {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/getRelationship/" + username, os.Getenv("FOLLOWERS_SERVICE_DOMAIN"), os.Getenv("FOLLOWERS_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://user-service:8080")
	if  token == ""{
		return model.RelationType{Relation: model.NotFollowing}
	}
	req.Header.Set("Authorization", token)
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
	}
	var relationType model.RelationType
	_ = json.NewDecoder(res.Body).Decode(&relationType)
	return relationType
}


func updateUserFollower(user dto.UserFollowerDTO, token string)  {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/updateUser", os.Getenv("FOLLOWERS_SERVICE_DOMAIN"), os.Getenv("FOLLOWERS_SERVICE_PORT"))
	usrJson, _  := json.Marshal(user)
	req, _ := http.NewRequest("PUT", requestUrl, bytes.NewBuffer(usrJson))
	req.Header.Set("Host", "http://user-service:8080")
	req.Header.Set("Authorization", token)
	client.Do(req)
}

func updatePosts(privacy bool, token string) {
	type BoolDTO struct{
		Privacy bool `json:"privacy"`
	}
	var privacyDTO BoolDTO
	privacyDTO.Privacy = privacy
	privacyJson, _  := json.Marshal(privacyDTO)
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/updatePostPrivacy", os.Getenv("POST_SERVICE_DOMAIN"), os.Getenv("POST_SERVICE_PORT"))
	req, _ := http.NewRequest("PUT", requestUrl, bytes.NewBuffer(privacyJson))
	req.Header.Set("Host", "http://user-service:8080")
	req.Header.Set("Authorization", token)
	client.Do(req)
}
