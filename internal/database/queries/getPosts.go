package queries

import (
	"github.com/jackc/pgx"
)

type Post struct {
	Id int64 `json:"id"`
	Descricao string `json:"descricao"`
	Src string `json:"src"`
	Alt string `json:"alt"`
}

func GetPosts(conn *pgx.Conn) ([]Post, error) {
	rows, err := conn.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

  var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Descricao, &post.Src, &post.Alt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
