package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	types "github.com/Glicio/go-api-gemini/api/types"
	"github.com/Glicio/go-api-gemini/internal/database"
	"github.com/Glicio/go-api-gemini/internal/database/mutations"
	"github.com/Glicio/go-api-gemini/internal/database/queries"
)

/**
 * GET /posts - Get all posts
 * POST /posts - Create a new post
*  EX: curl -X POST -H "Content-Type: application/json" -d '{"key": "secret","descricao": "teste", "src": "https://placekitten.com/200/200", "alt": "teste"}' http://localhost:3000/posts
*/
func Posts(w http.ResponseWriter, r *http.Request) {

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
	var novoPost types.Request
	err := json.NewDecoder(r.Body).Decode(&novoPost)
	var key string
	if novoPost.Key != "" {
		key = novoPost.Key
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing key"))
		return
	}

	if key != "secret" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	fmt.Println("Right key")

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Fprintf(w, "Error getting posts: %v", err)
		return
	}

  var toCreate mutations.CreatePostInput = mutations.CreatePostInput{
    Descricao: novoPost.Descricao,
    Src: novoPost.Src,
    Alt: novoPost.Alt,
  }
	fmt.Println("Creating post")
	createdPost, err := mutations.CreatePost(database.Conn, toCreate)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintln(w, "Erro ao criar o post")
    return
  }
	fmt.Printf("Created post %v", createdPost.Id)
	json.NewEncoder(w).Encode(createdPost)
	return

}
