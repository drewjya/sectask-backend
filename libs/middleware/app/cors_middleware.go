package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORS
func (cfg *CfgMiddlerware) SetCorsMiddlewares(e *echo.Group, scope string) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType, echo.HeaderAccept, "x-api-key", echo.HeaderAccessControlAllowOrigin},
		AllowMethods: middleware.DefaultCORSConfig.AllowMethods,
	}))
	// e.Use(cfg.checkApiKey)
	if scope == "private" {
		e.Use(cfg.checkPrivateApiKey)
	}
	if scope == "public" {
		e.Use(cfg.checkPublicApiKey)
	}
}

func (cfg *CfgMiddlerware) checkApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		api_key := c.Request().Header.Get("x-api-key")

		if api_key == "" || api_key != cfg.PrivateKey {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Api Key")
		}

		return next(c)
	}
}

func (cfg *CfgMiddlerware) checkPrivateApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		api_key := c.Request().Header.Get("x-api-key")

		if api_key == "" || api_key != cfg.PrivateKey {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Api Key")
		}

		return next(c)
	}
}

func (cfg *CfgMiddlerware) checkPublicApiKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		api_key := c.Request().Header.Get("x-api-key")

		if api_key == "" || api_key != cfg.PublicKey {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Api Key")
		}

		return next(c)
	}
}
