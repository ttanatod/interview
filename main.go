package main

import (
	"fmt"
	"log"

	//	"net/http"
	"os"

	"example.com/m/v2/models"
	routers "example.com/m/v2/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//connect DB
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PORT := os.Getenv("DB_PORT")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	PORT := os.Getenv("PORT")

	// dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", DB_HOST, DB_USER, DB_NAME, DB_PORT, DB_PASSWORD)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("Can't create conncetion to DB")
	}

	migrateError := db.AutoMigrate(&models.Member{}, &models.Field{}, &models.RentRecord{})
	if migrateError != nil {
		fmt.Printf("migrateError: %v\n", migrateError)
	}

	// request handler
	router := gin.New()
	api := router.Group("/api")
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HELLO GOLANG RESTFUL API.",
		})
	})

	routers.SetCollectionRoutes(api, db)
	port := fmt.Sprintf(":%v", PORT)
	fmt.Println("Server Running on Port", port)
	//http.ListenAndServe(port, router)
	router.Run(port)
}
