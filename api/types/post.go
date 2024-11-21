package api

type Post struct {
  //id is optional
  Id int64 `json:"id,omitempty"`
  Descricao string `json:"descricao"`
  Src string `json:"src"`
  Alt string `json:"alt"`
}

