package entity

import (
	"errors"
	"net/http"
	errs "sectask/libs/errors"
	"sectask/libs/httpresponse"

	"github.com/labstack/echo/v4"
)

func (re *rest) GetEntityHandler(c echo.Context) error {
	ctx := c.Request().Context()
	id, ok := c.Get("id").(string)
	if !ok {
		httpresponse.Error(c, http.StatusInternalServerError, httpresponse.ErrMessage{
			ErrMapping: errors.New(errs.Unauthorized),
			Message:    errs.Unauthorized,
		})
		return errors.New(errs.Unauthorized)
	}

	res, err := re.svc.Entity.GetEntityService(ctx, id)
	if err != nil {
		httpresponse.Error(c, http.StatusInternalServerError, httpresponse.ErrMessage{
			ErrMapping: err,
			Message:    err.Error(),
		})
		return err
	}

	httpresponse.Success(c, http.StatusOK, res)
	return nil
}
