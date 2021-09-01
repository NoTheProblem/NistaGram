package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"verification-service/dto"
	"verification-service/model"
	"verification-service/service"
)

type VerificationHandler struct {
	VerificationService *service.VerificationService
}

func (handler *VerificationHandler) CreateNewUserRequest(writer http.ResponseWriter, request *http.Request) {
	user , _ := getUserFromToken(request)
	fmt.Println(user.Username)
	makeDirectoryIfNotExists(user.Username)

	request.ParseMultipartForm(10 << 20)

	var file, fileHandler, err = request.FormFile("myFile")
	var firstName = request.FormValue("firstName")
	var lastName = request.FormValue("lastName")
	var category = request.FormValue("category")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", fileHandler.Filename)
	fmt.Printf("File Size: %+v\n", fileHandler.Size)
	fmt.Printf("MIME Header: %+v\n", fileHandler.Header)
	pictureType := filepath.Ext(fileHandler.Filename)
	fileName := fmt.Sprintf("*"+pictureType)
	tempFile, err := ioutil.TempFile(user.Username, fileName)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
	}
	fmt.Println(tempFile.Name())
	fmt.Println(tempFile)
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(writer, err.Error())
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	var verRequest model.VerificationRequest
	verRequest.Username = user.Username
	verRequest.FirstName = firstName
	verRequest.LastName = lastName
	verRequest.Category = category
	verRequest.Path = tempFile.Name()
	errS := handler.VerificationService.AddVerificationRequest(&verRequest)
	if errS != nil{
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)

}

func (handler *VerificationHandler) AnswerRequest(writer http.ResponseWriter, request *http.Request) {
	user , _ := getUserFromToken(request)
	if model.Role(user.Role) != model.Administrator {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	var verificationAnswer dto.VerificationAnswerDTO
	err := json.NewDecoder(request.Body).Decode(&verificationAnswer)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.VerificationService.AnswerRequest(&verificationAnswer,request.Header.Get("Authorization"))
	if err != nil{
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
	
}

func (handler *VerificationHandler) GetAllUnAnswered(writer http.ResponseWriter, request *http.Request) {
	user , _ := getUserFromToken(request)
	if model.Role(user.Role) != model.Administrator {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	unAnsweredRequests, _ :=handler.VerificationService.GetAllUnAnsweredRequests()

	writer.Header().Set("Content-Type", "application/json")
	publicPostsJson, err := json.Marshal(unAnsweredRequests)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(publicPostsJson)
	}

}

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

func getUserFromToken(r *http.Request) (model.Auth, error) {
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/authorize", os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://verification-service:8080")
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
