package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"post-service/dto"
	"post-service/model"
	"post-service/service"
	"strconv"
)

type PostHandler struct {
	PostService *service.PostService
}

func (handler *PostHandler) CreateNewPost(w http.ResponseWriter, r *http.Request) {

	user , err := getUserFromToken(r)
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	makeDirectoryIfNotExists(user.Username)

	r.ParseMultipartForm(10 << 20)
	var location = r.FormValue("location")
	var tags = r.FormValue("tags")
	var description = r.FormValue("description")
	var numberOfImagesStr = r.FormValue("numberOfImages")
	var isPublic, _ = strconv.ParseBool(r.FormValue("isPublic"))
	var numberOfImages, _ = strconv.Atoi(numberOfImagesStr)
	fileNames := make([]string, numberOfImages)
	fmt.Println(numberOfImages)
	for i := 0; i < numberOfImages; i++ {
		var file, fileHandler, err = r.FormFile("myFile" + strconv.Itoa(i))
		fmt.Println("myFile" + strconv.Itoa(i))
		fmt.Println(i)
		fmt.Println(strconv.Itoa(i))
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			fmt.Fprintf(w, err.Error())
			return
		}
		defer file.Close()
		fmt.Printf("File: %+v\n", i)
		fmt.Printf("Uploaded File: %+v\n", fileHandler.Filename)
		fmt.Printf("File Size: %+v\n", fileHandler.Size)
		fmt.Printf("MIME Header: %+v\n", fileHandler.Header)
		pictureType := filepath.Ext(fileHandler.Filename)
		fileName := fmt.Sprintf("*"+pictureType)
		tempFile, err := ioutil.TempFile(user.Username, fileName)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, err.Error())
		}
		fmt.Println(tempFile.Name())
		fmt.Println(tempFile)
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, err.Error())
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		fileNames[i]  = tempFile.Name()
	}

	fmt.Println(fileNames)
	// return that we have successfully uploaded our file!
	var post dto.PostDTO
	post.Location = location
	post.IsPublic = isPublic
	json.Unmarshal([]byte(tags), &post.Tags)

	post.Description = description
	handler.PostService.AddPost(post,user.Username,fileNames)
	w.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	publicPosts := handler.PostService.GetAll()
	publicPostsJson, err := json.Marshal(publicPosts)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(publicPostsJson)
	}
}

func (handler *PostHandler) GetHomeFeed(writer http.ResponseWriter, request *http.Request) {
	//username := getUsernameFromToken(request)
	publicPosts :=handler.PostService.GetHomeFeed("username")
	writer.Header().Set("Content-Type", "application/json")
	publicPostsJson, err := json.Marshal(publicPosts)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(publicPostsJson)
	}
}

func (handler *PostHandler) GetPostsByUsername(writer http.ResponseWriter, request *http.Request) {
	// TODO private?
	vars := mux.Vars(request)
	username := vars["username"]
	publicPosts :=handler.PostService.GetProfilePosts(username)
	writer.Header().Set("Content-Type", "application/json")
	publicPostsJson, err := json.Marshal(publicPosts)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(publicPostsJson)
	}


}

func (handler *PostHandler) GetPost(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	publicPosts :=handler.PostService.GetPostByID(id)
	writer.Header().Set("Content-Type", "application/json")
	publicPostsJson, err := json.Marshal(publicPosts)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(publicPostsJson)
	}
}

func (handler *PostHandler) CommentPost(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println("CommentHandler")
	var commentDTO dto.CommentDTO
	err = json.NewDecoder(request.Body).Decode(&commentDTO)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.CommentPost(commentDTO, user.Username)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) LikePost(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	var postId dto.IdDTO
	err = json.NewDecoder(request.Body).Decode(&postId)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.LikePost(postId.Id, user.Username)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) DislikePost(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	var postId dto.IdDTO
	err = json.NewDecoder(request.Body).Decode(&postId)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.DisLikePost(postId.Id, user.Username)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) ReportPost(writer http.ResponseWriter, request *http.Request) {
	_ , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	var postId dto.IdDTO
	err = json.NewDecoder(request.Body).Decode(&postId)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.ReportPost(postId.Id)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) GetAllUnansweredReports(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if model.Role(user.Role) != model.Administrator {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	reports :=handler.PostService.GetAllUnansweredReports()
	writer.Header().Set("Content-Type", "application/json")
	reportsJson, err := json.Marshal(reports)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(reportsJson)
	}

}

func (handler *PostHandler) AnswerReport(writer http.ResponseWriter, request *http.Request) {
	user , err := getUserFromToken(request)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	if model.Role(user.Role) != model.Administrator {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	var reportDTO dto.ReportDTO
	err = json.NewDecoder(request.Body).Decode(&reportDTO)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PostService.AnswerReport(reportDTO, request.Header.Get("Authorization"))
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) SearchTag(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	tag := vars["tag"]
	fmt.Println("tag")
	fmt.Println(tag)
	publicPosts :=handler.PostService.SearchTag(tag)
	writer.Header().Set("Content-Type", "application/json")
	publicPostsJson, err := json.Marshal(publicPosts)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(publicPostsJson)
	}
}

func (handler *PostHandler) SearchLocation(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	location := vars["location"]
	fmt.Println("location")
	fmt.Println(location)
	publicPosts :=handler.PostService.SearchLocation(location)
	writer.Header().Set("Content-Type", "application/json")
	publicPostsJson, err := json.Marshal(publicPosts)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(publicPostsJson)
	}
}

func getUserFromToken(r *http.Request) (model.Auth, error) {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/authorize", os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://post-service:8080")
	fmt.Println(r.Header.Get("Authorization"))
	if  r.Header.Get("Authorization") == ""{
		return model.Auth{}, errors.New("no logged user")
	}
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
	}

	var user model.Auth
	err := json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return model.Auth{}, err
	}

	if user.Username == ""{
		return model.Auth{}, errors.New("no such user")
	}
	return user, nil
}



func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}
