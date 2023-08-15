package openai

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"

	"main.go/core"
)

func GenerateImage(prompt string) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("[ERROR] error loading .env file", err)
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
		return
	}
	fmt.Println(respUrl.Data[0].URL)

	reqBase64 := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := c.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("[ERROR] image creation error: %v\n", err)
		return
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("[ERROR] base64 decode error: %v\n", err)
		return
	}

	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		fmt.Printf("[ERROR] PNG decode error: %v\n", err)
		return
	}

	fileName := core.GenerateRandomFileName() + ".png"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("[ERROR] file creation error: %v\n", err)
		return
	}
	defer file.Close()

	if err := png.Encode(file, imgData); err != nil {
		fmt.Printf("[ERROR] PNG encode error: %v\n", err)
		return
	}

	fmt.Printf("[SUCCESS] the image was saved as %s", fileName)
}
