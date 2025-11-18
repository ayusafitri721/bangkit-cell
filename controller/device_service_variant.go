package controller

import (
	"bangkit-cell/config"
	"bangkit-cell/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDeviceServiceVariants(c *gin.Context) {
	var variants []model.DeviceServiceVariant
	
	rows, err := config.DB.Query("SELECT id, device_id, service_id, tipe_part, harga_min, harga_max, catatan, created_at, updated_at FROM device_service_variants")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
			"status": 0,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var variant model.DeviceServiceVariant
		err := rows.Scan(
			&variant.ID,
			&variant.DeviceID,
			&variant.ServiceID,
			&variant.TipePart,
			&variant.HargaMin,
			&variant.HargaMax,
			&variant.Catatan,
			&variant.CreatedAt,
			&variant.UpdatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": err.Error(),
				"status": 0,
			})
			return
		}
		variants = append(variants, variant)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": variants,
		"status": 1,
	})
}