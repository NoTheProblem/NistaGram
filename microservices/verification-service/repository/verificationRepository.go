package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"verification-service/model"
)

type VerificationRepository struct {
	Database *gorm.DB
}

func (repository *VerificationRepository) AddRequest(request *model.VerificationRequest) error {
	result := repository.Database.Create(request)
	if result.RowsAffected == 0 {
		return fmt.Errorf("request not added")
	}
	fmt.Println("Request successfully added")
	return nil
}


func (repository *VerificationRepository) UpdateUserProfileInfo(user *model.VerificationRequest) error {
	result := repository.Database.Updates(user)
	if result.RowsAffected == 0 {
		return fmt.Errorf("request not update")
	}
	fmt.Println("\"Request successfully updated!")
	return nil
}

func (repository *VerificationRepository) GetAllUnAnsweredRequests() ([]model.VerificationRequest, error) {
	var requests []model.VerificationRequest
	err := repository.Database.Table("verification_requests").Find(&requests, "is_answered is false").Error
	if  err != nil {
		return nil, err
	}
	return requests, nil

}

func (repository *VerificationRepository) GetVerificationRequestById(uid uuid.UUID) (*model.VerificationRequest, error) {
	request := &model.VerificationRequest{}
	err := repository.Database.Table("verification_requests").First(&request, "id = ?", uid).Error
	if  err != nil {
		return nil, err
	}
	return request, nil

}

func (repository *VerificationRepository) UpdateVerificationRequest(request *model.VerificationRequest) error{
	result := repository.Database.Updates(request)
	if result.RowsAffected == 0 {
		return fmt.Errorf("user did not update")
	}
	return nil
}

