package model

type Relationship struct {
	UsernameFollower string `json:"usernameFollower"`
	UsernameFollowing string `json:"usernameFollowing"`
	NodeId int `json:"nodeId"`
}
