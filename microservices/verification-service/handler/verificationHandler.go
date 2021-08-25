package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/authorize", os.Getenv("AUTH_SERVICE_DOMAIN"), os.Getenv("AUTH_SERVICE_PORT"))
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req.Header.Set("Host", "http://verification-service:8080")
	req.Header.Set("Authorization", request.Header.Get("Authorization"))
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	body, err5 := ioutil.ReadAll(res.Body)
	if err5 != nil {
		log.Fatalln(err5)
	}
	//Convert the body to type string
	sb := string(body)
	username := sb[1:len(sb)-1]
	fmt.Println(username)
	makeDirectoryIfNotExists(username)

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
	tempFile, err := ioutil.TempFile(username, fileName)
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
	verRequest.Username = username
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
	// TODO check token user role

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
	// TODO check token user role

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
