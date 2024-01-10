package handlers

import (
	"prizepicks/jurassicpark/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCages(t *testing.T) {
	t.Skip() //just for initial validation, initial state will be no cages and should support no cages
	models.ConnectDatabase("../demo.db")
	cages := GetCages()
	assert.NotEmpty(t, cages)
}
