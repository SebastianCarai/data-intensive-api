package handlers

import (
	"data-intensive-api/cmd/internals/db"
	"data-intensive-api/cmd/internals/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)

func AddLikeHandler(w http.ResponseWriter, r *http.Request){
	authorId := utils.GetRandomInt(1, 10000)
	postId := utils.GetRandomInt(1, 500000)

	fmt.Println(postId)

	var likes_count int
	db.Pool.QueryRow(
		r.Context(), 
		`Select likes_count from posts where id = $1`, postId,
	).Scan(&likes_count)
	fmt.Printf("Before: %v\n", likes_count)
	
	like, err := db.Pool.Exec(
		r.Context(),
		`with ins as (
			insert into likes (user_id, post_id, created_at)
			values ($1, $2, now())
			on conflict do nothing
			returning post_id
		)
		update posts
		set likes_count = likes_count + 1
		where id in (select post_id from ins);`,
		authorId, postId,
	)

	if err != nil{
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23503"{
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if like.RowsAffected() == 0{
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("No rows affected"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Ok"))

}


func AddCommentHandler(w http.ResponseWriter, r *http.Request){
	authorId := utils.GetRandomInt(1, 10000)
	postId := utils.GetRandomInt(1, 500000)
	body := "New comment"

	fmt.Println(postId)

	var comments_count int
	db.Pool.QueryRow(
		r.Context(), 
		`Select comments_count from posts where id = $1`, postId,
	).Scan(&comments_count)
	fmt.Printf("Before: %v\n", comments_count)
	
	like, err := db.Pool.Exec(
		r.Context(),
		`with ins as (
			insert into comments (user_id, post_id, body, created_at)
			values ($1, $2, $3, now())
			on conflict do nothing
			returning post_id
		)
		update posts
		set comments_count = comments_count + 1
		where id in (select post_id from ins);`,
		authorId, postId, body,
	)

	if err != nil{
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23503"{
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}

		fmt.Printf("%v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if like.RowsAffected() == 0{
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("No rows affected"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Ok"))
}