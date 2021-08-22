package repository

import (
	"fmt"
	"gorm.io/gorm"
	"user-service/model"
)

type UserRepository struct {
	Database *gorm.DB
}

func (repository *UserRepository) RegisterUser(user *model.User) error {
	result := repository.Database.Create(user)
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not registered")
	}
	fmt.Println("User successfully registered! [user-service]")
	return nil
}

func (repository *UserRepository) UpdateUserProfileInfo(user *model.User) error {
	result := repository.Database.Updates(user)
	if result.RowsAffected == 0 {
		return fmt.Errorf("user did not update")
	}
	fmt.Println("User successfully updated!")
	return nil
}

func (repository *UserRepository) FindUserByUsername(username string) (*model.User, error){
	user := &model.User{}
	err := repository.Database.Table("users").First(&user, "username = ?", username).Error
	if  err != nil {
		return nil, err
	}
	return user, nil
}
