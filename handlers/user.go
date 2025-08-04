package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-postgres/models"
)

// GetUsers godoc
// @Summary ดึงข้อมูลผู้ใช้งานทั้งหมด
// @Description คืนค่าผู้ใช้งานทั้งหมดจากฐานข้อมูล
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} handlers.ErrorResponse
// @Router /users [get]
func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		users, err := models.Users().All(ctx, db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

