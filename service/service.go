package service

import (
	"sectask/repository/psql"
	"sectask/service/s_entity"

	"github.com/go-redis/redis"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

type Service struct {
	Entity s_entity.EntityService
}

func New(
	validator echo.Validator,
	psqlRepos *psql.PsqlRepositories,
	redis *redis.Client,
	srClient *resty.Client,
	samuelRedis *redis.Client,
	minio *minio.Client,
	samuelUrl string,
	samuelUrl2 string,
) Service {
	svc := Service{
		Entity: *s_entity.NewEntityService(validator, redis, samuelRedis, psqlRepos, srClient, minio, samuelUrl, samuelUrl2),
	}

	return svc
}
