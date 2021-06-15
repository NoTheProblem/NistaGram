package service

import "post-service/repository"

type PostService struct {
	PostRepository *repository.PostRepository
}