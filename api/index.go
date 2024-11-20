package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	types "github.com/Glicio/go-api-gemini/api/types"
)

func Index(w http.ResponseWriter, r *http.Request) {

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

  if(key == "secret"){
    var res types.Response = types.Response{Message: "Authorized", Status: "OK"}
    json.NewEncoder(w).Encode(res)
    return
  }

	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	fmt.Fprintf(w, "Hello, world!")
}

