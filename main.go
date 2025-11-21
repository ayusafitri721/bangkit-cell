package main

import (
	"bangkit-cell/config"
	"bangkit-cell/controller"
	
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	
	router := gin.Default()
	
	// Brand routes
	router.GET("/brands", controller.GetAllBrands)
	router.POST("/brands", controller.SetBrand)
	
	// Device routes
	router.GET("/devices", controller.GetAllDevices)
	
	// Service routes
	router.GET("/services", controller.GetAllServices)
	
	// Transaction routes
	router.GET("/transactions", controller.GetAllTransactions)
	
	// User routes
	router.GET("/users", controller.GetAllUsers)
	
	// Device Service Variant routes
	router.GET("/device-service-variants", controller.GetAllDeviceServiceVariants)
	
	// Transaction Detail routes
	router.GET("/transaction-details", controller.GetAllTransactionDetails)
	
	// Price Log routes
	router.GET("/price-logs", controller.GetAllPriceLogs)
	
	router.Run(":8000")
}