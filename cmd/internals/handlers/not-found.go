package handlers

import "net/http"

func PageNotFoundHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("Page not found. <a href='/'>Go back to homepage<a>"))
}