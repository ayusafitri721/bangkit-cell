package controller

import (
	"bangkit-cell/config"
	"bangkit-cell/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPriceLogs(c *gin.Context) {
	var logs []model.PriceLog
	
	rows, err := config.DB.Query("SELECT id, device_service_variant_id, user_id, old_harga_min, old_harga_max, new_harga_min, new_harga_max, tipe_part, changed_at FROM price_logs")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
			"status": 0,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var log model.PriceLog
		err := rows.Scan(
			&log.ID,
			&log.DeviceServiceVariantID,
			&log.UserID,
			&log.OldHargaMin,
			&log.OldHargaMax,
			&log.NewHargaMin,
			&log.NewHargaMax,
			&log.TipePart,
			&log.ChangedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": err.Error(),
				"status": 0,
			})
			return
		}
		logs = append(logs, log)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": logs,
		"status": 1,
	})
}