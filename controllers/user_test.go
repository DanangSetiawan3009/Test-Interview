package controllers

import (
	"bytes"
	"encoding/json"
	"kredit-plus/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisUser(t *testing.T) {
	r := gin.Default()

	userController := userController{}

	r.POST("/regis", userController.Regis)

	data := models.User{
		Fullname: "test",
		Email:    "test",
		Password: "test",
	}
	payload, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/regis", bytes.NewBuffer(payload))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
