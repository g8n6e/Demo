package handlers

import (
	"prizepicks/jurassicpark/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCagesHandler(t *testing.T) {
	models.ConnectDatabase("../demo.db")
	cages := GetCages()
	assert.NotEmpty(t, cages)
}
