package entity

import (
	"sectask/service"

	"github.com/labstack/echo/v4"
)

type EntityHandler interface {
	GetEntityHandler(c echo.Context) error
}

type (
	rest struct {
		svc service.Service
	}
)

func NewEntityHandler(
	svc service.Service,
) *rest {
	return &rest{
		svc: svc,
	}
}
