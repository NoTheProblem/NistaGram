package handler

import (
	"fmt"
	"followers-service/service"
	"github.com/gorilla/mux"
	"net/http"
)

type FollowHandler struct {
	FollowService *service.FollowService
}

func(handler *FollowHandler) Hello(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res, "Hello from controller!")
}

func (handler *FollowHandler) FollowRequest(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello from controller!")
	vars := mux.Vars(request)
	follower := vars["follower"]
	following := vars["following"]
	handler.FollowService.FollowRequest(follower, following)
}