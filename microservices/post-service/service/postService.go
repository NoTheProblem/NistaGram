package service

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"post-service/dto"
	"post-service/model"
	"post-service/repository"
	"time"
)

type PostService struct {
	PostRepository *repository.PostRepository
}


func (service *PostService) AddPost(postDto dto.PostDTO, username string, paths string ) (string, error) {
	post, err := mapPostDtoTOPost(&postDto, username, paths)
	if err != nil {
		return "", err
	}

	postId, err1 := service.PostRepository.AddPost(post)
	if err1 != nil {
		fmt.Println(err1)
		return "", err1
	}
	fmt.Println("postId: " + postId)
	return postId, nil
}

func (service *PostService) GetAll() interface{} {
	publicPostsDocuments := service.PostRepository.GetAll()

	publicPosts := CreatePostsFromDocuments(publicPostsDocuments)
	return publicPosts
}


func CreatePostsFromDocuments(PostsDocuments []bson.D) []model.Post {
	var publicPosts []model.Post
	for i := 0; i < len(PostsDocuments); i++ {
		var post model.Post
		bsonBytes, _ := bson.Marshal(PostsDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &post)
		publicPosts = append(publicPosts, post)
	}
	return publicPosts
}

func mapPostDtoTOPost(postDTO *dto.PostDTO, username string, paths string) (*model.Post, error) {
	var post model.Post
	post.Tags = postDTO.Tags
	post.Description = postDTO.Description
	post.Location = postDTO.Location
	post.NumberOfDislikes, post.NumberOfLikes, post.NumberOfReaches = 0 , 0 , 0
	post.IsAlbum, post.IsAdd = postDTO.IsAlbum, postDTO.IsAdd
	post.Owner = username
	post.Path = paths
	post.Date = time.Now()
	return &post, nil
}


func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}
