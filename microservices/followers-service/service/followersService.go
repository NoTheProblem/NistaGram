package service

import (
	"fmt"
	"followers-service/DTO"
	"followers-service/repository"
)

type FollowService struct {
	FollowRepository *repository.FollowRepository

}

func (service *FollowService) FollowRequest (followRequest string, follower string) error {
	fmt.Printf("Hello from service!")
	userFollowed :=  service.FollowRepository.Follow(followRequest,follower)
	return userFollowed
}

func (service *FollowService) RemoveFollower(following string, follower string) error {
	fmt.Printf("Hello from service!")
	userUnfollowed :=  service.FollowRepository.Unfollow(following, follower)
	return userUnfollowed

}

func (service *FollowService) Block(following string, follower string) error {
	fmt.Printf("Hello from service!")
	userBlock :=  service.FollowRepository.Block(following, follower )
	return userBlock

}

func (service *FollowService) Unblock(following string, follower string) error {
	fmt.Printf("Hello from service!")
	userUnblock :=  service.FollowRepository.Unblock(following, follower)
	return userUnblock

}

func (service *FollowService) AcceptRequest(following string, follower string) error {
	fmt.Printf("Hello from service!")
	userAcceptedRequest :=  service.FollowRepository.AcceptRequest(follower, follower)
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

func (service *FollowService) TurnNotificationsForUserOn(username string) error {
	userNotificationsTurnedOn := service.FollowRepository.TurnNotificationsForUserOn(username)
	return userNotificationsTurnedOn
	
}

func (service *FollowService) TurnNotificationsForUserOff(username string) error {
	userNotificationsTurnedOff := service.FollowRepository.TurnNotificationsForUserOff(username)
	return userNotificationsTurnedOff

}

func (service *FollowService) FindAllFollowersWithNotificationTurnOn(follower string) ([]string) {

	followersUsernames := service.FollowRepository.FindAllFollowersWithNotificationTurnOn(follower)
	return followersUsernames
}

func (service *FollowService) AddUser(user DTO.UserDTO) error {
	return service.FollowRepository.AddUser(user)
}

func (service *FollowService) UpdateUser(user DTO.UserDTO) error {
	return service.FollowRepository.UpdateUser(user)

}

func (service *FollowService) DeleteUser(username string) error {
	return service.FollowRepository.DeleteUser(username)

}