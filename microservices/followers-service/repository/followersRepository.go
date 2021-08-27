package repository

import (
	"fmt"
	"followers-service/DTO"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"sort"
)

type FollowRepository struct {
	DatabaseSession *neo4j.Session
}




func (u *FollowRepository) Follow(followRequest string, follower string) error  {
	session := *u.DatabaseSession

	result, err := session.Run("match (u:User{Username:$followingUsername}) return u.IsPrivate", map[string]interface{}{"followingUsername":followRequest})

	if err != nil {
		return err
	}



	var newFollowerIsPrivate bool
	if result.Next() {
		NewUserIsPrivate, _ := result.Record().GetByIndex(0).(bool)
		newFollowerIsPrivate = NewUserIsPrivate
		fmt.Println(newFollowerIsPrivate)

	}else {
		return fmt.Errorf("No user");
	}

	_, err2 := session.Run("match (u1:User),(u2:User) where u1.Username = $followerUsername and u2.Username = $followingUsername merge  (u1)-[f:follow{IsPrivate:$IsPrivate, notifications : FALSE}]->(u2) ",
		map[string]interface{}{"followingUsername":followRequest,"followerUsername":follower, "IsPrivate":newFollowerIsPrivate})

	if err2 != nil {
		return err
	}

	return nil

}

func (u *FollowRepository) ifExist(session neo4j.Session, following string, private bool) error {
	_, err := session.Run("merge (u:User{Username:$followingUsername, IsPrivate:$followingPrivate}) return u", map[string]interface{}{"followingUsername":following,"followingPrivate":private,})
	if err != nil {
		return err
	}
	return nil

}

func (u *FollowRepository) Unfollow(following string, follower string) error {

	session := *u.DatabaseSession

	_, err := session.Run("match (u1:User{Username:$followerUsername})-[f:follow]->(u2:User{Username:$followingUsername}) detach delete f return u1, u2 ",
		map[string]interface{}{"followingUsername":following,"followerUsername":follower})

	if err != nil {
		return err

	}
	return nil

}

func (u *FollowRepository) Block(following string, follower string) error {

	session := *u.DatabaseSession

	_, err := session.Run("match (u1:User{Username:$followerUsername})-[f:follow]->(u2:User{Username:$followingUsername}) detach delete f return u1, u2 ",
		map[string]interface{}{"followingUsername":following,"followerUsername":follower})

	if err != nil {
		return err

	}


	res, err2 := session.Run("match (u1:User{Username:$followerUsername})-[f:follow]->(u2:User{Username:$followingUsername}) detach delete f return u1, u2 ",
		map[string]interface{}{"followingUsername":follower,"followerUsername":following})

	if err2 != nil {
		return err2

	}


	fmt.Println(res.Next())


	res2, err3 := session.Run("match (u1:User),(u2:User) where u1.Username = $followerUsername and u2.Username = $followingUsername merge  (u1)-[b:block{}]->(u2); ",
		map[string]interface{}{"followingUsername":following,"followerUsername":follower})

	fmt.Println(res2.Next())

	if err3 != nil {
		return err
	}
	return nil



}

func (u *FollowRepository) Unblock(following string, follower string) error {

	session := *u.DatabaseSession
	_, err := session.Run("match (u1:User{Username:$followerUsername})-[b:block]->(u2:User{Username:$followingUsername}) detach delete b return u1, u2 ",
		map[string]interface{}{"followingUsername":following,"followerUsername":follower})

	if err != nil {
		return err

	}
	return nil

}

func (u *FollowRepository) AcceptRequest(following string, follower string) error {
	session := *u.DatabaseSession

	_, err := session.Run("match (u1:User{Username:$followerUsername})-[f:follow]->(u2:User{Username:$followingUsername}) set f.IsPrivate = false  return f;",
		map[string]interface{}{"followingUsername":following,"followerUsername":follower})

	if err != nil {
		return err

	}
	return nil

}

