package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-postgres/docs"
	"go-postgres/handlers"
)

// @title Go Postgres API
// @version 1.0
// @description API ตัวอย่างเชื่อมต่อ PostgreSQL ด้วย Gin และ Swagger
// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {
	// โหลด .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// อ่านค่าจาก ENV
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, sslMode,
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("เชื่อมต่อ PostgreSQL สำเร็จ!")

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/company", handlers.GetCompanies(db))
	r.GET("/department", handlers.GetDepartments(db))
	r.GET("/users", handlers.GetUsers(db))

	r.Run(":8080")
}
