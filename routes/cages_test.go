package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"prizepicks/jurassicpark/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCagesHandler(t *testing.T) {
	router := SetupTestRouter()
	req, _ := http.NewRequest("GET", "/cages/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var cages []models.Cage
	json.Unmarshal(w.Body.Bytes(), &cages)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, cages)
}
