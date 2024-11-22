package services

import (
	"bytes"
	ctx "context"
	// import standard libraries
	// Import the GenerativeAI package for Go
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)


func GenerateAltFromImage(image []byte, imageExtension string) (string) {
	err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }
  var apiKey = os.Getenv("GEMINI_API_KEY")

  client, err := genai.NewClient(ctx.Background(), option.WithAPIKey(apiKey)) 
  if err != nil {
    fmt.Println("Error creating client:", err)
  }
  defer client.Close()

  model := client.GenerativeModel("gemini-1.5-flash")

  var prompt = "Gere um alt text em português do Brasil para a seguinte imagem, gere apenas o texto de alt";

  response, err := model.GenerateContent(ctx.Background(), genai.Text(prompt), genai.ImageData(imageExtension,image))

  if err != nil {
    fmt.Println("Error generating content:", err)
  }
  parts := response.Candidates[0].Content.Parts

  fmt.Println("Parts:")
  finalText := new(bytes.Buffer)
  for _, part := range parts {
    fmt.Fprintf(finalText, "%s\n", part)
  }

  return finalText.String()
}


func GenerateDescriptionFromImage(image []byte, imageExtension string) (string) {
	err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }
  var apiKey = os.Getenv("GEMINI_API_KEY")

  client, err := genai.NewClient(ctx.Background(), option.WithAPIKey(apiKey)) 
  if err != nil {
    fmt.Println("Error creating client:", err)
  }
  defer client.Close()

  model := client.GenerativeModel("gemini-1.5-flash")

  var prompt = "Gere uma descrição em português do Brasil para a seguinte imagem, bem detalhada, gere apenas o texto de descrição";

  response, err := model.GenerateContent(ctx.Background(), genai.Text(prompt), genai.ImageData(imageExtension,image))

  if err != nil {
    fmt.Println("Error generating content:", err)
  }
  parts := response.Candidates[0].Content.Parts

  fmt.Println("Parts:")
  finalText := new(bytes.Buffer)
  for _, part := range parts {
    fmt.Fprintf(finalText, "%s\n", part)
  }

  return finalText.String()
}
