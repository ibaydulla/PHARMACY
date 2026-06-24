package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ibaydulla/internal/controllers"
	"github.com/ibaydulla/internal/utils"
)

func main() {

	// Database connection
	connStr := fmt.Sprintf(
		"user=%s dbname=%s host=%s password=%s sslmode=disable connect_timeout=5",
		"postgres",
		"pharmacy_db",
		"localhost",
		"123456",
	)

	utils.ConnectDB(connStr)
	defer utils.GetDB().Close(context.Background())

	// Gin router
	r := gin.Default()
	api := r.Group("/api")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Auth
	controllers.AuthRoute(api)

	// Existing routes
	controllers.UserRoute(api)
	controllers.PharmacyRoute(api)
	controllers.MedicinesRoute(api)
	controllers.OrderRoute(api)
	controllers.CategoryRoute(api)

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
