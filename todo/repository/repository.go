package repository

import (
	"errors"

	"github.com/arnoldtherigan15/final_project_golang/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	// var
	db.AutoMigrate(&domain.Todo{})
	if err := db.AutoMigrate(&domain.Status{}); err == nil && db.Migrator().HasTable(&domain.Status{}) {
		if err := db.First(&domain.Status{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			var status = []domain.Status{
				{StatusTxt: "New"},
				{StatusTxt: "OnGoing"},
				{StatusTxt: "Done"},
				{StatusTxt: "Deleted"},
			}
			db.Create(&status)
		}
	}
	return &repository{db}
}

func (r *repository) Create(todo *domain.Todo) (*domain.Todo, error) {
	err := r.db.Create(todo).Error
	if err != nil {
		return &domain.Todo{}, err
	}
	return todo, nil
}

func (r *repository) FindAll() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (r *repository) Update(todo *domain.Todo) (bool, error) {
	err := r.db.Save(todo).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FindByID(ID int) (*domain.Todo, error) {
	var todo domain.Todo
	err := r.db.Where("id = ?", ID).Find(&todo).Error
	if err != nil {
		return &todo, err
	}
	return &todo, nil
}

func (r *repository) Delete(todo *domain.Todo) (bool, error) {
	if err := r.db.Delete(todo).Error; err != nil {
		return false, err
	}
	return true, nil
}
