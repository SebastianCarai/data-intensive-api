package handlers

import (
	"data-intensive-api/db"
	"data-intensive-api/structs"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func FeedHandler(w http.ResponseWriter, r *http.Request){
	max := 10000
	min := 1
	randInt := rand.IntN(max - min) + min
	
	query, err := db.GetFeedQuery(randInt)

	if err != nil {
		return
	}

	rows, queryErr := db.Pool.Query(r.Context(), query)

	if queryErr != nil{
		fmt.Printf("Error while querying db: %v\n", queryErr)
		http.Error(w, "internal server error", 500)
		return
	}

	feed,feedErr := pgx.CollectRows(rows, pgx.RowToStructByName[structs.Post])
	
	if feedErr != nil{
		fmt.Printf("Error while converting sql rows to json: %v\n", feedErr)
		http.Error(w, "internal server error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}