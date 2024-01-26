package main

import (
	"log"
	"time"

	"sectask/domain/entity"
	"sectask/libs/minio"

	db "sectask/libs/technologies"

	"github.com/jinzhu/gorm"
	godotenv "github.com/joho/godotenv"
	envconfig "github.com/kelseyhightower/envconfig"

	_ "github.com/lib/pq"
	// _ "github.com/go-sql-driver/mysql"
)

const (
	appName    = "CA"
	appVersion = "0.0.0.alpha1"
	calldepth  = 4
)

type (
	config struct {
		DB          db.DBConfig       `envconfig:"DB"`
		Webserver   wb                `envconfig:"WEBSERVER"`
		Redis       db.RedisConfig    `envconfig:"REDIS"`
		SamuelRedis db.RedisConfig    `envconfig:"SAMUEL_REDIS"`
		App         appConfig         `envconfig:"APP"`
		Token       token             `envconfig:"TOKEN"`
		Minio       minio.MinioConfig `envconfig:"MINIO"`
	}

	wb struct {
		ListenAddress   string        `envconfig:"LISTEN_ADDRESS"`
		MaxConnections  int           `envconfig:"MAX_CONNECTION"`
		ReadTimeout     time.Duration `envconfig:"READ_TIMEOUT"`
		WriteTimeout    time.Duration `envconfig:"WRITE_TIMEOUT"`
		GracefulTimeout time.Duration `envconfig:"GRACEFUL_TIMEOUT"`
	}
	appConfig struct {
		Env        string     `envconfig:"ENV"`
		Version    string     `envconfig:"VERSION"`
		Secret     string     `envconfig:"SECRET"`
		PrivateKey privateKey `envconfig:"PRIVATE_KEY"`
		ApiKey     apiKey     `envconfig:"API"`
	}

	user struct {
		Client string `envconfig:"CLIENT"`
		Secret string `envconfig:"SECRET"`
		Url    string `envconfig:"URL"`
		ApiKey string `envconfig:"API_KEY"`
	}

	privateKey struct {
		Location   string `envconig:"LOCATION"`
		Passphrase string `envconig:"PASSPHRASE"`
	}

	apiKey struct {
		PrivateKey string `envconig:"PRIVATE_KEY"`
		PublicKey  string `envconig:"PUBLIC_KEY"`
	}

	token struct {
		Secret           string `envconfig:"SECRET"`
		RefreshSecret    string `envconfig:"REFRESH_SECRET"`
		ExpSecret        string `envconfig:"EXP_SECRET"`
		ExpRefreshSecret string `envconfig:"EXP_REFRESH_SECRET"`
	}
)

func load(filenames ...string) error {
	return godotenv.Load(filenames...)
}

// Parse parses environment variables.
func parse(prefix string, out interface{}) error {
	return envconfig.Process(prefix, out)
}

// LoadAndParse loads and parses environment variables.
func LoadAndParse(prefix string, out interface{}, filenames ...string) error {
	if err := load(filenames...); err != nil {
		return err
	}

	if err := parse(prefix, out); err != nil {
		return err
	}

	return nil
}

func loadConfig() *config {
	var cfg config

	err := LoadAndParse(appName, &cfg)
	if err != nil {
		panic("Failed to load environment configuration. err: " + err.Error())
	}
	log.Println("[" + appName + "] env: " + cfg.App.Env + " successfully loaded")

	return &cfg
}

/*
Function to auto migrate DB
*/
func autoMigrateDB(db *gorm.DB) {
	db.AutoMigrate(&entity.Entity{})
}
