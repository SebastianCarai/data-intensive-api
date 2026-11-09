package handlers

import (
	"fmt"
	"net/http"
)

func AddLikeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Added like")
}


func AddCommentHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Added comment")
}