func (u *FollowRepository) FindAllFollowingsUsername(follower string) DTO.UsersListDTO  {

	session := *u.DatabaseSession
	//Both public
	var followingsUsernames []string
	result, err := session.Run("match (u1:User{Username:$followerUsername, IsPrivate:false})-[f:follow{IsPrivate:false}]->(u:User{IsPrivate:false}) return u.Username; ",
		map[string]interface{}{"followerUsername":follower})
	for result.Next() {
		Username, _ := result.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err != nil {
		return DTO.UsersListDTO{}

	}
	//fmt.Println(followingsUsernames)

	//Follower public, following private
	result2, err2 := session.Run("match (u1:User{Username:$followerUsername, IsPrivate:false})-[f:follow{IsPrivate:false}]->(u:User{IsPrivate:true}) return u.Username; ",
		map[string]interface{}{"followerUsername":follower})
	for result2.Next() {
		Username, _ := result2.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err2 != nil {
		return DTO.UsersListDTO{}

	}

	//fmt.Println(followingsUsernames)

	//Follower private, following public
	result3, err3 := session.Run("match (u1:User{Username:$followerUsername, IsPrivate:true})-[f:follow{IsPrivate:false}]->(u:User{IsPrivate:false}) return u.Username; ",
		map[string]interface{}{"followerUsername":follower})
	for result3.Next() {
		Username, _ := result3.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err3 != nil {
		return DTO.UsersListDTO{}

	}
	//fmt.Println(followingsUsernames)


	//Follower private, following private
	var optFollowingsUsernames []string
	result4, err4 := session.Run("match (u1:User{Username:$followerUsername, IsPrivate:true})-[f:follow{IsPrivate:false}]->(u:User{IsPrivate:true}) return u.Username; ",
		map[string]interface{}{"followerUsername":follower})
	for result4.Next() {
		Username, _ := result4.Record().GetByIndex(0).(string)
		optFollowingsUsernames = append(optFollowingsUsernames, Username)
		fmt.Println(optFollowingsUsernames)

	}
	fmt.Println(optFollowingsUsernames)
	fmt.Println(result4.Next())
	if err4 != nil {
		return DTO.UsersListDTO{}

	}

	for _, optUsername := range optFollowingsUsernames {
		result5, err5 := session.Run("match (u1:User{Username:$optUsername})-[f:follow{IsPrivate:false}]->(u2:User{Username:$followerUsername}) return u1.Username;",
			map[string]interface{}{"followerUsername":follower, "optUsername": optUsername})

		if result5.Next() {
			followingsUsernames = append(followingsUsernames, optUsername)
		}
		if err5 != nil {
			return DTO.UsersListDTO{}
		}
	}

	fmt.Println(followingsUsernames)
	return  DTO.UsersListDTO{Usernames: followingsUsernames}


}

func (u *FollowRepository) FindAllFollowersUsername(following string)  DTO.UsersListDTO {
	session := *u.DatabaseSession


	//Both public
	var followingsUsernames []string
	result, err := session.Run("match (u1:User{IsPrivate:false})-[f:follow{IsPrivate:false}]->(u:User{Username:$followerUsername,IsPrivate:false}) return u1.Username; ",
		map[string]interface{}{"followerUsername":following})
	for result.Next() {
		Username, _ := result.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err != nil {
		return DTO.UsersListDTO{}

	}

	fmt.Println(followingsUsernames)

	//Follower private, following public
	result3, err3 := session.Run("match (u1:User{IsPrivate:true})-[f:follow{IsPrivate:false}]->(u:User{Username:$followerUsername, IsPrivate:false}) return u1.Username; ",
		map[string]interface{}{"followerUsername":following})
	for result3.Next() {
		Username, _ := result3.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err3 != nil {
		return DTO.UsersListDTO{}

	}


	fmt.Println(followingsUsernames)

	//Follower public, following private
	result2, err2 := session.Run("match (u1:User{ IsPrivate:false})-[f:follow{IsPrivate:false}]->(u:User{Username:$followerUsername,IsPrivate:true}) return u1.Username; ",
		map[string]interface{}{"followerUsername":following})
	for result2.Next() {
		Username, _ := result2.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err2 != nil {
		return DTO.UsersListDTO{}

	}
	fmt.Println(followingsUsernames)

	//Follower private, following private
	var optFollowingsUsernames []string
	result4, err4 := session.Run("match (u1:User{IsPrivate:true})-[f:follow{IsPrivate:false}]->(u:User{Username:$followerUsername, IsPrivate:true}) return u1.Username; ",
		map[string]interface{}{"followerUsername":following})
	for result4.Next() {
		Username, _ := result4.Record().GetByIndex(0).(string)
		optFollowingsUsernames = append(optFollowingsUsernames, Username)
		fmt.Println(optFollowingsUsernames)

	}
	fmt.Println(optFollowingsUsernames)
	fmt.Println(result4.Next())
	if err4 != nil {
		return DTO.UsersListDTO{}

	}

	for _, optUsername := range optFollowingsUsernames {
		result5, err5 := session.Run("match (u1:User{Username:$optUsername})-[f:follow{IsPrivate:false}]->(u2:User{Username:$followerUsername}) return u2.Username;",
			map[string]interface{}{"followerUsername":optUsername, "optUsername": following})

		if result5.Next() {
			followingsUsernames = append(followingsUsernames, optUsername)
		}
		if err5 != nil {
			return DTO.UsersListDTO{}
		}
	}


	fmt.Println(followingsUsernames)



	return DTO.UsersListDTO{Usernames: followingsUsernames}


}

func (u *FollowRepository) TurnNotificationsForUserOn(username string) error {
	session := *u.DatabaseSession

	_, err := session.Run("match (u1:User{Username:$followerUsername}) set u1.IsNotifications = true  return u1;",
		map[string]interface{}{"followerUsername":username,})

	if err != nil {
		fmt.Println(err)
		return err

	}
	return nil

}

func (u *FollowRepository) TurnNotificationsForUserOff(username string) error {
	session := *u.DatabaseSession

	_, err := session.Run("match (u1:User{Username:$followerUsername}) set u1.IsNotifications = false  return u1;",
		map[string]interface{}{"followerUsername":username,})

	if err != nil {
		fmt.Println(err)
		return err

	}
	return nil

}

func (u *FollowRepository) FindAllFollowersWithNotificationTurnOn(follower string) DTO.UsersListDTO {

	session := *u.DatabaseSession

	fmt.Println(follower)

	var followingWithNot []string;

	var followersUsernames []string = u.FindAllFollowersUsername(follower).Usernames
	fmt.Println(followersUsernames)


	for _, optUsername := range followersUsernames {
		result5, err5 := session.Run("match (u1:User{Username:$optUsername, IsNotifications:$notification}) return u1.Username;",
			map[string]interface{}{"notification":true, "optUsername": optUsername})

		if result5.Next() {
			followingWithNot = append(followingWithNot, optUsername)
		}
		if err5 != nil {
			return DTO.UsersListDTO{}
		}
	}


	return DTO.UsersListDTO{Usernames: followersUsernames}


	
}

func (u *FollowRepository) AddUser(user DTO.UserDTO) error {
	session := *u.DatabaseSession
	_, err := session.Run("merge (u:User{Username:$followingUsername, IsPrivate:$followingPrivate, IsNotifications:$notifications})",
		map[string]interface{}{"followingUsername":user.Username,"followingPrivate":user.IsPrivate,"notifications": user.IsNotifications})
	if err != nil {
		return err
	}
	return nil

}

func (u *FollowRepository) UpdateUser(user DTO.UserDTO) error {
	fmt.Println(user)
	session := *u.DatabaseSession
	_, err := session.Run("match (u:User{Username:$followingUsername}) set u.IsPrivate = $followingPrivate , u.IsNotifications = $notifications",
		map[string]interface{}{"followingUsername":user.Username,"followingPrivate":user.IsPrivate,"notifications": user.IsNotifications})
	if err != nil {
		return err
	}
	return nil

}

func (u *FollowRepository) DeleteUser(username string) error {
	session := *u.DatabaseSession
	_, err := session.Run("match (u:User{Username:$username}) detach delete u",
		map[string]interface{}{"username":username})
	if err != nil {
		return err
	}
	return nil

}

func (u *FollowRepository) GetRecommendedProfiles(username string)   DTO.UsersListDTO {
	var followingsUsernames []string
	followingsUsernames = u.FindAllFollowingsUsername(username).Usernames

	var followingsFollowingsUsernames []string
	var optRecommendationMyFollowingsFollowings []string

	for _, optUsername := range followingsUsernames {
		optRecommendationMyFollowingsFollowings = u.FindAllFollowingsUsername(optUsername).Usernames
		for _, optUsername2 := range optRecommendationMyFollowingsFollowings {
			if optUsername2!=username{
				var k = true
				for _, item := range followingsUsernames {
					if item == optUsername2 {
						k = false
					}
				}
				if k {
					followingsFollowingsUsernames=append(followingsFollowingsUsernames, optUsername2)
				}
			}
		}
	}
	fmt.Println(followingsFollowingsUsernames)

	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range followingsFollowingsUsernames {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	fmt.Println("Jedinstveni")
	fmt.Println(list)

	type recommendStruct struct {
		username      string
		score        int
	}


	var optRecommendation []string
	var sameUsers []string
	var recommendations []recommendStruct
	var recommend recommendStruct

	for _, optUsername := range list {
		optRecommendation = u.FindAllFollowingsUsername(optUsername).Usernames
		for _, firstList := range followingsUsernames {
			for _, secondList := range optRecommendation {
				if secondList == firstList {
					sameUsers = append(sameUsers, firstList)
				}
			}
		}
		fmt.Println(optRecommendation)
		fmt.Println(followingsUsernames)
		fmt.Println("Zajednicki pratioci")
		fmt.Println(sameUsers)



		var numberOfSameFollowing = len(sameUsers)
		sameUsers = nil
		recommend.username = optUsername
		recommend.score = numberOfSameFollowing;

		recommendations = append(recommendations,recommend)

		fmt.Println(numberOfSameFollowing)
		fmt.Println(recommendations)



	}

	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].score > recommendations[j].score
	})

	fmt.Println(recommendations)

	var winners []string

	if len(recommendations) >=5 {
		recommendations = recommendations[0:5]
		for _, winner := range recommendations {
			winners = append(winners, winner.username)
		}
	} else{
		for _, winner := range recommendations {
			winners = append(winners, winner.username)
		}
	}
	fmt.Println(winners)
	return DTO.UsersListDTO{ Usernames: winners}
}

func (u *FollowRepository) GetRelationship(username string, username2 string) DTO.RelationTypeDTO {

	//It's me?
	if username == username2{
		return DTO.RelationTypeDTO{ Relation: DTO.Following}
	}

	session := *u.DatabaseSession

	//Blocking user
	result, _ := session.Run("match (u1:User{Username:$followerUsername})-[b:block]->(u2:User{Username:$followingUsername}) return b; ",
		map[string]interface{}{"followerUsername":username,"followingUsername":username2})

	if result.Next() {
		return DTO.RelationTypeDTO{ Relation: DTO.Blocking}
	}

	//Blocked user
	result, _ = session.Run("match (u1:User{Username:$followerUsername})-[b:block]->(u2:User{Username:$followingUsername}) return b; ",
		map[string]interface{}{"followerUsername":username2,"followingUsername":username})

	if result.Next() {
		return DTO.RelationTypeDTO{ Relation: DTO.Blocked}
	}

	//Following user
	result, _ = session.Run("match (u1:User{Username:$followerUsername})-[f:follow{IsPrivate:false}]->(u2:User{Username:$followingUsername}) return f; ",
		map[string]interface{}{"followerUsername":username,"followingUsername":username2})

	if result.Next() {
		return DTO.RelationTypeDTO{ Relation: DTO.Following}
	}

	//NotAccepted user
	result, _ = session.Run("match (u1:User{Username:$followerUsername})-[f:follow{IsPrivate:true}]->(u2:User{Username:$followingUsername}) return f; ",
		map[string]interface{}{"followerUsername":username,"followingUsername":username2})

	if result.Next() {
		return DTO.RelationTypeDTO{ Relation: DTO.NotAccepted}
	}

	return DTO.RelationTypeDTO{ Relation: DTO.NotFollowing}

}

func (u *FollowRepository) GetUser(username string) DTO.UserDTO {
	session := *u.DatabaseSession

	result, _ := session.Run("MATCH (n:User {Username: 'b'}) RETURN n",
		map[string]interface{}{"Username":username})

	var user DTO.UserDTO
	fmt.Println(result.Record().GetByIndex(0))
	if result.Next() {
		userRecord, _ := result.Record().GetByIndex(0).(DTO.UserDTO)
		user = userRecord
	}
	fmt.Println(user)
	return user

}

func (u *FollowRepository) GetUnavailableUsers(username string) DTO.UsersListDTO {

	session := *u.DatabaseSession
	var unavailableUsernames []string
	result, _ := session.Run("match (u1:User{Username:$followerUsername})-[b:block]->(u) return u.Username; ",
		map[string]interface{}{"followerUsername":username})
	for result.Next() {
		Username, _ := result.Record().GetByIndex(0).(string)
		unavailableUsernames = append(unavailableUsernames, Username)
	}

	result, _ = session.Run("match (u1)-[b:block]->(u:User{Username:$followerUsername}) return u1.Username; ",
		map[string]interface{}{"followerUsername":username})
	for result.Next() {
		Username, _ := result.Record().GetByIndex(0).(string)
		unavailableUsernames = append(unavailableUsernames, Username)
	}

	return DTO.UsersListDTO{ Usernames: unavailableUsernames}
}



