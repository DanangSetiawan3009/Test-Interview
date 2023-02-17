package controllers

import (
	"kredit-plus/models"
	"kredit-plus/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{userService}
}

// CONTROLLER FOR USER
func (c *userController) Regis(ctx *gin.Context) {
	var userRegis models.User
	err := ctx.ShouldBindJSON(&userRegis)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if err := c.userService.Regis(userRegis); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *userController) Login(ctx *gin.Context) {
	var userLogin models.User
	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	token, err := c.userService.Login(userLogin)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
