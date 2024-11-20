package database

import (
	"fmt"
	"os"
	"strconv"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
  "github.com/Glicio/go-api-gemini/internal/database/mutations"
)

var Pool *pgx.ConnPool
var Conn *pgx.Conn


func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
  var host = os.Getenv("DB_HOST")
  var port = os.Getenv("DB_PORT")
  var user = os.Getenv("DB_USER")
  var password = os.Getenv("DB_PASSWORD")
  var databaseName = os.Getenv("DB_NAME")

  //pass port to uint16
  portInt, err := strconv.ParseUint(port, 10, 16)
  if err != nil {
    fmt.Println("Error parsing port:", err)
  }


	Pool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
      Host: host,
      Port: uint16(portInt),
      User: user,
      Password: password,
      Database: databaseName,
		},
		MaxConnections: 10,
	})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}
  
  Conn, err = Pool.Acquire()
  if err != nil {
    fmt.Println("Error acquiring connection:", err)
  }

  //create posts table
  _, err = Conn.Exec("CREATE TABLE IF NOT EXISTS posts (id SERIAL PRIMARY KEY, descricao TEXT, src TEXT, alt TEXT)")
  if err != nil {
    fmt.Println("Error creating table:", err)
  }

  //count posts
	rows, err := Conn.Query("SELECT COUNT(*) FROM posts")
	if err != nil {
		fmt.Println("Error counting posts:", err)
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			fmt.Println("Error scanning row:", err)
		}
	}

  if count == 0 {
    //create mock posts
    createPostInput1 := mutations.CreatePostInput{
      Descricao: "Foto de um gato",
      Src: "https://placecats.com/millie/200/200",
      Alt: "Um gato",
    }
    createPostInput2 := mutations.CreatePostInput{
      Descricao: "Foto de um gato 2",
      Src: "https://placekitten.com/200/200",
      Alt: "Um gatinho",
    }
    createPostInput3 := mutations.CreatePostInput{
      Descricao: "Foto de um gato 3",
      Src: "https://placekitten.com/200/200",
      Alt: "Um gatito",
    }

    _, err = mutations.CreatePost(Conn, createPostInput1)
    if err != nil {
      fmt.Println("Error creating post:", err)
    }
    _, err = mutations.CreatePost(Conn, createPostInput2)
    if err != nil {
      fmt.Println("Error creating post:", err)
    }
    _, err = mutations.CreatePost(Conn, createPostInput3)
    if err != nil {
      fmt.Println("Error creating post:", err)
    }

  }

}
