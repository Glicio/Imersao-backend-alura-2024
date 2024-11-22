package www 


import (
  "net/http"
  "os"
  "fmt"
)



func Static(w http.ResponseWriter, r *http.Request) {
  pwd, err := os.Getwd()
  if err != nil {
    fmt.Println("Error getting pwd:", err)
  }

  var file = r.URL.Path 
  var path = pwd + file

  fmt.Println("Serving static files from", path)
  http.ServeFile(w, r, path)
}

