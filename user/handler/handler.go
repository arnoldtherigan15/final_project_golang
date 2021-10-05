package handler

import (
	"net/http"
	"strconv"

	"github.com/arnoldtherigan15/final_project_golang/domain"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service domain.UserService
}

func NewHandler(group *echo.Group, service domain.UserService) {
	handler := &Handler{service}
	g := group.Group("/users")
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}

func isRequestValid(user *domain.User) (bool, error) {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return false, err
	}
	return true, nil
}

func (h *Handler) Create(c echo.Context) (err error) {
	var user domain.User
	if err = c.Bind(&user); err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	if ok, err := isRequestValid(&user); !ok {
		errMsg := domain.ErrorValidationFormatter(err)
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, errMsg)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	createdUser, err := h.service.Create(&user)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusInternalServerError, domain.ErrInternalServerError)
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	return c.JSON(http.StatusCreated, createdUser)
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
	var user domain.User
	err = c.Bind(&user)
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, domain.ErrBadRequest)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	if ok, err := isRequestValid(&user); !ok {
		errMsg := domain.ErrorValidationFormatter(err)
		errResponse := domain.ErrorFormatter(http.StatusUnprocessableEntity, errMsg)
		return c.JSON(http.StatusUnprocessableEntity, errResponse)
	}

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusBadRequest, domain.ErrBadParamInput)
		return c.JSON(http.StatusBadRequest, errResponse)
	}
	updatedTodo, err := h.service.Update(ID, &user)

	if err != nil {
		errResponse := domain.ErrorFormatter(http.StatusNotFound, domain.ErrNotFound)
		return c.JSON(http.StatusNotFound, errResponse)
	}

	response := map[string]bool{
		"is_update": updatedTodo,
	}

	return c.JSON(http.StatusOK, response)
}
