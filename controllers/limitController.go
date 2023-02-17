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

type limitController struct {
	service services.LimitService
}

func NewLimitController(service services.LimitService) *limitController {
	return &limitController{service}
}

func (c *limitController) CreateLimit(ctx *gin.Context) {
	var limitRequest models.LimitKonsumen
	err := ctx.ShouldBindJSON(&limitRequest)
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

	limit, err := c.service.Create(limitRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": limit,
	})
}

func (c *limitController) GetAllLimit(ctx *gin.Context) {
	limits, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": limits,
	})
}

func (c *limitController) GetByIdLimit(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	limit, err := c.service.FindByID(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": limit,
	})
}

func (c *limitController) UpdateLimit(ctx *gin.Context) {
	var limitRequest models.LimitKonsumen
	err := ctx.ShouldBindJSON(&limitRequest)
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
	limit, err := c.service.Update(id, limitRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": limit,
	})
}

func (c *limitController) DeleteLimit(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	limit, err := c.service.Delete(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": limit,
	})
}
