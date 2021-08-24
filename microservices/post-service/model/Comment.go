package model

type Comment struct {
	CommentText string `json:"text"`
	CommentDate string `json:"date"`
	CommentOwnerUsername string `json:"commentOwnerUsername"`
}
