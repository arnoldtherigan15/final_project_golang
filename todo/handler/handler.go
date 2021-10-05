package handler

import (
	"net/http"
	"strconv"

	"github.com/arnoldtherigan15/final_project_golang/domain"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service domain.TodoService
}

func NewHandler(group *echo.Group, service domain.TodoService) {
	handler := &Handler{service}
	g := group.Group("/todos")
	g.POST("", handler.Create)
	g.GET("", handler.FindAll)
	g.GET("/:id", handler.FindByID)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}

func isRequestValid(todo *domain.Todo) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(todo); err != nil {
		return false, err
	}
	return true, nil
}

func (h *Handler) Create(c echo.Context) (err error) {
	var todo domain.Todo
	if err = c.Bind(&todo); err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	if ok, err := isRequestValid(&todo); !ok {
		errMsg := domain.ErrorValidationFormatter(err)
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, errMsg)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	createdTodo, err := h.service.Create(&todo)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, domain.ErrInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusCreated, createdTodo)
}

func (h *Handler) FindAll(c echo.Context) (err error) {

	todos, err := h.service.FindAll()

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, domain.ErrInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusOK, todos)
}

func (h *Handler) Delete(c echo.Context) (err error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	isDeleted, err := h.service.Delete(ID)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusNotFound, errResponse)
	}

	response := map[string]bool{
		"is_delete": isDeleted,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Update(c echo.Context) (err error) {
	var todo domain.Todo
	err = c.Bind(&todo)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	if ok, err := isRequestValid(&todo); !ok {
		errMsg := domain.ErrorValidationFormatter(err)
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, errMsg)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	updatedTodo, err := h.service.Update(ID, &todo)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusNotFound, errResponse)
	}

	response := map[string]bool{
		"is_update": updatedTodo,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) FindByID(c echo.Context) (err error) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	todo, err := h.service.FindByID(ID)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusNotFound, errResponse)
	}

	return c.JSON(http.StatusOK, todo)
}
