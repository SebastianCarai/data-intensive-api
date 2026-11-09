package handlers

import (
	"data-intensive-api/cmd/internals/db"
	structs "data-intensive-api/cmd/internals/types"
	"data-intensive-api/cmd/internals/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)


func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Homapage"))
}

func UserHandler(w http.ResponseWriter, r *http.Request){

	now := time.Now()

	
	rows, err := db.Pool.Query(r.Context(), "select id, username from users where id=$1;", utils.GetRandomInt(1, 10000))

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

	fmt.Printf("GET /users: %vms\n", time.Since(now).Milliseconds())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}