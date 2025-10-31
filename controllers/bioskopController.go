package controllers

import (
	"database/sql"
	"fmt"
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

func (ctrl *BioskopController) GetBioskops(c *gin.Context) {
	sqlStatement := `SELECT id, nama, lokasi, rating FROM bioskop ORDER BY id ASC`

	rows, err := ctrl.DB.Query(sqlStatement)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	bioskops := []models.Bioskop{}

	for rows.Next() {
		var b models.Bioskop
		if err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan data"})
			return
		}
		bioskops = append(bioskops, b)
	}

	c.JSON(http.StatusOK, bioskops)
}


func (ctrl *BioskopController) GetBioskopByID(c *gin.Context) {
	id := c.Param("id")

	sqlStatement := `SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1`

	var b models.Bioskop

	err := ctrl.DB.QueryRow(sqlStatement, id).Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, b)
}


func (ctrl *BioskopController) UpdateBioskop(c *gin.Context) {
	id := c.Param("id")

	var input models.BioskopInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `
		UPDATE bioskop
		SET nama = $1, lokasi = $2, rating = $3
		WHERE id = $4
		RETURNING id, nama, lokasi, rating`

	var updatedBioskop models.Bioskop

	err := ctrl.DB.QueryRow(
		sqlStatement,
		input.Nama,
		input.Lokasi,
		input.Rating,
		id,
	).Scan(&updatedBioskop.ID, &updatedBioskop.Nama, &updatedBioskop.Lokasi, &updatedBioskop.Rating)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "No bioskop found to update"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update data"})
		return
	}

	c.JSON(http.StatusOK, updatedBioskop)
}


func (ctrl *BioskopController) DeleteBioskop(c *gin.Context) {
	id := c.Param("id")

	sqlStatement := `DELETE FROM bioskop WHERE id = $1`

	result, err := ctrl.DB.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete data"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check deletion results"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No bioskop found to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Bioskop with ID %s was successfully deleted", id)})
}