package main

import (
	"fmt"
	"net/http"

	api "github.com/Glicio/go-api-gemini/api"
	"github.com/Glicio/go-api-gemini/internal/database"
	"github.com/Glicio/go-api-gemini/www"
)

func main() {
	http.HandleFunc("/", www.Index)
  http.HandleFunc("/static/", www.Static)
  http.HandleFunc("/posts", api.Posts)
  http.HandleFunc("/upload", api.Upload)
  database.Init()
	fmt.Println("Listening on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
