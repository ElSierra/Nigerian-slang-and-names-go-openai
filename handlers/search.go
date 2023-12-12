package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/elsierra/go-echo-zik/lib"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

func PostHandler(c echo.Context) error {

	var reqBody io.ReadCloser = c.Request().Body

	defer reqBody.Close()

	body, _ := io.ReadAll(reqBody)
	var myData msg
	err := json.Unmarshal(body, &myData)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	C, err := lib.OpenAiClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, msg{Message: "Error"})
	}
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Just give the response only",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "what does the name " + myData.Message + "mean in Nigeria",
			},
		},
		Stream: true,
	}
	stream, err := C.CreateChatCompletionStream(c.Request().Context(), req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return err
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return err
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return err
		}

		c.JSON(http.StatusOK, msg{Message: response.Choices[0].Delta.Content})
	}
}
