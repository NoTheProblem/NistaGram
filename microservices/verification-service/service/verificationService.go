package service

import (
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"verification-service/dto"
	"verification-service/model"
	"verification-service/repository"
)

type VerificationService struct {
	VerificationRepository *repository.VerificationRepository
}

func (handler *VerificationService) AddVerificationRequest(request *model.VerificationRequest) error {

	request.DateSubmitted = time.Now()
	request.IsAnswered = false
	request.Id = uuid.New()
	err := handler.VerificationRepository.AddRequest(request)
	if err != nil {
		return err
	}
	return nil

}

func (handler *VerificationService) GetAllUnAnsweredRequests() ( []model.VerificationRequest ,error) {

	requests, err := handler.VerificationRepository.GetAllUnAnsweredRequests()
	if err != nil{
		return nil, err
	}
	for i, request := range requests {
		b, err := ioutil.ReadFile(request.Path)
		if err != nil {
			fmt.Print(err)
		}
		requests[i].Image = b
	}

	return requests, nil
}

func (handler *VerificationService) AnswerRequest(requestDTO *dto.VerificationAnswerDTO, token string) error {
	uid, err := uuid.Parse(requestDTO.Id)
	if err != nil {
		return nil
	}
	request, err := handler.VerificationRepository.GetVerificationRequestById(uid)
	if err != nil{
		return err
	}
	// TODO send notification?
	request.DateAnswered = time.Now()
	request.IsAnswered = true
	request.VerificationAnswer = requestDTO.VerificationAnswer
	if requestDTO.VerificationAnswer {

		client := &http.Client{}
		// TODO put env
		requestUrl := fmt.Sprintf("http://%s:%s/verify/", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))

		req, err := http.NewRequest(http.MethodPut, requestUrl + request.Username, nil)
		if err != nil {
			panic(err)
		}
		req.Header.Set("Host", "http://verification-service:8080")
		req.Header.Set("Authorization", token)
		_, err = client.Do(req)
		if err != nil {
			return err
		}

	}else{
		request.Answer = requestDTO.Answer
	}
	err = handler.VerificationRepository.UpdateVerificationRequest(request)
	if err != nil{
		return err
	}
	return nil
}
