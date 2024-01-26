package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	AccessKey string `envconfig:"ACCESS_KEY"`
	SecretKey string `envconfig:"SECRET_KEY"`
	SSL       bool   `envconfig:"SSL"`
	EndPoint  string `envconfig:"END_POINT"`
	Bucket    string `envconfig:"BUCKET"`
}

func InitMinio(mCfg MinioConfig) (*minio.Client, error) {
	return minio.New(mCfg.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(mCfg.AccessKey, mCfg.SecretKey, ""),
		Secure: mCfg.SSL,
	})
}
