package service

import (
	"fmt"
	"post-service/dto"
	"post-service/repository"
)

type PostService struct {
	PostRepository *repository.PostRepository
	TagRepository *repository.TagRepository
	CommentRepository *repository.CommentRepository
	LocationRepository *repository.LocationRepository
}

func (s PostService) AddPost(dto dto.PostDTO) error {
	fmt.Println(dto.IsAdd)
	fmt.Println(dto.Location)
	fmt.Println(dto.Description)
	fmt.Println(dto.Tags)
	fmt.Println(dto.IsAlbum)
	return nil
}