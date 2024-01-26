package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/google/uuid"
	log2 "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// SetCompleteLogMiddleware Middleware for logging request and response
func (cfg *CfgMiddlerware) newLogger() (*Log, error) {
	l := logrus.New()

	logf, err := rotatelogs.New(
		"storage/logs/access_log.%Y%m%d",

		// symlink current log to this file
		rotatelogs.WithLinkName("/tmp/app_access.log"),

		// max: 7 days to keep
		rotatelogs.WithMaxAge(24*7*time.Hour),

		// rotate every day
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return nil, err
	}

	l.Formatter = &logrus.JSONFormatter{}
	l.Out = logf
	l.Level = logrus.InfoLevel

	return &Log{Logger: l}, nil
}

func (cfg *CfgMiddlerware) Debug(info string, data interface{}) {
	l, err := cfg.newLogger()
	if err != nil {
		panic(err)
	}

	_, file, no, _ := runtime.Caller(1)

	l.Logger.WithFields(logrus.Fields{
		"file": file,
		"line": no,
		"time": time.Now().Format("2006-01-02 15:04:05"),
		"data": data,
	}).Info(info)
}

func (cfg *CfgMiddlerware) SetCompleteLogMiddleware(e *echo.Group) {
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		var bodyJSON interface{}
		var bodyRESP interface{}

		json.Unmarshal(reqBody, &bodyJSON)
		json.Unmarshal(resBody, &bodyRESP)

		logField := logrus.Fields{
			"at":            time.Now().Format("2006-01-02 15:04:05"),
			"remote_ip":     c.RealIP(),
			"protocol":      c.Request().Proto,
			"host":          c.Request().Host,
			"uri":           c.Request().RequestURI,
			"mobile_req_id": c.Request().Header.Get("request-id"),
			//"headers":   c.Request().Header,
			"params":   c.Request().URL.Query(),
			"request":  bodyJSON,
			"response": bodyRESP,
			//"request":  string(reqBody),
			//"response": string(resBody),
		}

		// for more complete log
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.WithFields(logField).Info("REQUEST&RESPONSE")
		// Info("")

		// l.Logger.WithFields(logField).Info("REQUEST LOG")
	}))
}

func Log1(level log2.Level, message string, context string, scope string, corr ...interface{}) {
	log2.SetFormatter(&log2.JSONFormatter{})
	// append optional correlation id to logger
	var correlation interface{}
	if len(corr) > 0 {
		correlation = corr[0]
	}
	entry := LogContext(context, scope, correlation)
	switch level {
	case log2.DebugLevel:
		entry.Debug(message)
	case log2.InfoLevel:
		entry.Info(message)
	case log2.WarnLevel:
		entry.Warn(message)
	case log2.ErrorLevel:
		entry.Error(message)
	case log2.FatalLevel:
		entry.Fatal(message)
	case log2.PanicLevel:
		entry.Panic(message)
	}
}

// LogContext function for logging the context of echo
// c string context
// s string scope
func LogContext(c string, s string, cor ...interface{}) *log2.Entry {
	// topic, ok := os.LookupEnv("LOG_TOPIC")
	// if !ok {
	// 	topic = TOPIC
	// }
	entry := log2.WithFields(log2.Fields{
		"context": c,
		"scope":   s,
		"tz":      time.Now().UTC().Format(time.RFC3339),
	})
	if len(cor) > 0 {
		if cor[0] != nil {
			entry = entry.WithFields(
				log2.Fields{
					"req_id": fmt.Sprintf("%+v", cor[0]),
				})
		}
	}
	return entry
}

func SetRqId(ctx context.Context) context.Context {
	rqId, _ := uuid.NewRandom()
	return context.WithValue(ctx, "request_id", rqId.String())
}
