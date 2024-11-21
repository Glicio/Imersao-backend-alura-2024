package www

import (
  "net/http"
  "os"
  "fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
  pwd, err := os.Getwd()
  if err != nil {
    fmt.Println("Error getting pwd:", err)
  }
  var path = pwd + "/www/index.html"
  fmt.Println("Serving file:", path)
  http.ServeFile(w, r, path)
}
