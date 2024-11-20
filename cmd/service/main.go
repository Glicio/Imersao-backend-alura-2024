package main

import (
	"fmt"
	"net/http"

	api "github.com/Glicio/go-api-gemini/api"
	"github.com/Glicio/go-api-gemini/internal/database"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
type Request struct {
	Message string `json:"message"`
  //optional key
  Key string `json:"key,omitempty"`
}

func main() {
	http.HandleFunc("/", api.Index)
  http.HandleFunc("/api", api.Api)
  database.Init()
	fmt.Println("Listening on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
