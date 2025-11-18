package controller

import (
	"bangkit-cell/config"
	"bangkit-cell/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBangkitCell(c *gin.Context) {
	var bangkitcells []model.Repair
	
	rows, err := config.DB.Query("SELECT * FROM repairs")
	if err != nil {
		println("Gagal masuk ke table repairs")
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
			"status": 0,
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var bangkitcell model.Repair
		err := rows.Scan(
			&bangkitcell.ID, 
			&bangkitcell.Customer, 
			&bangkitcell.Phone, 
			&bangkitcell.Issue, 
			&bangkitcell.Price, 
			&bangkitcell.Status, 
			&bangkitcell.CreatedAt, 
			&bangkitcell.UpdatedAt,
		)

		if err != nil {
			fmt.Println("Error Scan Data Repair")
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": err.Error(),
				"status": 0,
			})
			return
		}

		fmt.Println(bangkitcell)
		bangkitcells = append(bangkitcells, bangkitcell)

		if len(bangkitcells) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"result": bangkitcells,
				"status": 1,
			})
		}
	}
}