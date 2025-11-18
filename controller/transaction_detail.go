package controller

import (
	"bangkit-cell/config"
	"bangkit-cell/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTransactionDetails(c *gin.Context) {
	var details []model.TransactionDetail
	
	rows, err := config.DB.Query("SELECT id, transaction_id, device_service_variant_id, harga, harga_modal, created_at, updated_at FROM transaction_details")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
			"status": 0,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var detail model.TransactionDetail
		err := rows.Scan(
			&detail.ID,
			&detail.TransactionID,
			&detail.DeviceServiceVariantID,
			&detail.Harga,
			&detail.HargaModal,
			&detail.CreatedAt,
			&detail.UpdatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": err.Error(),
				"status": 0,
			})
			return
		}
		details = append(details, detail)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": details,
		"status": 1,
	})
}