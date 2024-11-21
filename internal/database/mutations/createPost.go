package mutations

import (
	types "github.com/Glicio/go-api-gemini/api/types"
	"github.com/jackc/pgx"
)

type CreatePostInput struct {
	Descricao string `json:"descricao"`
	Src       string `json:"src"`
	Alt       string `json:"alt"`
}

func CreatePost(conn *pgx.Conn, input CreatePostInput) (types.Post, error) {
	descricao := input.Descricao
	src := input.Src
	alt := input.Alt

	var newPost types.Post
	err := conn.QueryRow("INSERT INTO posts (descricao, src, alt) VALUES ($1, $2, $3) RETURNING *", descricao, src, alt).Scan(&newPost.Id, &newPost.Descricao, &newPost.Src, &newPost.Alt)
	if err != nil {
		return newPost, err
	}
	return newPost, nil
}
