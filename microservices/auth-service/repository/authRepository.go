package repository

import (
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
		return fmt.Errorf("User not registered!")
	}
	fmt.Println("User successfuly registered!")
	return nil
}

func (repository *AuthRepository) UpdateUser(user *model.User) error {
	result := repository.Database.Updates(user)
	if result.RowsAffected == 0 {
		return fmt.Errorf("User not update!")
	}
	fmt.Println("User successfuly updated!")
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


