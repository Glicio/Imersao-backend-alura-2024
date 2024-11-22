package api

import (
	"encoding/json"
	"fmt"
	"github.com/Glicio/go-api-gemini/internal/database"
	"github.com/Glicio/go-api-gemini/internal/database/mutations"
	"github.com/Glicio/go-api-gemini/services"
	"io"
	"net/http"
	"os"
)

type UploadRequest struct {
	Image     string `json:"image"`
	Descricao string `json:"descricao"`
	Src       string `json:"src"`
	Alt       string `json:"alt"`
}

func Upload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting pwd:", err)
		}
		var path = pwd + "/www/index.html"
		fmt.Println("Serving file:", path)
		http.ServeFile(w, r, path)
		return
	}

	err := r.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		fmt.Println("Error getting multipart reader:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("image")

	if err != nil {
		fmt.Println("Error getting file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var pwd string
	pwd, err = os.Getwd()

	if err != nil {
		fmt.Println("Error getting working directory:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//check if file extension is allowed
	ext := handler.Filename[len(handler.Filename)-4:]
	if ext != ".jpg" && ext != ".png" && ext != ".gif" && ext != "jpeg" { // flawed, but idc for this toy project
		fmt.Println("File extension not allowed:", ext)
		http.Error(w, "File extension not allowed", http.StatusBadRequest)
		return
	}

	//check if folder exists
	if _, err := os.Stat(pwd + "/static"); os.IsNotExist(err) {
		os.Mkdir(pwd+"/static", 0755) // Create a new directory
	}

	var newFilePath = pwd + "/static/" + handler.Filename

	out, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println("Error copying file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	file.Seek(0, 0)
	b, err := io.ReadAll(file)

	altText := services.GenerateAltFromImage(b, ext)
	desc := services.GenerateDescriptionFromImage(b, ext)
	fmt.Println("File uploaded successfully:", handler.Filename)
	var newPost = mutations.CreatePostInput{
		Descricao: desc,
		Src:       newFilePath,
		Alt:       altText,
	}

	post, err := mutations.CreatePost(database.Conn, newPost)

	if err != nil {
		fmt.Println("Error creating post:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Post created successfully:", post.Id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
