package domain

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	StatusId  int            `gorm:"primaryKey" json:"status_id"`
	StatusTxt string         `json:"status_txt" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type StatusRepository interface {
	Create(status Status) (Status, error)
	Update(status Status) (bool, error)
	FindByID(ID int) (Status, error)
	Delete(status Status) (bool, error)
}

type StatusService interface {
	Create(status Status) (Status, error)
	Update(ID int, status Status) (bool, error)
	Delete(ID int) (bool, error)
}
