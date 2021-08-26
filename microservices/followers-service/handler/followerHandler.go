package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"followers-service/DTO"
	"followers-service/model"
	"followers-service/service"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type FollowHandler struct {
	FollowService *service.FollowService
}


func (handler *FollowHandler) Follow(writer http.ResponseWriter, request *http.Request) {

	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	username := vars["username"]
	var res = handler.FollowService.FollowRequest(username, user.Username)
	fmt.Println(res)
}

func (handler *FollowHandler) RemoveFollower(writer http.ResponseWriter, request *http.Request) {

	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(request)
	username := vars["username"]
	var res = handler.FollowService.RemoveFollower(username, user.Username)
	fmt.Println(res)
}

func (handler *FollowHandler) Block(writer http.ResponseWriter, request *http.Request) {

	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	following := vars["user"]
	var res = handler.FollowService.Block(following, user.Username)
	fmt.Println(res)
}

func (handler *FollowHandler) Unblock(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	following := vars["user"]
	var res = handler.FollowService.Unblock(following, user.Username)
	fmt.Println(res)
}

func (handler *FollowHandler) AcceptRequest(writer http.ResponseWriter, request *http.Request) {

	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	follower := vars["follower"]
	var res = handler.FollowService.AcceptRequest(follower, user.Username)
	fmt.Println(res)
}

func (handler *FollowHandler) FindAllFollowing(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	following := handler.FollowService.FindAllFollowing(user.Username)
	/*if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}*/
	for _, optUsername := range following {
		fmt.Println(optUsername)
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(following)
}

func (handler *FollowHandler) FindAllFollowers(writer http.ResponseWriter, request *http.Request) {

	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	followers:= handler.FollowService.FindAllFollowers(user.Username)
	for _, optUsername := range followers {
		fmt.Println(optUsername)
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(followers)
}

func (handler *FollowHandler) TurnNotificationsForUserOn(writer http.ResponseWriter, request *http.Request) {

	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	userNotificationsTurnedOn := handler.FollowService.TurnNotificationsForUserOn(user.Username)
	fmt.Println(userNotificationsTurnedOn)
}

func (handler *FollowHandler) TurnNotificationsForUserOff(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	userNotificationsTurnedOff := handler.FollowService.TurnNotificationsForUserOff(user.Username)
	fmt.Println(userNotificationsTurnedOff)
}

func (handler *FollowHandler) FindAllFollowersWithNotificationTurnOn(writer http.ResponseWriter, request *http.Request) {
	// TODO isPrivate
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	followingNotOn := handler.FollowService.FindAllFollowersWithNotificationTurnOn(user.Username)
	for _, optUsername := range followingNotOn {
		fmt.Println(optUsername)
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(followingNotOn)
}

func (handler *FollowHandler) AddUser(writer http.ResponseWriter, request *http.Request) {
	var userDTO DTO.UserDTO
	err := json.NewDecoder(request.Body).Decode(&userDTO)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.FollowService.AddUser(userDTO)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (handler *FollowHandler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	var userDTO DTO.UserDTO
	err = json.NewDecoder(request.Body).Decode(&userDTO)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if user.Username != userDTO.Username {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = handler.FollowService.UpdateUser(userDTO)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (handler *FollowHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if model.Role(user.Role) != model.Administrator{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(request)
	username := vars["username"]
	err = handler.FollowService.DeleteUser(username)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}


func (handler *FollowHandler) IsFollowing(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(request)
	username := vars["username"]

	isFollowing := handler.FollowService.FollowRepository.IsFollowing(user.Username, username)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(isFollowing)
}


func getUserFromToken(r *http.Request) (model.Auth, error) {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/authorize", os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://post-service:8080")
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

func (handler *FollowHandler) GetRecommendedProfiles(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	recommend:= handler.FollowService.GetRecommendedProfiles(user.Username)
	for _, optUsername := range recommend {
		fmt.Println(optUsername)
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(recommend)
}

func (handler *FollowHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	var userDTO DTO.UserDTO
	userDTO = handler.FollowService.FollowRepository.GetUser(user.Username)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(userDTO)

}


