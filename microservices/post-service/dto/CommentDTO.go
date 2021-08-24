package dto


type CommentDTO struct {
	PostId string `json:"id"`
	CommentText string `json:"text"`
	CommentDate string `json:"date"`
}
