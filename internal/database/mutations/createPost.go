package mutations

import (
	"github.com/jackc/pgx"
)

type CreatePostInput struct {
	Descricao string `json:"descricao"`
	Src       string `json:"src"`
	Alt       string `json:"alt"`
}

func CreatePost(conn *pgx.Conn, input CreatePostInput) (int64, error) {
	descricao := input.Descricao
	src := input.Src
	alt := input.Alt

	var id int64
	err := conn.QueryRow("INSERT INTO posts (descricao, src, alt) VALUES ($1, $2, $3) RETURNING id", descricao, src, alt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
