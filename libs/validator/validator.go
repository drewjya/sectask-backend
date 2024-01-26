package validator

import (
	"encoding/json"
	"errors"
	"sectask/libs/str"

	"github.com/go-playground/validator"
)

type (
	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	var err error
	var checkType interface{}
	byteData, _ := json.Marshal(i)
	json.Unmarshal(byteData, &checkType)
	switch checkType.(type) {
	case []interface{}:
		err = cv.Validator.Var(i, "dive")
	case map[string]interface{}:
		err = cv.Validator.Struct(i)
	}
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errKey := str.Underscore(err.StructField())
			if errKey != "" {
				return errors.New(errKey + " " + err.Tag())
			}
		}
	}
	return nil
}
