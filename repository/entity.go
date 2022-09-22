package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Entity any

type BaseModel struct {
	ID        string         `gorm:"type:char(36);primary_key;" json:"ID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (base *BaseModel) BeforeCreate(_ *gorm.DB) (err error) {
	base.ID = uuid.NewString()
	return
}

type News struct {
	BaseModel
	Name         string
	ThumbnailURL string
	Content      string
	Tags         string
}
