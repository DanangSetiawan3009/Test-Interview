package main

import (
	"fmt"
	"kredit-plus/controllers"
	"kredit-plus/repositories"
	"kredit-plus/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:adm!n123@tcp(localhost:3306)/kredit_plus?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection error")
	}
	fmt.Println("DB Connection Success")

	// CUSTOMER
	customerRepository := repositories.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepository)
	customerController := controllers.NewCustomerController(customerService)

	// USER
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	// LIMIT
	limitRepository := repositories.NewLimitRepository(db)
	limitService := services.NewLimitService(limitRepository)
	limitController := controllers.NewLimitController(limitService)

	// TRANSAKSI
	transaksiRepository := repositories.NewTransaksiRepository(db)
	transaksiService := services.NewTransaksiService(transaksiRepository)
	transaksiController := controllers.NewTransaksiController(transaksiService)

	router := gin.Default()
	// user router
	router.POST("/regis", userController.Regis)
	router.POST("/login", userController.Login)
	// customer router
	router.POST("/customer", customerController.CreateCustomer)
	router.GET("/cuxtomer", customerController.GetAllCustomer)
	router.GET("/customer/:id", customerController.GetByIdCustomer)
	router.PUT("/customer/:id", customerController.UpdateCustomer)
	router.DELETE("/customer/:id", customerController.DeleteCustomer)
	// limit_konsumen router
	router.POST("/limit", limitController.CreateLimit)
	router.GET("/limit", limitController.GetAllLimit)
	router.GET("/limit/:id", limitController.GetByIdLimit)
	router.PUT("/limit/:id", limitController.UpdateLimit)
	router.DELETE("/limit/:id", limitController.DeleteLimit)
	// transaksi router
	router.POST("/transaksi", transaksiController.CreateTransaksi)
	router.GET("/transaksi", transaksiController.GetAllTransaksi)
	router.GET("/transaksi/:id", transaksiController.GetByIdTransaksi)
	router.PUT("/transaksi/:id", transaksiController.UpdateTransaksi)
	router.DELETE("/transaksi/:id", transaksiController.DeleteTransaksi)

	router.Run()
}
