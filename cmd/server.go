package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"sectask/delivery/rest"
	"sectask/libs/jwe"
	mwApp "sectask/libs/middleware/app"
	vldtr "sectask/libs/validator"
	"sectask/repository/psql"
	"sectask/service"

	"sectask/libs/minio"
	dbase "sectask/libs/technologies"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var ProjectFolder = flag.String("folder", "./", "absolute path of project folder")

// GitCommit : for show commit number
var GitCommit string

func serv() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	e := echo.New()
	cfg := loadConfig()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Success Run CA Version %v %v !!!", os.Getenv("CA_APP_VERSION"), os.Getenv("CA_APP_ENV")))
	})

	e.Validator = &vldtr.CustomValidator{Validator: validator.New()}
	log.Println("[" + appName + "] Successfully initializing validator ...")

	db, pg, err := dbase.InitDB(cfg.DB)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// auto migrate DB
	autoMigrateDB(db)

	dbV2, err := gorm.Open(postgres.New(postgres.Config{Conn: pg}))
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("[" + appName + "] Database successfully initialized ...")

	redis := dbase.InitRedis(cfg.Redis)
	log.Println("[" + appName + "] Redis successfully initialized ...")

	samuelRedis := dbase.InitRedis(cfg.SamuelRedis)
	log.Println("[" + appName + "] Samuel Redis successfully initialized ...")

	srClient := dbase.InitRestyClient()
	srClient.SetBaseURL(os.Getenv("SERVER_V2_URL"))
	log.Println("[" + appName + "] SR Client successfully initialized ...")

	psqlRepos := psql.NewPsqlRepositories(db, dbV2)

	minio, err := minio.InitMinio(cfg.Minio)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	svc := service.New(e.Validator, psqlRepos, redis, srClient, samuelRedis, minio, os.Getenv("CA_SAMUEL_URL"), os.Getenv("CA_SAMUEL_URL2"))

	mwApps := mwApp.New(
		cfg.Token.Secret,
		mwApp.ApiKey(cfg.App.ApiKey),
		cfg.App.Version,
		jwe.Credential{
			KeyLocation: cfg.App.PrivateKey.Location,
			Passphrase:  cfg.App.PrivateKey.Passphrase,
		},
	)

	rootURL := "/" + strings.ToLower(appName)
	healthRoute := e.Group(rootURL) // non api-key required
	rest.New(svc, mwApps, cfg.App.Env).Route(healthRoute)

	log.Println("[" + appName + "] Root Url: " + rootURL)
	log.Println("["+appName+"] (build commit hash: "+GitCommit+") running on port:", cfg.Webserver.ListenAddress)
	if cfg.Webserver.ListenAddress != "" {
		e.HidePort = true
		e.HideBanner = true
	}

	s := &http.Server{
		Addr:         cfg.Webserver.ListenAddress,
		ReadTimeout:  cfg.Webserver.ReadTimeout,
		WriteTimeout: cfg.Webserver.WriteTimeout,
	}

	go func() {
		sigchan := make(chan os.Signal, 1)

		signal.Notify(sigchan, os.Interrupt)

		<-sigchan

		db.Close()
		fmt.Println("database has been closed")

		redis.Close()
		fmt.Println("redis has been closed")

		samuelRedis.Close()
		fmt.Println("samuel redis has been closed")

		os.Exit(0)
	}()

	err = e.StartServer(s)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// if failed get /v10
	// go get "github.com/go-playground/validator/v10"
	// cd $GOPATH/src/github.com/go-playground/validator/
	// ln -s . v10

	// if failed get go-urn && universal-translator
	// go get -u "github.com/leodido/go-urn"
	// go get -u "github.com/go-playground/universal-translator"
}
