package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"net/http"
	"os"
	"post-service/dto"
	"post-service/model"
	"post-service/repository"
	"strings"
	"time"
)

type PostService struct {
	PostRepository *repository.PostRepository
}

func (service *PostService) AddPost(postDto dto.PostDTO, username string, paths []string) (string, error) {
	post, err := mapPostDtoTOPost(&postDto, username, paths)
	if err != nil {
		return "", err
	}

	postId, err1 := service.PostRepository.AddPost(post)
	if err1 != nil {
		return "", err1
	}
	return postId, nil
}

func (service *PostService) GetAll() interface{} {
	publicPostsDocuments := service.PostRepository.GetAll()

	publicPosts := CreatePostsFromDocuments(publicPostsDocuments)
	return publicPosts
}

func (service *PostService) GetHomeFeed(token string) interface{} {

	homePostsDocument := service.getPostsForHomeFeed(token)
	homePosts := CreatePostsFromDocuments(homePostsDocument)
	for i, s := range homePosts {
		for j, _ := range s.Path {
			b, err := ioutil.ReadFile(s.Path[j])
			if err != nil {
				fmt.Print(err)
			}
			var image model.PostImages
			image.Image = b
			homePosts[i].Images = append(homePosts[i].Images, image)
		}
	}
	return homePosts
}

func (service *PostService) getPostsForHomeFeed(token string) []bson.D {
	if token == ""{
		fmt.Println("public wall")
		return service.PostRepository.GetHomeFeedPublic()
	}
	followingUsers := getFollowingUsers(token)
	var postsDocument []bson.D
	for _, user := range followingUsers.Usernames {
		posts := service.PostRepository.GetProfilePosts(user)
		postsDocument = append(postsDocument, posts...)
	}
	return postsDocument
}


func (service *PostService) GetProfilePosts(username string, token string) (interface{}, error) {
	publicPostsDocuments := service.PostRepository.GetProfilePosts(username)

	publicPosts := CreatePostsFromDocuments(publicPostsDocuments)
	var relationType = getRelationType(username, token)

	if relationType.Relation == model.Blocked {
		return nil, errors.New("record not found")

	}

	if len(publicPosts) > 0 {
		if publicPosts[0].IsPrivate {
			switch relationType.Relation {
			case model.Blocking:
				return nil, errors.New("user blocked")
			case model.NotAccepted:
				return nil, errors.New("request not accepted")
			case model.NotFollowing:
				return nil, errors.New("private profile, send request")
			case model.Following:
				return publicPosts, nil
			}
		}
	}

	for i, s := range publicPosts {
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

	return publicPosts, nil
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
			fmt.Print(err)
		}
		var image model.PostImages
		image.Image = b
		post.Images = append(post.Images, image)
	}
	return post
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
	if errR != nil {
		return errR
	}
	return nil
}

func (service *PostService) LikePost(id string, username string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	postDocument := service.PostRepository.GetPost(uid)
	var post model.Post
	bsonBytes, _ := bson.Marshal(postDocument)
	_ = bson.Unmarshal(bsonBytes, &post)

	if stringInSlice(username, post.UsersLiked) {
		post.NumberOfLikes = post.NumberOfLikes - 1
		post.UsersLiked = removeStringFromSLice(username, post.UsersLiked)
		service.updateUserReactions(uid, username, "unLike")
	} else {
		post.NumberOfLikes = post.NumberOfLikes + 1
		post.UsersLiked = append(post.UsersLiked, username)
		service.updateUserReactions(uid, username, "like")
	}
	errR := service.PostRepository.UpdateLikes(&post)
	if errR != nil {
		return errR
	}
	return nil
}

func (service *PostService) DisLikePost(id string, username string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	postDocument := service.PostRepository.GetPost(uid)
	var post model.Post
	bsonBytes, _ := bson.Marshal(postDocument)
	_ = bson.Unmarshal(bsonBytes, &post)

	if stringInSlice(username, post.UsersDisliked) {
		post.NumberOfDislikes = post.NumberOfDislikes - 1
		post.UsersDisliked = removeStringFromSLice(username, post.UsersDisliked)
		service.updateUserReactions(uid, username, "unDislike")

	} else {
		post.NumberOfDislikes = post.NumberOfDislikes + 1
		post.UsersDisliked = append(post.UsersDisliked, username)
		service.updateUserReactions(uid, username, "dislike")

	}
	errR := service.PostRepository.UpdateDisLikes(&post)
	if errR != nil {
		return errR
	}
	return nil
}

