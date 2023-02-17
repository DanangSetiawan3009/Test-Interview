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

func TestCreateTransaksiSuccess(t *testing.T) {
	r := gin.Default()
	controller := transaksiController{}
	transaksi := models.Transaksi{
		LimitKonsumenId: 1,
		NoKontrak:       "912637221",
		Otr:             500000,
		AdminFee:        200000,
		JumlahCicilan:   500000,
		JumlahBunga:     100000,
		NamaAsset:       "Setrika",
	}

	r.POST("/transaksi", controller.CreateTransaksi)
	data, _ := json.Marshal(transaksi)
	req, err := http.NewRequest("POST", "/transaksi", bytes.NewBuffer(data))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}
