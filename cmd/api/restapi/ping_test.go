package restapi_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	// Créer un nouveau routeur Gin
	router := gin.New()

	// Créer un nouvel objet "server"
	s := setupNewAPIServer(t)

	// Ajouter une route pour la méthode Ping
	router.GET("/ping", s.Ping)

	// Créer une nouvelle requête GET pour l'URL "/ping"
	req, err := http.NewRequest("GET", "/ping", nil)
	require.NoError(t, err, fmt.Sprintf("Error creating request: %v", err))

	// Créer un enregistreur de réponse pour la requête
	w := httptest.NewRecorder()

	// Exécuter la requête à l'aide du routeur Gin
	router.ServeHTTP(w, req)

	// Vérifier que le code de statut HTTP est 200 (OK)
	require.Equal(t, http.StatusOK, w.Code, "unexpected HTTP status code")

	// Vérifier que le corps de la réponse est "pong"
	require.Equal(t, "pong", w.Body.String(), "unexpected response body")
}
