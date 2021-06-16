package repository

import (
	"gorm.io/gorm"
	"post-service/model"
)

type TagRepository struct {
	Database *gorm.DB
}

func (repo *TagRepository) CreateTag(tag *model.Tag) error {
	result := repo.Database.Create(tag)
	print(result.Error)
	print(result.RowsAffected)
	return nil
}


func (repo *TagRepository) ExistsByName(tagName string) bool {

	if err := repo.Database.First(&model.Tag{}, "name = ?", tagName).Error; err != nil {
		return false
	}
	return true
}