func (service *PostService) updateUserReactions(id uuid.UUID, username string, reactionType string) {
	var userReactions model.UserReaction
	userReactionsDocument, err := service.PostRepository.GetUserReactions(username)
	if err != nil {
		userReactions.Username = username
		service.PostRepository.CreateUserReaction(userReactions)
	} else {
		bsonBytes, _ := bson.Marshal(userReactionsDocument)
		_ = bson.Unmarshal(bsonBytes, &userReactions)
	}
	switch reactionType {
	case "like":
		userReactions.LikedPosts = append(userReactions.LikedPosts, id)
	case "unLike":
		userReactions.LikedPosts = removeUUIDFromSLice(id, userReactions.LikedPosts)
	case "dislike":
		userReactions.DislikedPosts = append(userReactions.DislikedPosts, id)
	case "unDislike":
		userReactions.LikedPosts = removeUUIDFromSLice(id, userReactions.LikedPosts)
	}
	service.PostRepository.UpdateUserReactions(userReactions)
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
	if err != nil {
		return err
	}
	return nil
}

func (service *PostService) GetAllUnansweredReports() interface{} {
	reportsDocuments := service.PostRepository.GetUnAnsweredReports()
	reports := CreateReportsFromDocuments(reportsDocuments)
	for i, s := range reports {
		post := service.GetPostByID(s.PostId.String())
		reports[i].Post = post
	}
	return reports
}

func (service *PostService) AnswerReport(reportDTO dto.ReportDTO, token string) error {
	uid, err := uuid.Parse(reportDTO.Id)
	if err != nil {
		return err
	}
	err = service.PostRepository.AnswerReport(uid, reportDTO.Penalty)
	if err != nil {
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

	}
	return nil
}

func (service *PostService) SearchLocation(location string, username string, token string) interface{} {
	publicPostsDocuments := service.PostRepository.GetPublicPosts()
	publicPosts := CreatePostsFromDocuments(publicPostsDocuments)
	var locationPosts []model.Post

	for _, p := range publicPosts {
		if strings.Contains(strings.ToLower(p.Location), strings.ToLower(location)) {
			locationPosts = append(locationPosts, p)
		}
	}

	for i, s := range locationPosts {
		for j, _ := range s.Path {
			b, err := ioutil.ReadFile(s.Path[j])
			if err != nil {
				fmt.Print(err)
			}
			var image model.PostImages
			image.Image = b
			locationPosts[i].Images = append(locationPosts[i].Images, image)
		}
	}
	if token == ""{
		return locationPosts
	}
	unavailableUsers := getUnavailableUsers(token)
	for i, post := range locationPosts {
		for _, username := range unavailableUsers.Usernames {
			if username == post.Owner{
				locationPosts = append(locationPosts[:i], locationPosts[i+1:]...)
				break
			}
		}
	}
	return locationPosts
	// TODO limit? pagable?
}

func (service *PostService) SearchTag(tag string, username string, token string) interface{} {
	publicPostsDocuments := service.PostRepository.GetPublicPosts()
	publicPosts := CreatePostsFromDocuments(publicPostsDocuments)
	var tagPosts []model.Post

	for _, p := range publicPosts {
		for _, t := range p.Tags {
			if strings.Contains(strings.ToLower(t), strings.ToLower(tag)) {
				tagPosts = append(tagPosts, p)
			}
		}
	}

	for i, s := range tagPosts {
		for j, _ := range s.Path {
			b, err := ioutil.ReadFile(s.Path[j])
			if err != nil {
				fmt.Print(err)
			}
			var image model.PostImages
			image.Image = b
			tagPosts[i].Images = append(tagPosts[i].Images, image)
		}
	}
	if token == ""{
		return tagPosts
	}
	unavailableUsers := getUnavailableUsers(token)
	for i, post := range tagPosts {
		for _, username := range unavailableUsers.Usernames {
			if username == post.Owner{
				tagPosts = append(tagPosts[:i], tagPosts[i+1:]...)
				break
			}
		}
	}

	return tagPosts
	// TODO limit? pagable?
}

