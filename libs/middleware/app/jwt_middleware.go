package app

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	errs "sectask/libs/errors"
	"sectask/libs/httpresponse"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// verify jwt middleware
func (cfg CfgMiddlerware) verify(c echo.Context, role string) (res map[string]interface{}, err error) {
	claims := &jwt.StandardClaims{}

	// header := ctx.Get("Authorization")
	header := c.Request().Header.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		return res, errors.New(errs.HeaderNotPresent)
	}

	// check claims and signing method
	token := strings.Replace(header, "Bearer ", "", -1)
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := cfg.TokenSecret
		return []byte(secret), nil
	})
	if err != nil {
		return res, errors.New(errs.InvalidToken)
	}

	// check token live time
	if claims.ExpiresAt < time.Now().Unix() {
		return res, errors.New(errs.ExpiredToken)
	}
	// jwe roll back encrypted id
	res, err = cfg.Jwe.Rollback(claims.Id)
	if err != nil {
		return res, errors.New(errs.Unauthorized)
	}
	if res == nil {
		return res, errors.New(errs.Unauthorized)
	}
	if role != "" && fmt.Sprintf("%v", res["role"]) != role {
		return res, errors.New(errs.InvalidRole)
	}

	return res, nil
}

// VerifyUser jwt middleware
func (cfg *CfgMiddlerware) VerifyUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		jweRes, err := cfg.verify(ctx, "")
		if err != nil {
			httpresponse.Error(ctx, http.StatusUnauthorized, httpresponse.ErrMessage{
				ErrMapping: err,
				Message:    errs.ErrValidateTokenMSG,
			})
			return err
		}

		// set id to uce case contract
		ctx.Set("user_id", jweRes["id"].(string))

		return next(ctx)
	}
}
