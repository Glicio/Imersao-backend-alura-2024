package main

import (
	"fmt"
	"net/http"

	api "github.com/Glicio/go-api-gemini/api"
	"github.com/Glicio/go-api-gemini/internal/database"
)

func main() {
	http.HandleFunc("/", api.Index)
  http.HandleFunc("/posts", api.Posts)
  database.Init()
	fmt.Println("Listening on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
