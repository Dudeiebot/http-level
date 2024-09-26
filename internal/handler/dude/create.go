package dude

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Dudeiebot/http-level/internal/model"
	"github.com/Dudeiebot/http-level/internal/service"
)

type CreateGopherHandler struct {
	service *service.GopherService
}

func NewCreateGopherHandler(service *service.GopherService) *CreateGopherHandler {
	return &CreateGopherHandler{
		service: service,
	}
}

func (h *CreateGopherHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		gopher := new(model.Dude)
		if err := c.Bind(gopher); err != nil {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				fmt.Sprintf("cannot bind gopher: %v", err),
			)
		}

		err := h.service.Create(c.Request().Context(), gopher)
		if err != nil {
			return fmt.Errorf("cannot create gopher: %w", err)
		}

		return c.JSON(http.StatusCreated, gopher)
	}
}
