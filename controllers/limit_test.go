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

func TestCreateLimitSuccess(t *testing.T) {
	r := gin.Default()
	controller := limitController{}
	limit := models.LimitKonsumen{
		KonsumenId: 1,
		Tenor1:     100000,
		Tenor2:     200000,
		Tenor3:     500000,
		Tenor4:     700000,
	}

	r.POST("/limit", controller.CreateLimit)
	data, _ := json.Marshal(limit)
	req, err := http.NewRequest("POST", "/limit", bytes.NewBuffer(data))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}
