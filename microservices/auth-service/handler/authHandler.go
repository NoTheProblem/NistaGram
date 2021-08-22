package handler

import (
	"auth-service/dto"
	"auth-service/model"
	"auth-service/service"
	"auth-service/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func (handler *AuthHandler) RegisterUser (res http.ResponseWriter, req *http.Request) {
	var registerDTO dto.RegisterDTO
	err := json.NewDecoder(req.Body).Decode(&registerDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.AuthService.RegisterUser(registerDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

func(handler *AuthHandler) Login(res http.ResponseWriter, req *http.Request){
	var logInDTO dto.LogInDTO
	err := json.NewDecoder(req.Body).Decode(&logInDTO)
	if err !=nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	var user *model.User
	user, err = handler.AuthService.FindByUsername(logInDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	match := service.CheckPasswordHash(logInDTO.Password, user.Password)
	fmt.Println("Match:   ", match)
	if !service.CheckPasswordHash(logInDTO.Password, user.Password){
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := util.CreateJWT(user.Username, &user.UserRole)
	response := dto.ResponseDTO{
		Token: token,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(responseJSON)
	res.Header().Set("Content-Type", "application/json")
}

func (handler *AuthHandler) Hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Pozdrav iz kontorolera")
}

func (handler *AuthHandler) PasswordChange(res http.ResponseWriter, req *http.Request) {
	var passwordChangerDTO dto.PasswordChangerDTO
	username := util.GetUsernameFromToken(req)
	err := json.NewDecoder(req.Body).Decode(&passwordChangerDTO)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(err)
	_, err = handler.AuthService.ChangePassword(username,passwordChangerDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (handler *AuthHandler) Authorize(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request)
	username := util.GetUsernameFromToken(request)
	_, er := handler.AuthService.Authenticate(username)
	if er != nil{
		fmt.Println(er)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	responseJSON, err := json.Marshal(username)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseJSON)
	writer.Header().Set("Content-Type", "application/json")

}

