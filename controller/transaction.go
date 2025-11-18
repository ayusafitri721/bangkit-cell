package controller

import (
	"bangkit-cell/config"
	"bangkit-cell/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	var transactions []model.Transaction
	
	rows, err := config.DB.Query("SELECT id, id_operator, status, metode_pembayaran, jumlah_bayar, kembalian, qris_reference, total, customer_name, customer_phone, created_at, updated_at FROM transactions")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
			"status": 0,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(
			&transaction.ID, 
			&transaction.IDOperator, 
			&transaction.Status, 
			&transaction.MetodePembayaran, 
			&transaction.JumlahBayar, 
			&transaction.Kembalian, 
			&transaction.QrisReference, 
			&transaction.Total, 
			&transaction.CustomerName, 
			&transaction.CustomerPhone, 
			&transaction.CreatedAt, 
			&transaction.UpdatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": err.Error(),
				"status": 0,
			})
			return
		}
		transactions = append(transactions, transaction)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": transactions,
		"status": 1,
	})
}
