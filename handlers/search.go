package handlers

import (
	"fmt"

	"net/http"
	"strings"

	"github.com/elsierra/go-echo-zik/lib"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

// msg struct should be defined to match the expected JSON structure.

// PostHandler handles POST requests and interacts with the OpenAI API.
func PostHandler(c echo.Context) error {
	var myData msg
	// Bind the request body to the myData struct.
	if err := c.Bind(&myData); err != nil {
		return c.JSON(http.StatusBadRequest, msg{Message: "Invalid request body"})
	}

	// Initialize the OpenAI client using a custom function (assumed to be correct).
	client, err := lib.OpenAiClient()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msg{Message: "Error creating OpenAI client"})
	}

	// Define the chat completion request.
	req := openai.ChatCompletionRequest{
		Model:     "gpt-4", // Use the correct model identifier.
		MaxTokens: 100,
		Messages: []openai.ChatCompletionMessage{
			// {
			// 	Role:    openai.ChatMessageRoleUser,
			// 	Content: "You're a dictionary of nigerian names and slangs, response should be a definition of the name or slang, in the this form 'name/slang' means 'definition' or 'name/slang' is 'definition' and include the origin of the name or slang in this format 'origin: origin of name/slang' or 'origin: origin of name/slang' and 'name/slang' means 'definition' or 'name/slang' is 'definition' and include the origin of the name or slang in this format 'origin: origin of name/slang' maximum of 60 characters",
			// },
	
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "What does the name or slang iyanu mean? in nigeria",
			},
			{Role: openai.ChatMessageRoleAssistant,
				Content: "{name:'Iyanu', origin:'Yorùbá', type:'name', meaning:'God's miracle', full:'Iyanuoluwa',sentence: 'Iyanu likes playing football' etymology:'Iyanu means miracle and oluwa means God' }}"},
			{
				Role:    openai.ChatMessageRoleUser,
				Content:  fmt.Sprintf("what about %s, use the above format in a stringified json, if you don't know just return null", myData.Message) ,
			},
		},
	}

	// Create the chat completion.
	response, err := client.CreateChatCompletion(c.Request().Context(),req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msg{Message: fmt.Sprintf("ChatCompletion error: %v", err)})
	}

	// Process the response.
	responseString := response.Choices[0].Message.Content
	fmt.Println(responseString)

	// Return the response as JSON.
	return c.JSON(http.StatusOK, msg{Message: strings.Trim(responseString, "\n")})
}
