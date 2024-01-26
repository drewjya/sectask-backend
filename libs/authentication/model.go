package authentication

import (
	"github.com/go-redis/redis"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthenticationConfig struct {
	Db          *gorm.DB
	RedisClient *redis.Client
	Ctx         echo.Context
	SRClient    *resty.Client
}

type AuthenticationResponse struct {
	UserID            string `json:"user_id"`
	ClientCode        string `json:"client_code"`
	AppVersion        string `json:"app_version"`
	DeviceType        string `json:"device_type"`
	DeviceCode        string `json:"device_code"`
	SamuelAccessToken string `json:"samuel_access_token"`
}
