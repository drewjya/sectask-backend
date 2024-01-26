package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (cfg *CfgMiddlerware) Middlewares(e *echo.Group, scope string) {
	cfg.SetCorsMiddlewares(e, scope)
	cfg.SetCompleteLogMiddleware(e)
	cfg.SetMainMiddlewares(e)
}

// MAIN
func (cfg *CfgMiddlerware) SetMainMiddlewares(e *echo.Group) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../static",
	}))

	e.Use(cfg.serverHeader)
	e.Use(cfg.VerifyUser)
	// e.Use(cfg.checkApiKey)
}

func (cfg *CfgMiddlerware) serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "sectask-api/"+cfg.Version)
		return next(c)
	}
}
