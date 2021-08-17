package service

import (
	"fmt"
	"followers-service/repository"
)

type FollowService struct {
	FollowRepository *repository.FollowRepository
}

func (service *FollowService) FollowRequest (follower string, following string) {
	fmt.Printf("Hello from service!")
	service.FollowRepository.RegisterUser(follower,following)
}