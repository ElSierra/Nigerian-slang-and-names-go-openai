package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/elsierra/go-echo-zik/lib"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

// Define the msg struct if it's not already defined elsewhere.

func PostHandler(c echo.Context) error {
	var myData msg
	// Use Bind method to read the request body and bind it to myData struct.
	if err := c.Bind(&myData); err != nil {
		return c.JSON(http.StatusBadRequest, msg{Message: "Invalid request body"})
	}

	// Initialize the OpenAI client.
	C, err := lib.OpenAiClient()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msg{Message: "Error creating OpenAI client"})
	}

	// Define the chat completion request.
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4,
		MaxTokens: 60,
		Messages: []openai.ChatCompletionMessage{
			// {
			// 	Role:    openai.ChatMessageRoleUser,
			// 	Content: "You're a dictionary of nigerian names and slangs, response should be a definition of the name or slang, in the this form 'name/slang' means 'definition' or 'name/slang' is 'definition' and include the origin of the name or slang in this format 'origin: origin of name/slang' or 'origin: origin of name/slang' and 'name/slang' means 'definition' or 'name/slang' is 'definition' and include the origin of the name or slang in this format 'origin: origin of name/slang' maximum of 60 characters",
			// },

			{
				Role:    openai.ChatMessageRoleUser,
				Content: "What does the name or slang iyanuoluwa mean? in nigeria",
			},
			{Role: openai.ChatMessageRoleAssistant,
				Content: "{name:'Iyanuoluwa', origin:'Yorùbá', type:'name', meaning:'God's miracle'}}"},
			{
				Role:    openai.ChatMessageRoleUser,
				Content:  fmt.Sprintf("what about %s, use the above format in a stringified json, if you don't know just return null", myData.Message) ,
			},
		},
		Stream: false,
	}

	// Create the chat completion stream.
	stream, err := C.CreateChatCompletionStream(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msg{Message: fmt.Sprintf("ChatCompletionStream error: %v", err)})
	}
	defer stream.Close()

	// Process the stream response.
	var responseString string
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break // End of stream, break the loop.
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, msg{Message: fmt.Sprintf("Stream error: %v", err)})
		}
		responseString += response.Choices[0].Delta.Content
		fmt.Println(responseString)
	}

	// Return the response as JSON.
	return c.JSON(http.StatusOK, msg{Message: strings.ReplaceAll(responseString, "\n", "")})
}
