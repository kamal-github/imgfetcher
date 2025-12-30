package internal

import (
	"fmt"
	"os"
)

func GetPixabayAPIKey() (string, error) {
	key := os.Getenv("PIXABAY_API_KEY")
	if key == "" {
		return "", fmt.Errorf("PIXABAY_API_KEY environment variable not set")
	}
	return key, nil
}
