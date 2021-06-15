package repository

import "gorm.io/gorm"

type PostRepository struct {
	Database *gorm.DB
}

