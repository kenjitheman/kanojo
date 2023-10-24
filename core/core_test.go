package core

import (
	"testing"
)

func TestGenerateRandomFileName(t *testing.T) {
	fileName := GenerateRandomFileName()
	if len(fileName) != fileNameLength {
		t.Errorf("Generated filename length is not as expected")
	}
	for _, char := range fileName {
		if !isLetter(char) {
			t.Errorf("Generated filename contains invalid characters")
		}
	}
}

func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func TestRandomness(t *testing.T) {
	const numIterations = 100
	fileNames := make(map[string]bool)
	for i := 0; i < numIterations; i++ {
		fileName := GenerateRandomFileName()
		if fileNames[fileName] {
			t.Errorf("Generated duplicate filenames")
		}
		fileNames[fileName] = true
	}
}
