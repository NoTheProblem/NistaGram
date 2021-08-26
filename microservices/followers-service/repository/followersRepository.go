package repository

import (
	"fmt"
	"followers-service/DTO"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowRepository struct {
	DatabaseSession *neo4j.Session
}


func (u *FollowRepository) Follow(newFollow DTO.FollowRequestDTO) bool  {
	// TODO token -- pazi ovde je DTO
	session := *u.DatabaseSession
	var follower = "Private2"
	var followerIsPrivate = true

	err1 := u.ifExist(session, newFollow.FollowingUsername, newFollow.IsPrivate)
	err2 := u.ifExist(session, follower, followerIsPrivate)

	if (err1!=nil || err2!=nil){
		return false
	}

	var privacy = newFollow.IsPrivate;

	_, err := session.Run("match (u1:User),(u2:User) where u1.Username = $followerUsername and u2.Username = $followingUsername merge  (u1)-[f:follow{isPrivate:$isPrivate, notifications : FALSE}]->(u2) ",
		map[string]interface{}{"followingUsername":newFollow.FollowingUsername,"followerUsername":follower, "isPrivate":privacy})

	if err != nil {
		return true
	}
	return false

}

func (u *FollowRepository) ifExist(session neo4j.Session, following string, private bool) error {
	_, err := session.Run("merge (u:User{Username:$followingUsername, isPrivate:$followingPrivate}) return u", map[string]interface{}{"followingUsername":following,"followingPrivate":private,})
	if err != nil {
		return err
	}
	return nil

}

func (u *FollowRepository) Unfollow(following string) bool {
	// TODO token username
	session := *u.DatabaseSession
	var follower = "Slav";


	_, err := session.Run("match (u1:User{Username:$followerUsername})-[f:follow]->(u2:User{Username:$followingUsername}) detach delete f return u1, u2 ",
		map[string]interface{}{"followingUsername":following,"followerUsername":follower})

	if err != nil {
		return true

	}
	return false

}

func (u *FollowRepository) Block(following string) bool {
	// TODO token username
	// Unfollow - druga strana
	session := *u.DatabaseSession
	var follower = "Slav"

	_, err := session.Run("match (u1:User{Username:$followerUsername})-[f:follow]->(u2:User{Username:$followingUsername}) detach delete f return u1, u2 ",
		map[string]interface{}{"followingUsername":following,"followerUsername":follower})

	if err != nil {
		return true

	}


	res, err2 := session.Run("match (u1:User{Username:$followerUsername})-[f:follow]->(u2:User{Username:$followingUsername}) detach delete f return u1, u2 ",
		map[string]interface{}{"followingUsername":follower,"followerUsername":following})

	if err2 != nil {
		return true

	}


	fmt.Println(res.Next())


	res2, err3 := session.Run("match (u1:User),(u2:User) where u1.Username = $followerUsername and u2.Username = $followingUsername merge  (u1)-[b:block{}]->(u2); ",
		map[string]interface{}{"followingUsername":following,"followerUsername":follower})

	fmt.Println(res2.Next())

	if err3 != nil {
		return true
	}
	return false



}

func (u *FollowRepository) Unblock(following string) bool {
	// TODO token username
	session := *u.DatabaseSession

	var follower = "Slav";


	_, err := session.Run("match (u1:User{Username:$followerUsername})-[b:block]->(u2:User{Username:$followingUsername}) detach delete b return u1, u2 ",
		map[string]interface{}{"followingUsername":following,"followerUsername":follower})

	if err != nil {
		return true

	}
	return false

}

func (u *FollowRepository) AcceptRequest(follower string) bool {
	// TODO token username
	session := *u.DatabaseSession
	var following = "Private2";

	_, err := session.Run("match (u1:User{Username:$followerUsername})-[f:follow]->(u2:User{Username:$followingUsername}) set f.isPrivate = false  return f;",
		map[string]interface{}{"followingUsername":following,"followerUsername":follower,})

	if err != nil {
		return true

	}
	return false

}

func (u *FollowRepository) FindAllFollowingsUsername(follower string) ([]string)  {

	session := *u.DatabaseSession
	//Both public
	var followingsUsernames []string
	result, err := session.Run("match (u1:User{Username:$followerUsername, isPrivate:false})-[f:follow{isPrivate:false}]->(u:User{isPrivate:false}) return u.Username; ",
		map[string]interface{}{"followerUsername":follower})
	for result.Next() {
		Username, _ := result.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err != nil {
		return nil

	}
	//fmt.Println(followingsUsernames)

	//Follower public, following private
	result2, err2 := session.Run("match (u1:User{Username:$followerUsername, isPrivate:false})-[f:follow{isPrivate:false}]->(u:User{isPrivate:true}) return u.Username; ",
		map[string]interface{}{"followerUsername":follower})
	for result2.Next() {
		Username, _ := result2.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err2 != nil {
		return nil

	}
	//fmt.Println(followingsUsernames)

	//Follower private, following public
	result3, err3 := session.Run("match (u1:User{Username:$followerUsername, isPrivate:true})-[f:follow{isPrivate:false}]->(u:User{isPrivate:false}) return u.Username; ",
		map[string]interface{}{"followerUsername":follower})
	for result3.Next() {
		Username, _ := result3.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err3 != nil {
		return nil

	}
	//fmt.Println(followingsUsernames)


	//Follower private, following private
	var optFollowingsUsernames []string
	result4, err4 := session.Run("match (u1:User{Username:$followerUsername, isPrivate:true})-[f:follow{isPrivate:false}]->(u:User{isPrivate:true}) return u.Username; ",
		map[string]interface{}{"followerUsername":follower})
	for result4.Next() {
		Username, _ := result4.Record().GetByIndex(0).(string)
		optFollowingsUsernames = append(optFollowingsUsernames, Username)
		fmt.Println(optFollowingsUsernames)

	}
	fmt.Println(optFollowingsUsernames)
	fmt.Println(result4.Next())
	if err4 != nil {
		return nil

	}

	for _, optUsername := range optFollowingsUsernames {
		result5, err5 := session.Run("match (u1:User{Username:$optUsername})-[f:follow{isPrivate:false}]->(u2:User{Username:$followerUsername}) return u1.Username;",
			map[string]interface{}{"followerUsername":follower, "optUsername": optUsername})

		if result5.Next() {
			followingsUsernames = append(followingsUsernames, optUsername)
		}
		if err5 != nil {
			return nil
		}
	}


	fmt.Println(followingsUsernames)
	return  followingsUsernames


}

func (u *FollowRepository) FindAllFollowersUsername(following string)  ([]string) {
	session := *u.DatabaseSession


	//Both public
	var followingsUsernames []string
	result, err := session.Run("match (u1:User{isPrivate:false})-[f:follow{isPrivate:false}]->(u:User{Username:$followerUsername,isPrivate:false}) return u1.Username; ",
		map[string]interface{}{"followerUsername":following})
	for result.Next() {
		Username, _ := result.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err != nil {
		return nil

	}

	fmt.Println(followingsUsernames)

	//Follower private, following public
	result3, err3 := session.Run("match (u1:User{isPrivate:true})-[f:follow{isPrivate:false}]->(u:User{Username:$followerUsername, isPrivate:false}) return u1.Username; ",
		map[string]interface{}{"followerUsername":following})
	for result3.Next() {
		Username, _ := result3.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err3 != nil {
		return nil

	}


	fmt.Println(followingsUsernames)

	//Follower public, following private
	result2, err2 := session.Run("match (u1:User{ isPrivate:false})-[f:follow{isPrivate:false}]->(u:User{Username:$followerUsername,isPrivate:true}) return u1.Username; ",
		map[string]interface{}{"followerUsername":following})
	for result2.Next() {
		Username, _ := result2.Record().GetByIndex(0).(string)
		followingsUsernames = append(followingsUsernames, Username)

	}
	if err2 != nil {
		return nil

	}
	fmt.Println(followingsUsernames)

	//Follower private, following private
	var optFollowingsUsernames []string
	result4, err4 := session.Run("match (u1:User{isPrivate:true})-[f:follow{isPrivate:false}]->(u:User{Username:$followerUsername, isPrivate:true}) return u1.Username; ",
		map[string]interface{}{"followerUsername":following})
	for result4.Next() {
		Username, _ := result4.Record().GetByIndex(0).(string)
		optFollowingsUsernames = append(optFollowingsUsernames, Username)
		fmt.Println(optFollowingsUsernames)

	}
	fmt.Println(optFollowingsUsernames)
	fmt.Println(result4.Next())
	if err4 != nil {
		return nil

	}

	for _, optUsername := range optFollowingsUsernames {
		result5, err5 := session.Run("match (u1:User{Username:$optUsername})-[f:follow{isPrivate:false}]->(u2:User{Username:$followerUsername}) return u2.Username;",
			map[string]interface{}{"followerUsername":optUsername, "optUsername": following})

		if result5.Next() {
			followingsUsernames = append(followingsUsernames, optUsername)
		}
		if err5 != nil {
			return nil
		}
	}


	fmt.Println(followingsUsernames)



	return followingsUsernames


}

func (u *FollowRepository) TurnNotificationsForUserOn(username string) bool {
	session := *u.DatabaseSession

	var follower = "Private";


	_, err := session.Run("match (u1:User{Username:$followerUsername})-[f:follow]->(u2:User{Username:$followingUsername}) set f.notifications = true  return f;",
		map[string]interface{}{"followingUsername":username,"followerUsername":follower,})

	if err != nil {
		fmt.Println(err)
		return true

	}
	return false

}

func (u *FollowRepository) TurnNotificationsForUserOff(username string) bool {
	session := *u.DatabaseSession

	var follower = "Slav";


	_, err := session.Run("match (u1:User{Username:$followerUsername})-[f:follow]->(u2:User{Username:$followingUsername}) set f.notifications = false  return f;",
		map[string]interface{}{"followingUsername":username,"followerUsername":follower,})

	if err != nil {
		return true

	}
	return false

}

func (u *FollowRepository) FindAllFollowersWithNotificationTurnOn(follower string) ([]string) {

	session := *u.DatabaseSession

	fmt.Println(follower)

	var followingWithNot []string;

	var followersUsernames []string = u.FindAllFollowersUsername(follower);
	fmt.Println(followersUsernames)


	for _, optUsername := range followersUsernames {
		result5, err5 := session.Run("match (u1:User{Username:$optUsername})-[f:follow{notifications: true}]->(u2:User{Username:$followerUsername}) return u2.Username;",
			map[string]interface{}{"followerUsername":follower, "optUsername": optUsername})

		if result5.Next() {
			followingWithNot = append(followingWithNot, optUsername)
		}
		if err5 != nil {
			return nil
		}
	}


	return followingWithNot;


	
}



