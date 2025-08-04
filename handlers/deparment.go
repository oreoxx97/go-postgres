package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-postgres/models"
)

// GetDepartments godoc
// @Summary ดึงข้อมูลแผนกทั้งหมด
// @Description คืนค่ารายชื่อแผนกจากฐานข้อมูล
// @Tags Department
// @Produce json
// @Success 200 {array} models.Department
// @Failure 500 {object} handlers.ErrorResponse
// @Router /department [get]
func GetDepartments(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		departments, err := models.Departments().All(ctx, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, departments)
	}
}
