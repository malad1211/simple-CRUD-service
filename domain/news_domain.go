package domain

type News struct {
	BaseModel
	Name         string
	ThumbnailURL string
	Content      string
	Tags         string
}

type NewsRepository interface {
	Repository[News]
}

type NewsService interface {
	Create(news News) (*News, error)
	Update(ID string, news News) error
	Delete(ID string) error
	Get(ID string) (*News, error)
}
