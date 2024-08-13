package handler

import (
	"github.com/labstack/echo/v4"
	"learning-project/internal/app"
	"learning-project/internal/module/user/interfaces"
	"learning-project/internal/module/user/payload"
	"learning-project/internal/server/middleware"
	"net/http"
)

type AttendanceHandler struct {
	service interfaces.AttendanceService
}

func NewAttendanceHandler(service interfaces.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{
		service: service,
	}
}

func (h *AttendanceHandler) Create(c echo.Context) error {
	currentUser := middleware.GetCurrentUser(c)

	pyld := payload.AttendanceCreate{
		UserID: currentUser.ID,
	}

	if err := c.Bind(&pyld); err != nil {
		return c.JSON(http.StatusBadRequest, app.ResponseFailed(err.Error()))
	}

	if err := c.Validate(pyld); err != nil {
		return c.JSON(http.StatusBadRequest, app.ResponseFailed(err.Error()))
	}

	result, msg, err := h.service.Create(c.Request().Context(), pyld)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, app.ResponseFailed(msg))
	}

	return c.JSON(http.StatusOK, app.ResponseSuccess(result, msg))
}

func (h *AttendanceHandler) Get(c echo.Context) error {
	currentUser := middleware.GetCurrentUser(c)

	pyld := payload.AttendanceGet{
		UserID:   &currentUser.ID,
		WithUser: true,
	}

	if err := c.Bind(&pyld); err != nil {
		return c.JSON(http.StatusBadRequest, app.ResponseFailed(err.Error()))
	}

	result, msg, err := h.service.Get(c.Request().Context(), pyld)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, app.ResponseFailed(msg))
	}

	return c.JSON(http.StatusOK, app.ResponseSuccess(result, msg))
}
