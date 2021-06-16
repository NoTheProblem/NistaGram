package repository

import (
	"gorm.io/gorm"
	"post-service/model"
)

type LocationRepository struct {
	Database *gorm.DB
}

func (repo *LocationRepository) CreateLocation(location *model.Location) error {
	result := repo.Database.Create(location)
	print(result.Error)
	print(result.RowsAffected)
	return nil
}


func (repo *LocationRepository) ExistsByName(locationName string) bool {

	if err := repo.Database.First(&model.Tag{}, "name = ?", locationName).Error; err != nil {
		return false
	}
	return true
}

