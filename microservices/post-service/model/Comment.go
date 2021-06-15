package model

type Comment struct {
	ID int `json:"PostId"`
	CommentText string `json:"CommentText"`
	CommentDate string `json:"CommentDate"`
	CommentPost Post `json:"posts" gorm:"foreignKey:PostID"`

}
