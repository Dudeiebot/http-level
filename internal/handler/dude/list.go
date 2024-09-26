package dude

import (
	"fmt"
	"net/http"

	"github.com/ankorstore/yokai/log"
	"github.com/ankorstore/yokai/trace"
	"github.com/labstack/echo/v4"

	"github.com/Dudeiebot/http-level/internal/service"
)

type ListGophersHandler struct {
	service *service.GopherService
}

func NewListGophersHandler(service *service.GopherService) *ListGophersHandler {
	return &ListGophersHandler{
		service: service,
	}
}

func (h *ListGophersHandler) Handle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		ctx, span := trace.CtxTracerProvider(ctx).
			Tracer("dude-api").
			Start(ctx, "ListDudeHandler span")
		defer span.End()

		log.CtxLogger(ctx).Info().Msg("called ListDudeHandler")

		gophers, err := h.service.List(ctx)
		if err != nil {
			return fmt.Errorf("cannot list all peoples: %w", err)
		}

		return c.JSON(http.StatusOK, gophers)
	}
}
