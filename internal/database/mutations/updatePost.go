package mutations

import (
	"fmt"
	"github.com/jackc/pgx"
	types "github.com/Glicio/go-api-gemini/api/types"
)

type UpdatePostInput struct {
  ID string `json:"id"`
	Descricao string `json:"descricao,omitempty"`
	Src       string `json:"src,omitempty"`
	Alt       string `json:"alt,omitempty"`
}

func UpdatePost(conn *pgx.Conn, input UpdatePostInput) (types.Post, error) {

  id := input.ID
	descricao := input.Descricao
	src := input.Src
	alt := input.Alt

	var updatedPost types.Post

	_, err := conn.Exec("UPDATE posts SET descricao = $1, src = $2, alt = $3 WHERE id = $4", descricao, src, alt, id)
	if err != nil {
		return updatedPost, err
	}
  fmt.Println("updated post with id: ", id)

	return updatedPost, nil
}
