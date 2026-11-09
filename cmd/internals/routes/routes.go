package routes

import (
	"data-intensive-api/cmd/internals/handlers"
	"net/http"
)



func SetupRoutes(mux *http.ServeMux){
	// 404
	mux.HandleFunc("GET /", handlers.PageNotFoundHandler)


	// Homepage
	mux.HandleFunc("GET /{$}", handlers.HomeHandler)


	// Users
	mux.HandleFunc("GET /users", handlers.UserHandler)


	// Posts
	mux.HandleFunc("GET /feed", handlers.FeedHandler)
	mux.HandleFunc("GET /posts", handlers.PostDetailsHandler)

	mux.HandleFunc("POST /posts", handlers.NewPostHandler)


	// Likes and Comments
	mux.HandleFunc("POST /posts/likes", handlers.AddLikeHandler)
	mux.HandleFunc("POST /posts/comments", handlers.AddCommentHandler)
}