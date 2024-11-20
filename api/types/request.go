package api

type Request struct {
	Message string `json:"message"`
  //optional key
  Key string `json:"key,omitempty"`
}
