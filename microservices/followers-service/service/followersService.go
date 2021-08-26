package service

import (
	"fmt"
	"followers-service/DTO"
	"followers-service/repository"
)

type FollowService struct {
	FollowRepository *repository.FollowRepository

}

func (service *FollowService) FollowRequest (newFollow DTO.FollowRequestDTO) bool {
	fmt.Printf("Hello from service!")
	userFollowed :=  service.FollowRepository.Follow(newFollow)
	return userFollowed
}

func (service *FollowService) RemoveFollower(following string) bool {
	fmt.Printf("Hello from service!")
	userUnfollowed :=  service.FollowRepository.Unfollow(following)
	return userUnfollowed

}

func (service *FollowService) Block(following string) bool {
	fmt.Printf("Hello from service!")
	userBlock :=  service.FollowRepository.Block(following)
	return userBlock

}

func (service *FollowService) Unblock(following string) bool {
	fmt.Printf("Hello from service!")
	userUnblock :=  service.FollowRepository.Unblock(following)
	return userUnblock

}

func (service *FollowService) AcceptRequest(follower string) bool {
	fmt.Printf("Hello from service!")
	userAcceptedRequest :=  service.FollowRepository.AcceptRequest(follower)
	return userAcceptedRequest

}

func (service *FollowService) FindAllFollowing(follower string) ([]string)  {

	followingUsernames := service.FollowRepository.FindAllFollowingsUsername(follower)
	return followingUsernames
}

func (service *FollowService) FindAllFollowers(follower string) ([]string)  {
	followersUsernames := service.FollowRepository.FindAllFollowersUsername(follower)
	return followersUsernames

}

func (service *FollowService) TurnNotificationsForUserOn(username string) bool {
	userNotificationsTurnedOn := service.FollowRepository.TurnNotificationsForUserOn(username)
	return userNotificationsTurnedOn
	
}

func (service *FollowService) TurnNotificationsForUserOff(username string) bool {
	userNotificationsTurnedOff := service.FollowRepository.TurnNotificationsForUserOff(username)
	return userNotificationsTurnedOff

}

func (service *FollowService) FindAllFollowersWithNotificationTurnOn(follower string) ([]string) {

	followersUsernames := service.FollowRepository.FindAllFollowersWithNotificationTurnOn(follower)
	return followersUsernames
}