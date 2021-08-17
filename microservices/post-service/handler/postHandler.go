package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"post-service/service"
)

type PostHandler struct {
	PostService *service.PostService
}

func (handler *PostHandler) CreateNewPost(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/auth/authorize", nil)
	fmt.Println( r.Header.Get("Authorization"))
	fmt.Println( r.Header.Get("Host"))
	req.Header.Set("Host", "http://localhost:4200")
	req.Header.Set("Authorization", r.Header.Get("Authorization"))
	fmt.Println(req)
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("didnt even call auth")
		fmt.Println(err2)
		return
	}
	body, err5 := ioutil.ReadAll(res.Body)
	if err5 != nil {
		log.Fatalln(err5)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Println(sb)
	param := mux.Vars(r)
	username := param["username"]
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
	fileName := fmt.Sprintf("*.jpg")
	makeDirectoryIfNotExists(username)
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
	w.WriteHeader(http.StatusOK)


}

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}



