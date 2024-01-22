package main

import (
	"GopherNoMoreDev/cmd/api/restapi"
	gophernomoredev_api "GopherNoMoreDev/cmd/api/restapi/openapi"
	"GopherNoMoreDev/internal/application"
	"io"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type Application interface {
	gophernomoredev_api.ServerInterface

	Start() error
	Close() error
}

type app struct {
	log *logger.Logger

	application Application
	routines    []io.Closer
}

func NewApplication() (*app, error) {
	gin.SetMode(gin.ReleaseMode)

	var log = &logger.Logger{
		Out:       os.Stderr,
		Formatter: new(logger.TextFormatter),
		Hooks:     make(logger.LevelHooks),
		Level:     logger.InfoLevel,
	}

	api := application.NewGopherNoMoreDevAPI(log)

	server, err := restapi.NewServer(log, "http://localhost:8080", 8080, api)
	if err != nil {
		return nil, err
	}

	return &app{
		log:         log,
		application: server,
		routines:    []io.Closer{server},
	}, nil
}

func (a *app) Start() chan error {
	errCh := make(chan error)

	go func() {
		err := a.application.Start()
		if err != nil {
			errCh <- err
			return
		}
	}()

	return errCh
}

func (a *app) Stop() {
	for _, r := range a.routines {
		err := r.Close()
		if err != nil {
			a.log.WithError(err).Errorf("failed to close %s routine", reflect.TypeOf(r).Name())
		}
	}
}
