package api

//post creation request
type Request struct {
  Key string `json:"key,omitempty"`
  Descricao string `json:"descricao,omitempty"`
  Src string `json:"src,omitempty"`
	Alt string `json:"alt,omitempty"`
}
