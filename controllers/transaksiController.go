package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"kredit-plus/models"
	"kredit-plus/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type transaksiController struct {
	service services.TransaksiService
}

func NewTransaksiController(service services.TransaksiService) *transaksiController {
	return &transaksiController{service}
}

func (c *transaksiController) CreateTransaksi(ctx *gin.Context) {
	var transaksiRequest models.Transaksi
	err := ctx.ShouldBindJSON(&transaksiRequest)
	if err != nil {
		errorMssg := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("Error on field: %s, Condition: %s", e.Field(), e.ActualTag())
			errorMssg = append(errorMssg, errorMsg)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMssg,
		})
		return
	}

	transaksi, err := c.service.Create(transaksiRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": transaksi,
	})
}

func (c *transaksiController) GetAllTransaksi(ctx *gin.Context) {
	transaksi, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": transaksi,
	})
}

func (c *transaksiController) GetByIdTransaksi(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	transaksi, err := c.service.FindByID(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": transaksi,
	})
}

func (c *transaksiController) UpdateTransaksi(ctx *gin.Context) {
	var transaksiRequest models.Transaksi
	err := ctx.ShouldBindJSON(&transaksiRequest)
	if err != nil {
		errorMssg := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("Error on field: %s, Condition: %s", e.Field(), e.ActualTag())
			errorMssg = append(errorMssg, errorMsg)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMssg,
		})
		return
	}

	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	transaksi, err := c.service.Update(id, transaksiRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": transaksi,
	})
}

func (c *transaksiController) DeleteTransaksi(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	transaksi, err := c.service.Delete(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": transaksi,
	})
}
