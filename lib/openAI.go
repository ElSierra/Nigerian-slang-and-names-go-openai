package lib

import (
	"errors"
	"os"

	"github.com/sashabaranov/go-openai"
)

func OpenAiClient() (*openai.Client, error) {
	API_KEY := os.Getenv("OPENAI_API_KEY")
	if API_KEY == "" {
		return nil, errors.New("OPENAI_API_KEY is not set")

	}
	C := openai.NewClient(API_KEY)
	return C, nil
}
