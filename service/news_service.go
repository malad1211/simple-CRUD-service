package service

import (
	"inspiredlab/domain"
)

type NewsService struct {
	repo domain.NewsRepository
}

func NewNewsService(repo domain.NewsRepository) *NewsService {
	return &NewsService{repo: repo}
}

func (u *NewsService) Create(news domain.News) (*domain.News, error) {
	created, err := u.repo.Create(&news)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (u *NewsService) Update(ID string, news domain.News) error {
	_, err := u.repo.Update(
		&domain.News{
			BaseModel: domain.BaseModel{
				ID: ID,
			},
		},
		&news,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *NewsService) Delete(ID string) error {
	err := u.repo.Delete(&domain.News{
		BaseModel: domain.BaseModel{
			ID: ID,
		}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *NewsService) Get(ID string) (*domain.News, error) {
	news, err := u.repo.FindOne(
		&domain.News{
			BaseModel: domain.BaseModel{
				ID: ID,
			},
		},
	)
	return news, err
}
