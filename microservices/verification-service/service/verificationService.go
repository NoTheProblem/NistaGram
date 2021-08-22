package service

import (
	"github.com/google/uuid"
	"time"
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
