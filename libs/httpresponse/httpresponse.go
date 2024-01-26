package httpresponse

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

var defaultErrorResp []byte

// Response ...
type Response struct {
	Status      bool        `json:"status"`
	Data        interface{} `json:"data"`
	Meta        interface{} `json:"meta"`
	Code        int         `json:"stat_code"`
	Messages    string      `json:"stat_msg"`
	ErrResponse interface{} `json:"err_response,omitempty"`
}

type ErrMessage struct {
	ErrMapping error
	Message    string
}

// Success wraps and writes the data to the echo.Context.
func Success(w echo.Context, httpStatusCode int, data interface{}) {
	resp := Response{
		Status:      true,
		Data:        data,
		ErrResponse: "",
		Code:        httpStatusCode,
		Meta:        struct{}{},
		Messages:    "Success",
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		jsonResp = defaultErrorResp
		httpStatusCode = http.StatusInternalServerError
		resp.Code = httpStatusCode
	}

	w.Response().Header().Set("Content-Type", "application/json")
	w.Response().WriteHeader(httpStatusCode)
	w.Response().Write(jsonResp)
}

func SuccessWithMeta(w echo.Context, httpStatusCode int, data interface{}, meta interface{}) {
	resp := Response{
		Status:      true,
		Data:        data,
		ErrResponse: "",
		Code:        httpStatusCode,
		Meta:        meta,
		Messages:    "Success",
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		jsonResp = defaultErrorResp
		httpStatusCode = http.StatusInternalServerError
		resp.Code = httpStatusCode
	}

	w.Response().Header().Set("Content-Type", "application/json")
	w.Response().WriteHeader(httpStatusCode)
	w.Response().Write(jsonResp)
}

// Error wraps and writes the errors to the echo.Context.
func Error(w echo.Context, httpStatusCode int, errs ErrMessage) {
	resp := Response{
		Status:      false,
		Data:        struct{}{},
		ErrResponse: errs.Message,
		Code:        httpStatusCode,
		Meta:        struct{}{},
		Messages:    errs.ErrMapping.Error(),
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		jsonResp = defaultErrorResp
		httpStatusCode = http.StatusInternalServerError
	}

	w.Response().Header().Set("Content-Type", "application/json")
	w.Response().WriteHeader(httpStatusCode)
	w.Response().Write(jsonResp)
}
