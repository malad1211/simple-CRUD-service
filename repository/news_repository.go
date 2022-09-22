package repository

import (
	"gorm.io/gorm"
	"inspiredlab/domain"
)

type NewsRepository struct {
	Repository[domain.News, News]
}

func NewNewsRepository(db *gorm.DB) *NewsRepository {
	return &NewsRepository{
		Repository: New[domain.News, News](db),
	}
}
