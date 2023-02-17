package controllers

import (
	"bytes"
	"encoding/json"
	"kredit-plus/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomerSuccess(t *testing.T) {
	r := gin.Default()
	controller := customerController{}
	customer := models.Konsumen{
		Nik:          "223123118274",
		Fullname:     "test",
		Legalname:    "testing",
		TempatLahir:  "test",
		TanggalLahir: &time.Time{},
		Gaji:         1000000,
		FotoKtp:      "test",
		FotoSelfie:   "test",
	}

	r.POST("/customer", controller.CreateCustomer)
	data, _ := json.Marshal(customer)
	req, err := http.NewRequest("POST", "/customer", bytes.NewBuffer(data))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}
