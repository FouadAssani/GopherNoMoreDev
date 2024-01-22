package restapi

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed openapi.yml
var openAPIFile string

func (s *server) GetOpenapi(c *gin.Context) {
	c.Status(http.StatusOK)
	fmt.Fprintln(c.Writer, openAPIFile)
}
