package handler

import (
	"auth-service/dto"
	"auth-service/model"
	"auth-service/service"
	"auth-service/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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


func (handler *AuthHandler) PasswordChange(res http.ResponseWriter, req *http.Request) {
	var passwordChangerDTO dto.PasswordChangerDTO
	username := util.GetUsernameFromToken(req)
	err := json.NewDecoder(req.Body).Decode(&passwordChangerDTO)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = handler.AuthService.ChangePassword(username,passwordChangerDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (handler *AuthHandler) Authorize(writer http.ResponseWriter, request *http.Request) {
	username := util.GetUsernameFromToken(request)
	role := util.GetRoleFromToken(request)
	roleDB, er := handler.AuthService.Authenticate(username)
	if er != nil{
		fmt.Println(er)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if role != int(roleDB){
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	auth := dto.AuthDTO{
		Username: username,
		Role:     role,
	}
	responseJSON, err := json.Marshal(auth)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseJSON)
	writer.Header().Set("Content-Type", "application/json")

}

func (handler *AuthHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	username := vars["username"]
	role := util.GetRoleFromToken(request)
	if model.Role(role) != model.Administrator {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	handler.AuthService.DeleteUser(username)
	writer.WriteHeader(http.StatusAccepted)
}

func (handler *AuthHandler) RegisterBusiness(writer http.ResponseWriter, request *http.Request) {
	var businessDTO dto.BusinessDTO
	err := json.NewDecoder(request.Body).Decode(&businessDTO)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.AuthService.RegisterBusiness(businessDTO)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (handler *AuthHandler) GetPendingBusinessRequests(writer http.ResponseWriter, request *http.Request) {
	role := util.GetRoleFromToken(request)
	if model.Role(role) != model.Administrator{
		http.Error(writer, "Only admins have this permission" ,http.StatusUnauthorized)
		return
	}
	writer.Header().Set("Content-Type", "application/json")

	requests, err := handler.AuthService.AuthRepository.GetAllPendingBusinessRequests()
	if err != nil{
		http.Error(writer, err.Error() ,http.StatusBadRequest)
		return
	}
	requestsJson, err := json.Marshal(requests)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(requestsJson)
	}
}

func (handler *AuthHandler) AnswerBusinessRequest(writer http.ResponseWriter, request *http.Request) {
	var answerDTO dto.BusinessRequestAnswer
	err := json.NewDecoder(request.Body).Decode(&answerDTO)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	role := util.GetRoleFromToken(request)
	if model.Role(role) != model.Administrator {
		http.Error(writer, "Only admins have this permission" ,http.StatusUnauthorized)
		return
	}
	err = handler.AuthService.AnswerBusinessRequest(answerDTO)
	if err != nil{
		http.Error(writer, err.Error() ,http.StatusBadRequest)

	}
	writer.WriteHeader(http.StatusOK)
}

