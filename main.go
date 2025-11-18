package main

import (
	"bangkit-cell/config"
	"bangkit-cell/controller"	
	"github.com/gin-gonic/gin"
)

func main() {
	// Tes database dulu
	config.ConnectDatabase()
	
	// Setup Gin Router
	router := gin.Default()
	
	router.GET("/getallrepairs", controller.GetAllBangkitCell)
	
	router.Run(":8000")

	// v1 := router.Group("/v1")
	// v2 := router.Group("/v2")

	// root "/"
	// v1.GET("/", rootHandler)
	// v2.GET("/", rootHandlerv2)

	// /repair
	// v1.GET("/repair", repairHandler)

	// /repair/:customer/:phone
	// v1.GET("/repair/:customer/:phone", repairWithParamHandler)

	// /service?type=
	// v1.GET("/service", serviceQueryHandler)

	// POST /repair/order
	// v1.POST("/repair/order", postRepairOrderHandler)
}

// func rootHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"app":    "HP Repair Service",
// 		"status": "running",
// 	})
// }

// func rootHandlerv2(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"app":     "HP Repair Service",
// 		"status":  "running",
// 		"version": "2.0",
// 	})
// }

// func repairHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"customer": "Budi Santoso",
// 		"phone":    "Samsung Galaxy S21",
// 		"issue":    "Layar pecah",
// 		"price":    500000,
// 	})
// }

// func repairWithParamHandler(c *gin.Context) {
// 	customer := c.Param("customer")
// 	phone := c.Param("phone")

// 	c.JSON(http.StatusOK, gin.H{
// 		"customer": customer,
// 		"phone":    phone,
// 		"status":   "OK",
// 	})
// }

// func serviceQueryHandler(c *gin.Context) {
// 	serviceType := c.Query("type")

// 	c.JSON(http.StatusOK, gin.H{
// 		"type":   serviceType,
// 		"status": "OK",
// 	})
// }

// type RepairOrderInput struct {
// 	Customer   string `json:"customer" binding:"required"`
// 	Phone      string `json:"phone" binding:"required"`
// 	Issue      string `json:"issue" binding:"required"`
// 	Price      int    `json:"price" binding:"required"`
// }

// func postRepairOrderHandler(c *gin.Context) {
// 	var repairInput RepairOrderInput

// 	err := c.ShouldBindJSON(&repairInput)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"Customer": repairInput.Customer,
// 		"Phone":    repairInput.Phone,
// 		"Issue":    repairInput.Issue,
// 		"Price":    repairInput.Price,
// 	})
// }