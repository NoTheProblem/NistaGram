package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"user-service/dto"
	"user-service/model"
	"user-service/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) RegisterUser (res http.ResponseWriter, req *http.Request) {
	var registerDTO dto.UserRegisterDTO
	err := json.NewDecoder(req.Body).Decode(&registerDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.UserService.RegisterUser(registerDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

func (handler *UserHandler) UpdateProfileInfo(writer http.ResponseWriter, request *http.Request) {
	var userProfileDTO dto.UserEditDTO
	user, errLogged := getUserFromToken(request)
	if errLogged != nil{
		http.Error(writer,errLogged.Error(),http.StatusUnauthorized)
		return
	}
	err := json.NewDecoder(request.Body).Decode(&userProfileDTO)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.UpdateProfileInfo(userProfileDTO, user.Username)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateNotificationSettings(writer http.ResponseWriter, request *http.Request) {
	var userNotificationDTO dto.UserNotificationDTO
	user, errLogged := getUserFromToken(request)
	if errLogged != nil{
		http.Error(writer,errLogged.Error(),http.StatusUnauthorized)
		return
	}
	err := json.NewDecoder(request.Body).Decode(&userNotificationDTO)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.UpdateProfileNotification(userNotificationDTO, user.Username, request.Header.Get("Authorization"))
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdatePrivacySettings(writer http.ResponseWriter, request *http.Request) {
	var userPrivacyDTO dto.UserPrivacyDTO
	user, errLogged := getUserFromToken(request)
	if errLogged != nil{
		http.Error(writer,errLogged.Error(),http.StatusUnauthorized)
		return
	}

	err := json.NewDecoder(request.Body).Decode(&userPrivacyDTO)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(err)
	err = handler.UserService.UpdateUserPrivacy(userPrivacyDTO, user.Username, request.Header.Get("Authorization"))
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) LoadMyProfile(writer http.ResponseWriter, request *http.Request) {
	authUser, errLogged := getUserFromToken(request)
	if errLogged != nil{
		http.Error(writer,errLogged.Error(),http.StatusUnauthorized)
		return
	}
	writer.Header().Set("Content-Type", "application/json")

	user, userErr := handler.UserService.UserRepository.FindUserByUsername(authUser.Username)
	if userErr != nil{
		writer.WriteHeader(http.StatusBadRequest)
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(userJson)
	}
}

func (handler *UserHandler) GetUserProfile(writer http.ResponseWriter, request *http.Request) {
	requester, errLogged := getUserFromToken(request)
	if errLogged != nil{
		requester.Username = ""
	}
	vars := mux.Vars(request)
	username := vars["username"]
	writer.Header().Set("Content-Type", "application/json")

	user, userErr := handler.UserService.GetUserProfile(username, requester.Username, request.Header.Get("Authorization"))
	if userErr != nil{
		writer.WriteHeader(http.StatusBadRequest)
		http.Error(writer,userErr.Error(),http.StatusBadRequest)
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(userJson)
	}
}

func (handler *UserHandler) VerifyProfile(writer http.ResponseWriter, request *http.Request) {
	user , _ := getUserFromToken(request)
	if model.Role(user.Role) != model.Administrator {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(request)
	username := vars["username"]
	err := handler.UserService.VerifyProfile(username)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)

}

func (handler *UserHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	username := vars["username"]
	userRequester , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if model.Role(userRequester.Role) != model.Administrator {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	handler.UserService.DeleteUser(username)
	writer.WriteHeader(http.StatusAccepted)
}

func (handler *UserHandler) SearchPublicUsers(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	username := vars["username"]
	userRequester , _ := getUserFromToken(request)
	publicUsers :=handler.UserService.SearchPublicUsers(username, userRequester.Username,request.Header.Get("Authorization"))
	writer.Header().Set("Content-Type", "application/json")
	publicUsersJson, err := json.Marshal(publicUsers)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(publicUsersJson)
	}

}


func getUserFromToken(r *http.Request) (model.Auth, error) {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/authorize", os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://user-service:8080")
	fmt.Println(r.Header.Get("Authorization"))
	if  r.Header.Get("Authorization") == ""{
		return model.Auth{}, errors.New("no logged user")
	}
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
	}

	var user model.Auth
	err := json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return model.Auth{}, err
	}

	if user.Username == ""{
		return model.Auth{}, errors.New("no such user")
	}

	return user, nil
}
