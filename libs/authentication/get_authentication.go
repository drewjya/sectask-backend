package authentication

import (
	"errors"
	errs "sectask/libs/errors"
	"sectask/libs/httpresponse"
	"sectask/libs/str"
	"sectask/libs/table"
	"sectask/libs/try"
	"strings"
)

func (s *AuthenticationConfig) GetAuthentication() (AuthenticationResponse, httpresponse.ErrMessage) {
	var result AuthenticationResponse
	userID, _ := s.Ctx.Get("user_id").(string)
	authHeader := s.Ctx.Request().Header.Get("Authorization")
	bearerToken := try.ArrayStringToString(strings.Split(authHeader, "Bearer "), 1, "")
	if userID == "" {
		return result, httpresponse.ErrMessage{
			ErrMapping: errors.New(errs.UserNotFound),
			Message:    userID,
		}
	}

	err := s.Db.Debug().Table(table.TableAccountOpenings).Select("user_id,client_code,app_version,device_type,device_code").Where("user_id = ? and deleted_at IS NULL", userID).Find(&result).Error
	if err != nil {
		return result, httpresponse.ErrMessage{
			ErrMapping: errors.New(errs.InternalDBError),
			Message:    err.Error(),
		}
	} else if result.ClientCode == "" {
		return result, httpresponse.ErrMessage{
			ErrMapping: errors.New(errs.ClientCodeNotFound),
			Message:    userID,
		}
	}

	result.SamuelAccessToken = str.TrimDoubleQuote(s.RedisClient.Get("accessToken" + result.ClientCode).Val())
	if result.SamuelAccessToken == "" {
		autoLoginURL := "/v1/api/samuel/auto-login"
		s.SRClient.SetAuthToken(bearerToken).R().Post(autoLoginURL)
		result.SamuelAccessToken = str.TrimDoubleQuote(s.RedisClient.Get("accessToken" + result.ClientCode).Val())
		if result.SamuelAccessToken == "" {
			return result, httpresponse.ErrMessage{
				ErrMapping: errors.New(errs.InvalidToken),
				Message:    result.ClientCode,
			}
		}
	}
	return result, httpresponse.ErrMessage{}
}
