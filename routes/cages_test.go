package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"prizepicks/jurassicpark/handlers"
	"prizepicks/jurassicpark/models"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCages(t *testing.T) {
	router := SetupTestRouter()
	req, _ := http.NewRequest("GET", "/cages/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var cages []models.Cage
	json.Unmarshal(w.Body.Bytes(), &cages)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, cages)
}

func TestGetCageById(t *testing.T) {
	models.ConnectDatabase("../demo.db")
	cages := handlers.GetCages()
	if cages != nil {
		router := SetupTestRouter()
		req, _ := http.NewRequest("GET", "/cage/"+strconv.Itoa(cages[0].ID), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		var cage models.Cage
		json.Unmarshal(w.Body.Bytes(), &cage)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, cage)
		return
	}
	t.Skip("No cages to test")
}
