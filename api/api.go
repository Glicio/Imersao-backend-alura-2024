package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	types "github.com/Glicio/go-api-gemini/api/types"
	"github.com/Glicio/go-api-gemini/internal/database"
	"github.com/Glicio/go-api-gemini/internal/database/queries"
)

func Api(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
    //blocking shit i gess...
		posts, err := queries.GetPosts(database.Conn)
		if err != nil {
			fmt.Fprintf(w, "Error getting posts: %v", err)
			return
		}
		fmt.Println("sending posts")
		json.NewEncoder(w).Encode(posts)
		return
	}
	//parse the request body
	var req types.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	var key string
	if req.Key != "" {
		key = req.Key
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing key"))
		return
	}

	if key == "secret" {
		fmt.Println("Right key")
		w.Header().Set("Content-Type", "application/json")
    //blocking shit i gess...
		posts, err := queries.GetPosts(database.Conn)
		if err != nil {
			fmt.Fprintf(w, "Error getting posts: %v", err)
			return
		}
		fmt.Println("sending posts")
		json.NewEncoder(w).Encode(posts)
		return

	}

	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	fmt.Fprintf(w, "Hello, world!")
}
