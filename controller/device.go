package controller

import (
	"bangkit-cell/config"
	"bangkit-cell/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDevices(c *gin.Context) {
	var devices []model.Device
	
	rows, err := config.DB.Query("SELECT id, brand_id, model, tipe, created_at, updated_at FROM devices")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
			"status": 0,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var device model.Device
		err := rows.Scan(&device.ID, &device.BrandID, &device.Model, &device.Tipe, &device.CreatedAt, &device.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": err.Error(),
				"status": 0,
			})
			return
		}
		devices = append(devices, device)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": devices,
		"status": 1,
	})
}