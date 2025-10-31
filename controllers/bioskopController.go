package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"SB-GO-BATCH-71-RAHMA/models"
)

type BioskopController struct {
	DB *sql.DB
}

func (ctrl *BioskopController) CreateBioskop(c *gin.Context) {
	var input models.BioskopInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
		INSERT INTO bioskop (nama, lokasi, rating)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int

	err := ctrl.DB.QueryRow(
		sqlStatement,
		input.Nama,
		input.Lokasi,
		input.Rating,
	).Scan(&id)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Printf("Error query: %v\n", err)
		return
	}

	newBioskop := models.Bioskop{
		ID:     id,
		Nama:   input.Nama,
		Lokasi: input.Lokasi,
		Rating: input.Rating,
	}

	c.JSON(http.StatusCreated, newBioskop)
}