package handler

import (
	"learning-project/internal/app"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/payload"
	"learning-project/internal/server/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service interfaces.UserService
}

func NewUserHandler(service interfaces.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Login(c echo.Context) error {
	requestBody := payload.LoginRequest{}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, app.ResponseFailed(err.Error()))
	}

	if err := c.Validate(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, app.ResponseFailed(err.Error()))
	}

	result, msg, err := h.service.Login(c.Request().Context(), requestBody)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, app.ResponseFailed(msg))
	}

	return c.JSON(http.StatusOK, app.ResponseSuccess(result, msg))
}

func (h *UserHandler) Update(c echo.Context) error {
	currentUser := middleware.GetCurrentUser(c)

	request := payload.UserUpdate{
		ID: currentUser.ID,
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, app.ResponseFailed(err.Error()))
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, app.ResponseFailed(err.Error()))
	}

	result, msg, err := h.service.Update(c.Request().Context(), request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, app.ResponseFailed(msg))
	}

	return c.JSON(http.StatusOK, app.ResponseSuccess(result, msg))
}
