package utils 

import (
	"context"
	"fmt"
	"os"


)

var db *pgx.Conn 

func ConnectDB(config string) {
	conn,  err := pgx.Connect(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "DB error: %v\n", err)
		os.Exit(1)
	}
	db = conn
}

func GetDB() *pggx.Conn {
	return db
}
