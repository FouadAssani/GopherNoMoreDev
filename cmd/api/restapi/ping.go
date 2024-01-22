package restapi

import (
	"net/http"

	gophernomoredev_api "GopherNoMoreDev/cmd/api/restapi/openapi"

	"github.com/gin-gonic/gin"
)

func (s *server) Ping(c *gin.Context) {
	resp := gophernomoredev_api.PingOutput("pong")
	c.String(http.StatusOK, resp)
}
