package routes

import (
	"data-intensive-api/handlers"
	"net/http"
)



func SetupRoutes(mux *http.ServeMux){
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/users", handlers.UserHandler)
	mux.HandleFunc("/feed", handlers.FeedHandler)
}