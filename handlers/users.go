package handlers

import (
	"data-intensive-api/db"
	"data-intensive-api/structs"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)


func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from air"))
}

func UserHandler(w http.ResponseWriter, r *http.Request){

	max := 10000
	min := 1
	randInt := rand.IntN(max - min) + min
	now := time.Now()

	
	rows, err := db.Pool.Query(r.Context(), "select id, username from users where id=$1;", randInt)

	if err != nil{
		fmt.Printf("Error while querying db: %v\n", err)
		http.Error(w, "internal server error", 500)
		return
	}

	users,usersErr := pgx.CollectRows(rows, pgx.RowToStructByName[structs.User])
	if usersErr != nil{
		fmt.Printf("Error while converting sql rows to json: %v\n", usersErr)
		http.Error(w, "internal server error", 500)
		return
	}

	fmt.Println(time.Since(now))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}