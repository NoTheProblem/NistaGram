package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"post-service/dto"
	"post-service/service"
)

type PostHandler struct {
	PostService *service.PostService
}

func (handler *PostHandler) CreateNewPost(w http.ResponseWriter, r *http.Request) {

	username := getUsernameFromToken(r)
	makeDirectoryIfNotExists(username)

	r.ParseMultipartForm(10 << 20)

	var file, fileHandler, err = r.FormFile("myFile")
	var description = r.FormValue("description")
	var location = r.FormValue("location")
	var tags = r.FormValue("tags")
	fmt.Println(description)
	fmt.Println(location)
	fmt.Println(tags)
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", fileHandler.Filename)
	fmt.Printf("File Size: %+v\n", fileHandler.Size)
	fmt.Printf("MIME Header: %+v\n", fileHandler.Header)
	pictureType := filepath.Ext(fileHandler.Filename)
	fileName := fmt.Sprintf("*"+pictureType)
	tempFile, err := ioutil.TempFile(username, fileName)
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
	// return that we have successfully uploaded our file!
	var post dto.PostDTO
	post.Location = location
	json.Unmarshal([]byte(tags), &post.Tags)
	//post.Tags = tags
	post.Description = description
	handler.PostService.AddPost(post,username,tempFile.Name())
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



func (handler *PostHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	handler.PostService.PostRepository.Delete()
	writer.WriteHeader(http.StatusOK)

}

func (handler *PostHandler) GetPostsByUsername(writer http.ResponseWriter, request *http.Request) {
	//username := getUsernameFromToken(request)
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

func getUsernameFromToken(r *http.Request) string {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/authorize", os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://user-service:8080")
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(res)
	body, err5 := ioutil.ReadAll(res.Body)
	if err5 != nil {
		log.Fatalln(err5)
	}
	//Convert the body to type string
	sb := string(body)
	username := sb[1:len(sb)-1]
	return username
}



func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}
