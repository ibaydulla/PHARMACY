package main

import (
"context"
"fmt"
"log"
"os"
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
