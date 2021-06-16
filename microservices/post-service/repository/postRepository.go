package repository

import (
	"gorm.io/gorm"
	"post-service/model"
)

type PostRepository struct {
	Database *gorm.DB
}

func (repo *PostRepository) CreatePost(post *model.Post) error {
	result := repo.Database.Create(post)
	print(result.Error)
	print(result.RowsAffected)
	return nil
}
