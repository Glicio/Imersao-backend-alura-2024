package api

import (
	"encoding/json"
	"fmt"
	types "github.com/Glicio/go-api-gemini/api/types"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("GET not allowed"))
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
    fmt.Println("secret")
		var res types.Response = types.Response{Message: "A Torre Eiffel ilumina à noite, com milhares de luzes cintilando, criando um espetáculo mágico em Paris.", Status: "OK"}
    w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
		return
	}

	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	fmt.Fprintf(w, "Hello, world!")
}
