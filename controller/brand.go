package controller

import (
	"bangkit-cell/config"
	"bangkit-cell/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBrands(c *gin.Context) {
	var brands []model.Brand
	
	rows, err := config.DB.Query("SELECT id, nama, negara_asal, created_at, updated_at FROM brands")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
			"status": 0,
		})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var brand model.Brand
		err := rows.Scan(&brand.ID, &brand.Nama, &brand.NegaraAsal, &brand.CreatedAt, &brand.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"result": err.Error(),
				"status": 0,
			})
			return
		}
		brands = append(brands, brand)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": brands,
		"status": 1,
	})
}

func SetBrand(c *gin.Context) {
	var brand model.Brand

	err := c.ShouldBindJSON(&brand)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err,
			"status": http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": brand,
	})
}