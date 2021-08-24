package service

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"os"
	"post-service/dto"
	"post-service/model"
	"post-service/repository"
	"time"
)

type PostService struct {
	PostRepository *repository.PostRepository
}


func (service *PostService) AddPost(postDto dto.PostDTO, username string, paths []string ) (string, error) {
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


func (service *PostService) GetHomeFeed(username string) interface{}{
	publicPostsDocuments := service.PostRepository.GetHomeFeedPublic()

	publicPosts := CreatePostsFromDocuments(publicPostsDocuments)
	for i, s:= range publicPosts{
		for j, _ := range s.Path {
			b, err := ioutil.ReadFile(s.Path[j])
			if err != nil {
				fmt.Print(err)
			}
			var image model.PostImages
			image.Image = b
			publicPosts[i].Images = append(publicPosts[i].Images, image)
		}
	}

	return  publicPosts

}

func (service *PostService) GetProfilePosts(username string) interface{} {
	publicPostsDocuments := service.PostRepository.GetProfilePosts(username)

	publicPosts := CreatePostsFromDocuments(publicPostsDocuments)
	for i, s:= range publicPosts{
		for j, _ := range s.Path {
			b, err := ioutil.ReadFile(s.Path[j])
			if err != nil {
				fmt.Print(err)
			}
			var image model.PostImages
			image.Image = b
			publicPosts[i].Images = append(publicPosts[i].Images, image)
		}
	}

	return  publicPosts
}

func (service *PostService) GetPostByID(id string) interface{} {
	uid, err := uuid.Parse(id)
	if err != nil {
		print(err)
		return nil
	}
	postDocument := service.PostRepository.GetPost(uid)
	var post model.Post
	bsonBytes, _ := bson.Marshal(postDocument)
	_ = bson.Unmarshal(bsonBytes, &post)
	for j, _ := range post.Path {
		b, err := ioutil.ReadFile(post.Path[j])
		if err != nil {
			fmt.Print(err)
		}
		var image model.PostImages
		image.Image = b
		post.Images = append(post.Images, image)
	}
	return  post
}

func (service *PostService) CommentPost(commentDTO dto.CommentDTO, username string) error {
	fmt.Println("CommentService")
	uid, err := uuid.Parse(commentDTO.PostId)
	if err != nil {
		return err
	}
	postDocument := service.PostRepository.GetPost(uid)
	var post model.Post
	bsonBytes, _ := bson.Marshal(postDocument)
	_ = bson.Unmarshal(bsonBytes, &post)
	var comment model.Comment
	comment.CommentText = commentDTO.CommentText
	comment.CommentDate = commentDTO.CommentDate
	comment.CommentOwnerUsername = username
	post.PostComments = append(post.PostComments, comment)

	errR := service.PostRepository.AddComment(&post)
	if errR != nil{
		return errR
	}
	return  nil
}

func (service *PostService) LikePost(id string) error {
	fmt.Println("LikeService")
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	postDocument := service.PostRepository.GetPost(uid)
	var post model.Post
	bsonBytes, _ := bson.Marshal(postDocument)
	_ = bson.Unmarshal(bsonBytes, &post)
	post.NumberOfLikes = post.NumberOfLikes + 1
	errR := service.PostRepository.UpdateLikes(&post)
	if errR != nil{
		return errR
	}
	return  nil
}

func (service *PostService) DisLikePost(id string) error {
	fmt.Println("DislikeService")
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	postDocument := service.PostRepository.GetPost(uid)
	var post model.Post
	bsonBytes, _ := bson.Marshal(postDocument)
	_ = bson.Unmarshal(bsonBytes, &post)
	post.NumberOfLikes = post.NumberOfLikes - 1
	errR := service.PostRepository.UpdateLikes(&post)
	if errR != nil{
		return errR
	}
	return  nil
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

func mapPostDtoTOPost(postDTO *dto.PostDTO, username string, paths []string) (*model.Post, error) {
	var post model.Post
	post.Id, _ = uuid.NewUUID()
	post.Tags = postDTO.Tags
	post.Description = postDTO.Description
	post.Location = postDTO.Location
	post.NumberOfDislikes, post.NumberOfLikes, post.NumberOfReaches = 0 , 0 , 0
	post.IsAdd =  postDTO.IsAdd
	if len(paths) > 1 {
		post.IsAlbum = true
	}else {
		post.IsAlbum = false
	}
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
