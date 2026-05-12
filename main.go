package main

import (
	"context"
	"fmt"
	"log"
	"os"

	
)

func main() {
	//DB connection 

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable connect_timeout=5", "postgres" , "localhost", "5432")
	utils.ConnectDB(connStr)
	defer utils.GetDB().Close(context.Background())

	// HTTP serve
	r := gin.Default()

	rg := r.group("/api")
	controllers.UserRoutes(rg)
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(":=8080"); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	} 
}