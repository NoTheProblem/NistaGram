package repository

import (
	"fmt"
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

func (repository *VerificationRepository) FindUserByUsername(username string) (*model.VerificationRequest, error){
	request := &model.VerificationRequest{}
	err := repository.Database.Table("verification_requests").First(&request, "username = ?", username).Error
	if  err != nil {
		return nil, err
	}
	return request, nil
}
