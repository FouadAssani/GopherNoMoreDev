package restapi_test

import (
	"GopherNoMoreDev/cmd/api/restapi"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupNewAPIServer(t *testing.T) restapi.Server {
	log := &logger.Logger{}
	publicURL := ""
	port := 8080

	server, err := restapi.NewServer(log, publicURL, port, nil)

	require.Nil(t, err)
	require.NotNil(t, server)

	return server
}

func TestServer_GetRouter(t *testing.T) {

	server := setupNewAPIServer(t)

	router := server.GetRouter()

	assert.NotNil(t, router)
}

func TestCorsMiddleware(t *testing.T) {
	server := setupNewAPIServer(t)

	resp := preflightRequest(server.GetRouter(), "/ping", "http://www.exemple.com", http.MethodGet)
	require.Equal(t, http.StatusNoContent, resp.Code, "body: %s", resp.Body.String())

	headers := []string{
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Methods",
		"Access-Control-Allow-Headers",
		"Access-Control-Max-Age",
	}
	for _, h := range headers {
		require.NotEmpty(t, resp.Header().Get(h), h)
	}
}

func preflightRequest(engine *gin.Engine, path, origin, requestMethod string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodOptions, path, bytes.NewBuffer([]byte{}))
	req.Header.Set("Origin", origin)
	req.Header.Set("Access-Control-Request-Method", requestMethod)

	recorder := httptest.NewRecorder()
	engine.ServeHTTP(recorder, req)

	return recorder
}
