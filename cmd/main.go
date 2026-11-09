package main

import (
	"data-intensive-api/cmd/internals/db"
	"data-intensive-api/cmd/internals/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Db *pgxpool.Pool

func main(){

	db.ConnectDB()

	mux := http.NewServeMux()

	routes.SetupRoutes(mux)

	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}