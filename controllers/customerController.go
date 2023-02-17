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

type customerController struct {
	service services.CustomerService
}

func NewCustomerController(service services.CustomerService) *customerController {
	return &customerController{service}
}

// CONTROLLER FOR CUSTOMER
func (c *customerController) CreateCustomer(ctx *gin.Context) {
	var custRequest models.Konsumen
	err := ctx.ShouldBindJSON(&custRequest)
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

	cust, err := c.service.Create(custRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": cust,
	})
}

func (c *customerController) GetAllCustomer(ctx *gin.Context) {
	customers, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": customers,
	})
}

func (c *customerController) GetByIdCustomer(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	cust, err := c.service.FindByID(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": cust,
	})
}

func (c *customerController) UpdateCustomer(ctx *gin.Context) {
	var custRequest models.Konsumen
	err := ctx.ShouldBindJSON(&custRequest)
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
	cust, err := c.service.Update(id, custRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": cust,
	})
}

func (c *customerController) DeleteCustomer(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	cust, err := c.service.Delete(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": cust,
	})
}
