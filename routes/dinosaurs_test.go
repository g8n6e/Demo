package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"prizepicks/jurassicpark/handlers"
	"prizepicks/jurassicpark/models"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDinosaurs(t *testing.T) {
	router := SetupTestRouter()
	req, _ := http.NewRequest("GET", "/dinosaurs/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var dinosaurs []models.Dinosaur
	json.Unmarshal(w.Body.Bytes(), &dinosaurs)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, dinosaurs)
}

func TestGetDinosaurById(t *testing.T) {
	models.ConnectDatabase("../demo.db")
	dinosaurs, err := handlers.GetDinosaurs()
	if err != nil {
		assert.Fail(t, "error getting dinosaurs")
	}
	if dinosaurs != nil {
		router := SetupTestRouter()
		req, _ := http.NewRequest("GET", "/dinosaur/"+strconv.Itoa(dinosaurs[0].ID), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		var dinosaur models.Dinosaur
		json.Unmarshal(w.Body.Bytes(), &dinosaur)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, dinosaur)
		return
	}
	t.Skip("No dinosaurs to test")
}

func TestUpdateDinosaur(t *testing.T) {
	models.ConnectDatabase("../demo.db")
	dinosaurs, err := handlers.GetDinosaurs()
	if err != nil {
		assert.Fail(t, "error getting dinosaurs")
	}
	if dinosaurs != nil {
		router := SetupTestRouter()
		r := strings.NewReader("{\"cageId\":1}")
		req, _ := http.NewRequest("PATCH", "/dinosaur/"+strconv.Itoa(dinosaurs[0].ID), r)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		var dinosaur models.Dinosaur
		json.Unmarshal(w.Body.Bytes(), &dinosaur)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, dinosaur)
		return
	}
	t.Skip("No dinosaurs to test")
}