func (service *PostService) UpdatePostsPrivacy(username string, privacy bool) {
	service.PostRepository.UpdatePostsPrivacyByOwner(username, privacy)
}

func (service *PostService) GetReactedPosts(username string) interface{} {
	var userReactions model.UserReaction
	userReactionsDocument, err := service.PostRepository.GetUserReactions(username)
	if err != nil {
		return userReactions
	}
	bsonBytes, _ := bson.Marshal(userReactionsDocument)
	_ = bson.Unmarshal(bsonBytes, &userReactions)

	var reactedPostsDTO dto.UserPostReactionDTO

	for _, s := range userReactions.LikedPosts {
		reactedPostsDTO.LikedPosts = append(reactedPostsDTO.LikedPosts, service.GetPostByID(s.String()))
	}
	for _, s := range userReactions.DislikedPosts {
		reactedPostsDTO.DislikedPosts = append(reactedPostsDTO.DislikedPosts, service.GetPostByID(s.String()))
	}
	return reactedPostsDTO
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
	post.IsPrivate = postDTO.IsPrivate
	post.Location = postDTO.Location
	post.NumberOfDislikes, post.NumberOfLikes, post.NumberOfReaches = 0, 0, 0
	post.IsAdd = postDTO.IsAdd
	if len(paths) > 1 {
		post.IsAlbum = true
	} else {
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

func sendDeleteRequests(username string, token string) {
	client := &http.Client{}
	// AUTH
	requestUrl := fmt.Sprintf("http://%s:%s/deleteUser/"+username, os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest(http.MethodDelete, requestUrl, nil)
	req.Header.Set("Host", "http://post-service:8080")
	req.Header.Set("Authorization", token)
	client.Do(req)

	//USER
	requestUrl = fmt.Sprintf("http://%s:%s/deleteUser/"+username, os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	req, _ = http.NewRequest(http.MethodDelete, requestUrl, nil)
	req.Header.Set("Host", "http://post-service:8080")
	req.Header.Set("Authorization", token)
	client.Do(req)
}

func stringInSlice(s string, list []string) bool {
	for _, b := range list {
		if b == s {
			return true
		}
	}
	return false
}

func removeStringFromSLice(s string, l []string) []string {
	for i, v := range l {
		if v == s {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func removeUUIDFromSLice(s uuid.UUID, l []uuid.UUID) []uuid.UUID {
	for i, v := range l {
		if v == s {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func getRelationType(username string, token string) model.RelationType {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/getRelationship/"+username, os.Getenv("FOLLOWERS_SERVICE_DOMAIN"), os.Getenv("FOLLOWERS_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://user-service:8080")
	if token == "" {
		return model.RelationType{Relation: model.NotFollowing}
	}
	req.Header.Set("Authorization", token)
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
	}
	var relationType model.RelationType
	_ = json.NewDecoder(res.Body).Decode(&relationType)
	return relationType
}

func getUnavailableUsers(token string) model.UsersList {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/getUnavailableUsers", os.Getenv("FOLLOWERS_SERVICE_DOMAIN"), os.Getenv("FOLLOWERS_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://post-service:8080")
	if  token == ""{
		return model.UsersList{Usernames: nil}
	}
	req.Header.Set("Authorization", token)
	res, err2 := client.Do(req)
	if err2 != nil {
		return model.UsersList{Usernames: nil}
	}
	var users model.UsersList
	_ = json.NewDecoder(res.Body).Decode(&users)
	return users
}

func getFollowingUsers(token string) model.UsersList {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/following", os.Getenv("FOLLOWERS_SERVICE_DOMAIN"), os.Getenv("FOLLOWERS_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://post-service:8080")
	if  token == ""{
		return model.UsersList{Usernames: nil}
	}
	req.Header.Set("Authorization", token)
	res, err2 := client.Do(req)
	if err2 != nil {
		return model.UsersList{Usernames: nil}
	}
	var users model.UsersList
	_ = json.NewDecoder(res.Body).Decode(&users)
	return users
}
