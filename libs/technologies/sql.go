package technologies

import (
	// "database/sql"

	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type DB struct {
	// Connection *sqlx.DB
	Connection *gorm.DB
}

type DBConfig struct {
	Dialect         string        `envconfig:"DIALECT"`
	Host            string        `envconfig:"HOST"`
	Port            int           `envconfig:"PORT"`
	Name            string        `envconfig:"NAME"`
	Username        string        `envconfig:"USER"`
	Password        string        `envconfig:"PASSWORD"`
	MaxConnOpen     int           `envconfig:"MAX_CONN_OPEN"`
	MaxConnLifetime time.Duration `envconfig:"MAX_CONN_LIFETIME"`
	MaxConnIdle     int           `envconfig:"MAX_CONN_IDLE"`
	MigrateLocation string        `envconfig:"MIGRATE_LOCATION"`
}

// var sqlxConnect = sqlx.Connect

// InitDB init the database from config to database connection
func InitDB(cfg DBConfig) (*gorm.DB, *sql.DB, error) {
	log.Println(cfg)
	dsn, err := cfg.toDSN()
	if err != nil {
		return nil, nil, err
	}
	// gorm
	db, err := gorm.Open(cfg.Dialect, dsn)
	if err != nil {
		return db, nil, err
	}
	// sqlx Connect
	// db, err := sqlxConnect(cfg.Dialect, dsn)
	// if err != nil {
	// 	return db, err
	// }
	dba := db.DB()
	err = dba.Ping()
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		return db, dba, err
	}
	dba.SetMaxOpenConns(cfg.MaxConnOpen)
	dba.SetMaxIdleConns(cfg.MaxConnIdle)
	dba.SetConnMaxLifetime(cfg.MaxConnLifetime)
	return db, dba, nil
}

func Migrate(cfg DBConfig, types string) error {
	// https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
	// migrate create -ext sql -dir db/migration -seq init_schema
	dsn, err := cfg.toDSN()
	if err != nil {
		return err
	}

	log.Println("Migrating Up Database Table")
	db, err := sql.Open(cfg.Dialect, dsn)
	if err != nil {
		log.Println(err)
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Println(err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+cfg.MigrateLocation, cfg.Dialect, driver)
	if err != nil {
		log.Println(err)
		return err
	}

	if types == "UP" {
		if err := m.Up(); err != nil {
			log.Println(err)
			return err
		}
		log.Println("Migrating Up success")

	} else if types == "DOWN" {
		if err := m.Down(); err != nil {
			log.Println(err)
			return err
		}
		log.Println("Migrating Down success")
	} else {
		return errors.New("wrong types of migration, only: 'UP' || 'DOWN'")
	}

	db.Close()
	return nil
}

func (cfg DBConfig) toDSN() (string, error) {
	switch cfg.Dialect {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name), nil
	case "postgres":
		// return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name), nil
		return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name), nil
	default:
		return "", fmt.Errorf("Dialect is not supported. expected: (msql|postgres), got: %s", cfg.Dialect)
	}
}
