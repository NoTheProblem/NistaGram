package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"post-service/service"
)

type PostHandler struct {
	PostService *service.PostService
}

func (handler *PostHandler) CreateNewPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	r.ParseMultipartForm(10 << 20)

	var file, fileHandler, err = r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", fileHandler.Filename)
	fmt.Printf("File Size: %+v\n", fileHandler.Size)
	fmt.Printf("MIME Header: %+v\n", fileHandler.Header)
	fileName := fmt.Sprintf("%s_-*.jpg", username)
	tempFile, err := ioutil.TempFile("repository", fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tempFile.Name())
	fmt.Println(tempFile)
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")

}



