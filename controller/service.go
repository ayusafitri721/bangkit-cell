package controller

import (
	"bangkit-cell/config"
	"bangkit-cell/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllServices(c *gin.Context) {
	var services []model.Service
	
	rows, err := config.DB.Query("SELECT id, nama, deskripsi, created_at, updated_at FROM services")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
			"status": 0,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var service model.Service
		err := rows.Scan(&service.ID, &service.Nama, &service.Deskripsi, &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": err.Error(),
				"status": 0,
			})
			return
		}
		services = append(services, service)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": services,
		"status": 1,
	})
}