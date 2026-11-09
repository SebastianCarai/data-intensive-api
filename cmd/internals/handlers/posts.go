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

func FeedHandler(w http.ResponseWriter, r *http.Request){
	now := time.Now()
	

	query, err := db.GetFeedQuery(utils.GetRandomInt(1, 10000))

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

	fmt.Printf("GET /feed: %vms\n", time.Since(now).Milliseconds())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feed)
}


func PostDetailsHandler(w http.ResponseWriter, r *http.Request){
	postId := utils.GetRandomInt(1, 500000)

	query, err := db.GetPostDetailsQuery(postId)

	if err != nil {
		return
	}

	rows, queryErr := db.Pool.Query(r.Context(), query)

	if queryErr != nil{
		fmt.Printf("Error while querying db: %v\n", queryErr)
		http.Error(w, "internal server error", 500)
		return
	}

	post, postErr := pgx.CollectOneRow(rows, pgx.RowToStructByName[structs.Post])

	if postErr != nil{
		fmt.Printf("Error while converting post details row to json: %v\n", postErr)
		http.Error(w, "internal server error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}


func NewPostHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("New post added")
}