package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-postgres/models"
)

// GetCompanies godoc
// @Summary ดึงข้อมูลบริษัททั้งหมด
// @Description คืนค่ารายชื่อบริษัทจากฐานข้อมูล
// @Tags Company
// @Produce json
// @Success 200 {array} models.Company
// @Failure 500 {object} handlers.ErrorResponse
// @Router /company [get]
func GetCompanies(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		companies, err := models.Companies().All(ctx, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, companies)
	}
}

