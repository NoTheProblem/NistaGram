package handler

import (
	"encoding/json"
	"fmt"
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
	username := getUsername(request)
	if username == ""{
		writer.WriteHeader(http.StatusUnauthorized)
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
	username := getUsername(request)
	if username == ""{
		writer.WriteHeader(http.StatusUnauthorized)
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
	username := getUsername(request)
	if username == ""{
		writer.WriteHeader(http.StatusUnauthorized)
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
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/authorize", os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://user-service:8080")
	req.Header.Set("Authorization", request.Header.Get("Authorization"))
	res, err2 := client.Do(req)
	fmt.Println("7")
	if err2 != nil {
		fmt.Println("8")
		fmt.Println(err2)
		fmt.Println("8")

	}
	fmt.Println(res)
	body, err5 := ioutil.ReadAll(res.Body)
	if err5 != nil {
		fmt.Println("9")
		log.Fatalln(err5)
	}
	fmt.Println("10")

	fmt.Println(body)

	//Convert the body to type string
	sb := string(body)
	fmt.Println("11")

	fmt.Println("username: " + sb)
	username := sb[1:len(sb)-1]
	writer.Header().Set("Content-Type", "application/json")

	if username == ""{
		writer.WriteHeader(http.StatusUnauthorized)
	}
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


func getUsername(r *http.Request) string{
	client := &http.Client{}
	requestUrl := fmt.Sprintf("https://%s:%s/authorize", os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "https://user-service:8080")
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return ""
	}
	body, err5 := ioutil.ReadAll(res.Body)
	if err5 != nil {
		log.Fatalln(err5)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Println("username: " + sb)
	username := sb[1:len(sb)-1]
	return username
}
