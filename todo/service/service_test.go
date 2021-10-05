package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/arnoldtherigan15/final_project_golang/domain"
	"github.com/arnoldtherigan15/final_project_golang/domain/mocks"
	_todoService "github.com/arnoldtherigan15/final_project_golang/todo/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTodoService_Create(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := &domain.Todo{
		Title:          "Belajar Golang",
		Description:    "Pointer, middleware dan api",
		DueDate:        "12-12-2021",
		PersonInCharge: "arnold",
		Status:         "New",
	}
	mockEmptyTodo := &domain.Todo{}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.Mock.On("Create", mock.Anything).Return(mockTodo, nil).Once()
		service := _todoService.NewService(mockTodoRepo)
		todo, err := service.Create(mockTodo)
		assert.NoError(t, err)
		assert.NotNil(t, todo)

		mockTodoRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockTodoRepo.On("Create", mock.Anything).Return(mockEmptyTodo, errors.New("Unexpected")).Once()

		service := _todoService.NewService(mockTodoRepo)

		todo, err := service.Create(mockEmptyTodo)

		assert.Error(t, err)
		assert.Equal(t, mockEmptyTodo, todo)

		mockTodoRepo.AssertExpectations(t)
	})
}

func TestTodoService_FindAll(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockArrTodo := []*domain.Todo{
		&domain.Todo{
			Title:          "Belajar Golang",
			Description:    "Pointer, middleware dan api",
			DueDate:        "12-12-2021",
			PersonInCharge: "arnold",
			Status:         "New",
		},
	}
	mockEmptyTodo := []*domain.Todo{
		&domain.Todo{},
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.Mock.On("FindAll", mock.Anything).Return(mockArrTodo, nil).Once()
		service := _todoService.NewService(mockTodoRepo)
		todos, err := service.FindAll()
		assert.NoError(t, err)
		assert.NotNil(t, todos)
		mockTodoRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockTodoRepo.On("FindAll", mock.Anything).Return(mockEmptyTodo, errors.New("Unexpected")).Once()

		service := _todoService.NewService(mockTodoRepo)

		todos, err := service.FindAll()

		assert.Error(t, err)
		assert.Equal(t, mockEmptyTodo, todos)

		mockTodoRepo.AssertExpectations(t)
	})
}

func TestTodoservice_Update(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := &domain.Todo{
		ID:             1,
		Title:          "Belajar Golang",
		Description:    "Pointer, middleware dan api",
		DueDate:        "12-12-2021",
		PersonInCharge: "arnold",
		Status:         "New",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}
	mockEmptyTodo := &domain.Todo{}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("FindByID", mock.Anything).Return(mockTodo, nil).Once()
		mockTodoRepo.On("Update", mock.Anything, mock.Anything).Return(true, nil).Once()
		service := _todoService.NewService(mockTodoRepo)
		isUpdated, err := service.Update(1, mockTodo)

		assert.NoError(t, err)
		assert.Equal(t, isUpdated, true)
		assert.NotNil(t, isUpdated)

		mockTodoRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockTodoRepo.On("FindByID", mock.Anything).Return(mockTodo, nil).Once()
		mockTodoRepo.On("Update", mock.Anything).Return(false, errors.New("unexpected")).Once()

		service := _todoService.NewService(mockTodoRepo)
		isUpdated, err := service.Update(1, mockEmptyTodo)
		assert.Error(t, err)
		assert.Equal(t, isUpdated, false)

		mockTodoRepo.AssertExpectations(t)
	})

}

func TestTodoservice_FindByID(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := &domain.Todo{
		ID:             1,
		Title:          "Belajar Golang",
		Description:    "Pointer, middleware dan api",
		DueDate:        "12-12-2021",
		PersonInCharge: "arnold",
		Status:         "New",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}
	mockEmptyTodo := &domain.Todo{}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("FindByID", mock.Anything).Return(mockTodo, nil).Once()
		service := _todoService.NewService(mockTodoRepo)
		todo, err := service.FindByID(1)

		assert.NoError(t, err)
		assert.Equal(t, mockTodo, todo)
		assert.NotNil(t, todo)

		mockTodoRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockTodoRepo.On("FindByID", mock.Anything).Return(mockEmptyTodo, errors.New("unexpected")).Once()

		service := _todoService.NewService(mockTodoRepo)
		todo, err := service.FindByID(1)
		assert.Error(t, err)
		assert.Equal(t, todo, mockEmptyTodo)

		mockTodoRepo.AssertExpectations(t)
	})

}

func TestTodoservice_Delete(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodo := &domain.Todo{
		ID:             1,
		Title:          "Belajar Golang",
		Description:    "Pointer, middleware dan api",
		DueDate:        "12-12-2021",
		PersonInCharge: "arnold",
		Status:         "New",
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockTodoRepo.On("FindByID", mock.Anything).Return(mockTodo, nil).Once()
		mockTodoRepo.On("Delete", mock.Anything).Return(true, nil).Once()
		service := _todoService.NewService(mockTodoRepo)
		isDeleted, err := service.Delete(1)

		assert.NoError(t, err)
		assert.Equal(t, isDeleted, true)
		assert.NotNil(t, isDeleted)

		mockTodoRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockTodoRepo.On("FindByID", mock.Anything).Return(mockTodo, nil).Once()
		mockTodoRepo.On("Delete", mock.Anything).Return(false, errors.New("unexpected")).Once()
		service := _todoService.NewService(mockTodoRepo)
		isDeleted, err := service.Delete(1)
		assert.Error(t, err)
		assert.Equal(t, isDeleted, false)

		mockTodoRepo.AssertExpectations(t)
	})

}
