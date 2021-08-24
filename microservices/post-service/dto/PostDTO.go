package dto

type PostDTO struct {
	Description string `json:"description"`
	IsAdd bool `json:"isAdd"`
	IsAlbum bool `json:"isAlbum"`
	Location string `json:"location"`
	Tags []string `json:"tags"`
	Image []byte `json:"image"`
}
