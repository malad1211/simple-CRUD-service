package domain

import (
	"time"
)

type Model interface {
}

type BaseModel struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository[M Model] interface {
	Create(model *M) (*M, error)
	Save(model *M) (*M, error)
	FindOne(model *M) (*M, error)
	Delete(model *M) error
	Update(where *M, update *M) (*M, error)
}

type Sort map[string]int

type Pageable struct {
	Size int64
	Page int64
	Sort Sort
}

func NewPageable(page int64, size int64) Pageable {
	// index db is zero
	if page > 0 {
		page--
	}
	return Pageable{Size: size, Page: page}
}
func (p Pageable) Skip() *int64 {
	size := p.Page * p.Size
	return &size
}

type Page struct {
	Pageable
	HasNext      bool
	HasPrevious  bool
	TotalElement int
	TotalPage    int
	List         interface{}
}

func NewPage(pageable Pageable, hasNext bool, hasPrevious bool, totalElement int, totalPage int, list interface{}) *Page {
	return &Page{Pageable: pageable, HasNext: hasNext, HasPrevious: hasPrevious, TotalElement: totalElement, TotalPage: totalPage, List: list}
}
