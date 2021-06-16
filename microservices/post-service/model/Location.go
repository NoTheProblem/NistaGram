package model


type Location struct {
	ID int `json:"id"`
	Name string `json:"locationName"`
	LocationPosts  []Post

}
