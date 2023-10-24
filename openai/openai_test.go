package openai

import (
	"testing"
	"time"
)

func TestGenerateImage(t *testing.T) {
	prompt := "Generate an image of a cat."
	startTime := time.Now()
	imageURL := GenerateImage(prompt)
	elapsed := time.Since(startTime)
	if imageURL == "" {
		t.Errorf("Generated image URL is empty")
	}
	if elapsed > 10*time.Second {
		t.Errorf("Response time is greater than 10 seconds: %v", elapsed)
	}
}

func TestGenerateImageWithDifferentPrompt(t *testing.T) {
	prompt := "Generate an image of a dog."
	startTime := time.Now()
	imageURL := GenerateImage(prompt)
	elapsed := time.Since(startTime)
	if imageURL == "" {
		t.Errorf("Generated image URL is empty")
	}
	if elapsed > 10*time.Second {
		t.Errorf("Response time is greater than 10 seconds: %v", elapsed)
	}
}
