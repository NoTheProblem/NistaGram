package repository

import (
	"gorm.io/gorm"
	"post-service/model"
)

type CommentRepository struct {
	Database *gorm.DB
}

func (repo *CommentRepository) CreateComment(comment *model.Comment) error {
	result := repo.Database.Create(comment)
	print(result.Error)
	print(result.RowsAffected)
	return nil
}


