package handler

import (
	"encoding/json"
	"fmt"
	"followers-service/DTO"
	"followers-service/service"
	"github.com/gorilla/mux"
	"net/http"
)

type FollowHandler struct {
	FollowService *service.FollowService
}


func (handler *FollowHandler) Follow(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello from controller!")
	var followRequest DTO.FollowRequestDTO
	err := json.NewDecoder(request.Body).Decode(&followRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var res = handler.FollowService.FollowRequest(followRequest)
	fmt.Println(res)
}

func (handler *FollowHandler) RemoveFollower(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello from controller!")
	vars := mux.Vars(request)
	following := vars["following"]
	var res = handler.FollowService.RemoveFollower(following)
	fmt.Println(res)
}

func (handler *FollowHandler) Block(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello from controller!")
	vars := mux.Vars(request)
	following := vars["following"]
	var res = handler.FollowService.Block(following)
	fmt.Println(res)
}

func (handler *FollowHandler) Unblock(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello from controller!")
	vars := mux.Vars(request)
	following := vars["following"]
	var res = handler.FollowService.Unblock(following)
	fmt.Println(res)
}

func (handler *FollowHandler) AcceptRequest(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello from controller!")
	vars := mux.Vars(request)
	follower := vars["follower"]
	var res = handler.FollowService.AcceptRequest(follower)
	fmt.Println(res)
}

func (handler *FollowHandler) FindAllFollowing(writer http.ResponseWriter, request *http.Request) {
	// TODO token username + isPrivate
	var follower = "Public";
	following := handler.FollowService.FindAllFollowing(follower)
	/*if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}*/
	for _, optUsername := range following {
		fmt.Println(optUsername);
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(following)
}

func (handler *FollowHandler) FindAllFollowers(writer http.ResponseWriter, request *http.Request) {
	// TODO token username + isPrivate
	var follower = "Public";
	followers:= handler.FollowService.FindAllFollowers(follower)
	for _, optUsername := range followers {
		fmt.Println(optUsername);
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(followers)
}

func (handler *FollowHandler) TurnNotificationsForUserOn(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	username := vars["username"]
	userNotificationsTurnedOn := handler.FollowService.TurnNotificationsForUserOn(username)
	fmt.Println(userNotificationsTurnedOn)
}

func (handler *FollowHandler) TurnNotificationsForUserOff(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	username := vars["username"]
	userNotificationsTurnedOff := handler.FollowService.TurnNotificationsForUserOff(username)
	fmt.Println(userNotificationsTurnedOff)
}

func (handler *FollowHandler) FindAllFollowersWithNotificationTurnOn(writer http.ResponseWriter, request *http.Request) {
	// TODO token username + isPrivate
	var follower = "Public";
	followingNotOn := handler.FollowService.FindAllFollowersWithNotificationTurnOn(follower)
	for _, optUsername := range followingNotOn {
		fmt.Println(optUsername);
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(followingNotOn)
}




