package app

import (
	"sectask/libs/jwe"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Middlewares interface {
	Middlewares(e *echo.Group, scope string)
	SetCorsMiddlewares(e *echo.Group, scope string)
	// SetJwtMiddlewares(g *echo.Group)
	Debug(info string, data interface{})
	SetCompleteLogMiddleware(e *echo.Group)
	SetMainMiddlewares(e *echo.Group)
	VerifyUser(next echo.HandlerFunc) echo.HandlerFunc
	VerifyBasiAuth(next echo.HandlerFunc) echo.HandlerFunc
}

type (
	CfgMiddlerware struct {
		TokenSecret string
		ApiKey
		Version string
		Jwe     jwe.Credential
	}

	ApiKey struct {
		PrivateKey string
		PublicKey  string
	}

	Log struct {
		Logger *logrus.Logger
	}
)

func New(
	TokenSecret string,
	ApiKey ApiKey,
	Version string,
	Jwe jwe.Credential,
) Middlewares {
	return &CfgMiddlerware{
		TokenSecret: TokenSecret,
		ApiKey:      ApiKey,
		Version:     Version,
		Jwe:         Jwe,
	}
}
