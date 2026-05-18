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
	//DB connection

	connStr := fmt.Sprintf("user=%s dbname=%s host=%s password=%s  sslmode=disable connect_timeout=5", "postgres", "pharmacy_db", "localhost", "5432")
	utils.ConnectDB(connStr)
	defer utils.GetDB().Close(context.Background())

	// HTTP serve
	r := gin.Default()

	rg := r.Group("/api")
	controllers.UserRoute(rg)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":8080"); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
