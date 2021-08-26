package service

import (
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
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

func (service *PostService) GetPostByID(id string) model.Post {
	uid, _ := uuid.Parse(id)
	postDocument := service.PostRepository.GetPost(uid)
	var post model.Post
	bsonBytes, _ := bson.Marshal(postDocument)
	_ = bson.Unmarshal(bsonBytes, &post)
	for j, _ := range post.Path {
		b, err := ioutil.ReadFile(post.Path[j])
		if err != nil {
			fmt.Println("1")
			fmt.Print(err)
		}
		var image model.PostImages
		image.Image = b
		post.Images = append(post.Images, image)
	}
	return  post
}

func (service *PostService) CommentPost(commentDTO dto.CommentDTO, username string) error {
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

func (service *PostService) ReportPost(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	var report model.Report
	report.Id, _ = uuid.NewUUID()
	report.PostId = uid
	report.DateReported = time.Now()
	report.IsAnswered = false
	id, err = service.PostRepository.AddReport(&report)
	if err != nil{
		return err
	}
	return nil
}

func (service *PostService) GetAllUnansweredReports() interface{} {
	reportsDocuments := service.PostRepository.GetUnAnsweredReports()
	reports := CreateReportsFromDocuments(reportsDocuments)
	for i, s:= range reports{
		fmt.Println(s)
		post := service.GetPostByID(s.PostId.String())
		reports[i].Post = post
	}
	return  reports
}

func (service *PostService) AnswerReport(reportDTO dto.ReportDTO, token string) error {
	uid, err := uuid.Parse(reportDTO.Id)
	if err != nil {
		return err
	}
	err = service.PostRepository.AnswerReport(uid, reportDTO.Penalty)
	if err != nil{
		return err
	}
	reportDocument := service.PostRepository.GetReportById(uid)
	var report model.Report
	bsonBytes, _ := bson.Marshal(reportDocument)
	_ = bson.Unmarshal(bsonBytes, &report)

	if model.Penalty(reportDTO.Penalty) == model.RemoveContent {
		err = service.PostRepository.DeletePost(report.PostId)
		if err != nil {
			return err
		}
	}
	if model.Penalty(reportDTO.Penalty) == model.DeleteProfile {
		postDocument := service.GetPostByID(report.PostId.String())
		var post model.Post
		bsonBytes, _ := bson.Marshal(postDocument)
		_ = bson.Unmarshal(bsonBytes, &post)
		service.PostRepository.DeleteUserPosts(post.Owner)
		sendDeleteRequests(post.Owner, token)
		// TODO send to auth and user service

	}
	return nil
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

func CreateReportsFromDocuments(ReportsDocuments []bson.D) []model.Report {
	var reports []model.Report
	for i := 0; i < len(ReportsDocuments); i++ {
		var report model.Report
		bsonBytes, _ := bson.Marshal(ReportsDocuments[i])
		_ = bson.Unmarshal(bsonBytes, &report)
		reports = append(reports, report)
	}
	return reports
}

func mapPostDtoTOPost(postDTO *dto.PostDTO, username string, paths []string) (*model.Post, error) {
	var post model.Post
	post.Id, _ = uuid.NewUUID()
	post.Tags = postDTO.Tags
	post.Description = postDTO.Description
	post.IsPublic = postDTO.IsPublic
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

func sendDeleteRequests(username string, token string)  {
	client := &http.Client{}
	// AUTH
	requestUrl := fmt.Sprintf("http://%s:%s/deleteUser/" + username, os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest(http.MethodDelete, requestUrl, nil)
	req.Header.Set("Host", "http://post-service:8080")
	req.Header.Set("Authorization", token)
	client.Do(req)

	//USER
	requestUrl = fmt.Sprintf("http://%s:%s/deleteUser/" + username, os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	req, _ = http.NewRequest(http.MethodDelete, requestUrl, nil)
	req.Header.Set("Host", "http://post-service:8080")
	req.Header.Set("Authorization", token)
	client.Do(req)
}




