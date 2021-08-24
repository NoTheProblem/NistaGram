package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"user-service/dto"
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
	username, errLogged := getUsernameFromToken(request)
	if errLogged != nil{
		http.Error(writer,errLogged.Error(),http.StatusUnauthorized)
		return
	}
	err := json.NewDecoder(request.Body).Decode(&userProfileDTO)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(err)
	err = handler.UserService.UpdateProfileInfo(userProfileDTO, username)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdateNotificationSettings(writer http.ResponseWriter, request *http.Request) {
	var userNotificationDTO dto.UserNotificationDTO
	username, errLoged := getUsernameFromToken(request)
	if errLoged != nil{
		http.Error(writer,errLoged.Error(),http.StatusUnauthorized)
		return
	}
	err := json.NewDecoder(request.Body).Decode(&userNotificationDTO)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(err)
	err = handler.UserService.UpdateProfileNotification(userNotificationDTO, username)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) UpdatePrivacySettings(writer http.ResponseWriter, request *http.Request) {
	var userPrivacyDTO dto.UserPrivacyDTO
	username, errLoged := getUsernameFromToken(request)
	if errLoged != nil{
		http.Error(writer,errLoged.Error(),http.StatusUnauthorized)
		return
	}

	err := json.NewDecoder(request.Body).Decode(&userPrivacyDTO)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(err)
	err = handler.UserService.UpdateUserPrivacy(userPrivacyDTO, username)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) LoadMyProfile(writer http.ResponseWriter, request *http.Request) {
	username, errLoged := getUsernameFromToken(request)
	if errLoged != nil{
		http.Error(writer,errLoged.Error(),http.StatusUnauthorized)
		return
	}
	writer.Header().Set("Content-Type", "application/json")


	user, userErr := handler.UserService.UserRepository.FindUserByUsername(username)
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
	requester, errLoged := getUsernameFromToken(request)
	if errLoged != nil{
		requester = ""
	}
	vars := mux.Vars(request)
	username := vars["username"]
	writer.Header().Set("Content-Type", "application/json")

	user, userErr := handler.UserService.GetUserProfile(username, requester)
	if userErr != nil{
		writer.WriteHeader(http.StatusBadRequest)
		http.Error(writer,userErr.Error(),400)
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(userJson)
	}
}


func getUsernameFromToken(r *http.Request) (string, error) {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/authorize", os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://user-service:8080")
	fmt.Println( r.Header.Get("Authorization"))
	if  r.Header.Get("Authorization") == ""{
		fmt.Println("nema autorizacije")
		return "", errors.New("no logged user")
	}
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(res)
	body, err5 := ioutil.ReadAll(res.Body)
	if err5 != nil {
		log.Fatalln(err5)
	}
	//Convert the body to type string
	sb := string(body)
	username := sb[1:len(sb)-1]
	if username == ""{
		return "", errors.New("no such user")
	}
	return username, nil
}
