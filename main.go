package main

import (
	"data-intensive-api/db"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

var Db *sql.DB


func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from air"))
}

func usersHandler(w http.ResponseWriter, r *http.Request){

	rows, err := Db.Query("select * from users limit 10;")
	if err != nil{
		fmt.Printf("Error while querying db: %v\n", err)
	}

	data,convertErr := db.ConvertRowsToJson(rows)
	if convertErr != nil{
		fmt.Printf("Error while converting sql rows to json: %v\n", convertErr)
	}


	w.Write([]byte(data))
}


func main(){
	Db = db.ConnectDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/users", usersHandler)


	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}