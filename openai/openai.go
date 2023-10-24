package openai

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func GenerateImage(prompt string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("[ERROR] error loading .env file", err)
        panic(err)
	}
	c := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()

	reqUrl := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}

	respUrl, err := c.CreateImage(ctx, reqUrl)
	if err != nil {
		fmt.Printf("[ERROR] image creation error: %v\n", err)
	}

	return respUrl.Data[0].URL
}
