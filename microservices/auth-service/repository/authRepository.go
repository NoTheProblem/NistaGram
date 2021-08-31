package repository

import (
	"auth-service/dto"
	"auth-service/model"
	"fmt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	Database *gorm.DB
}

func (repository *AuthRepository) RegisterUser(user *model.User) error {
	result := repository.Database.Create(user)
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not registered")
	}
	fmt.Println("User successfully registered! [auth-repository]")
	return nil
}


func (repository *AuthRepository) FindUserByUsername(username string) (*model.User, error){
	user := &model.User{}
	err := repository.Database.Table("users").First(&user, "username = ?", username).Error
	if  err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *AuthRepository) Delete(username string) {
	repository.Database.Table("users").Where("username = ?", username).Delete(&model.User{})
}

func (repository *AuthRepository) RegisterBusiness(business *model.BusinessRequests) error {
	result := repository.Database.Create(business)
	if result.RowsAffected == 0 {
		return fmt.Errorf("business not registered")
	}
	fmt.Println("Business successfully registered! [auth-repository]")
	return nil
}

func (repository *AuthRepository) GetAllPendingBusinessRequests() ([]model.BusinessRequests, error)  {
	var requests []model.BusinessRequests
	err := repository.Database.Table("business_requests").Find(&requests, "status", model.Pending).Error
	if  err != nil {
		return nil, err
	}
	return requests, nil
}

func (repository *AuthRepository) FindBusinessRequestByUsername(username string) (*model.BusinessRequests, error) {
	businessRequest := &model.BusinessRequests{}
	err := repository.Database.Table("business_requests").First(&businessRequest, "username = ?", username).Error
	if  err != nil {
		return nil, err
	}
	return businessRequest, nil
}

func (repository *AuthRepository) UpdateBusinessRequestStatus(d *dto.BusinessRequestAnswer) error {
	err := repository.Database.Table("business_requests").Where("username = ?", d.Username).Update("status",d.Status)
	if err != nil{
		return err.Error
	}
	return nil
}

func (repository *AuthRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := repository.Database.Table("users").Find(&users).Error
	if  err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *AuthRepository) ChangeRole(change dto.AuthDTO) error {
	err := repository.Database.Table("users").Where("username = ?", change.Username).Update("user_role",change.Role)
	if err != nil{
		return err.Error
	}
	return nil
}


