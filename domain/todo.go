package domain

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID             int            `json:"id"`
	Title          string         `json:"title" validate:"required"`
	Description    string         `json:"description" validate:"required"`
	DueDate        string         `json:"due_date" validate:"required"`
	PersonInCharge string         `json:"person_in_charge" validate:"required"`
	Status         string         `json:"status" validate:"required"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

type TodoRepository interface {
	Create(todo *Todo) (*Todo, error)
	Update(todo *Todo) (bool, error)
	FindByID(ID int) (*Todo, error)
	FindAll() ([]*Todo, error)
	Delete(todo *Todo) (bool, error)
}

type TodoService interface {
	Create(todo *Todo) (*Todo, error)
	Update(ID int, todo *Todo) (bool, error)
	FindByID(ID int) (*Todo, error)
	FindAll() ([]*Todo, error)
	Delete(ID int) (bool, error)
}
