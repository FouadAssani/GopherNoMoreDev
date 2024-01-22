package restapi

import (
	gophernomoredev_api "GopherNoMoreDev/cmd/api/restapi/openapi"
	"GopherNoMoreDev/internal/application"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

const (
	ShustdownTimeout = 5 * time.Second
)

type Server interface {
	io.Closer
	gophernomoredev_api.ServerInterface

	Start() error
	GetRouter() *gin.Engine
}

type server struct {
	port            int
	log             *logger.Logger
	router          *gin.Engine
	httpServer      *http.Server
	application_api application.GopherNoMoreDevAPI
}

func NewServer(
	log *logger.Logger,
	publicURL string,
	port int,
	application_api application.GopherNoMoreDevAPI,

) (*server, error) {
	s := &server{
		log:             log,
		port:            port,
		router:          gin.New(),
		application_api: application_api,
	}

	gin.SetMode(gin.ReleaseMode)

	s.router.NoRoute(func(ctx *gin.Context) {
		fmt.Println(ctx.Request.URL)

		ctx.JSON(http.StatusNotFound, errorOutput{Code: http.StatusNotFound, Message: "not found"})
	})

	s.router = gophernomoredev_api.RegisterHandlersWithOptions(s.router, s, gophernomoredev_api.GinServerOptions{BaseURL: "/api/v0"})

	s.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", s.port),
		Handler:           s.router,
		ReadHeaderTimeout: 2 * time.Second,
	}

	return s, nil
}

type errorOutput struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *server) Start() error {
	s.log.Infof("listening on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *server) Close() error {
	s.log.Infof("shutting down service on %s", s.httpServer.Addr)

	ctx, cancel := context.WithTimeout(context.Background(), ShustdownTimeout)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}

func (s *server) GetRouter() *gin.Engine {
	return s.router
}
