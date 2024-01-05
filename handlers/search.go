package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"unicode"

	"net/http"
	"strings"

	"github.com/elsierra/go-echo-zik/internal/database"
	"github.com/elsierra/go-echo-zik/lib"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(strings.ToLower(s))
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func (apiCfg ApiConfig) PostHandler(c echo.Context) error {
	var myData msg
	// Bind the request body to the myData struct.
	responseFormat := `{"word":"%s","origin":"%s","type":"%s","definition":"%s","fullWord":"%s","sentence":"%s","etymology":"%s"}`

	if err := c.Bind(&myData); err != nil {
		return c.JSON(http.StatusBadRequest, msg{Message: "Invalid request body"})
	}

	// Initialize the OpenAI client using a custom function (assumed to be correct).
	client, err := lib.OpenAiClient()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msg{Message: "Error creating OpenAI client"})
	}
	myData.Message = CapitalizeFirst(myData.Message)

	wordExists, errWord := apiCfg.DB.GetOneWord(c.Request().Context(), myData.Message)

	if errWord == nil {
		return c.JSON(http.StatusOK, ToDictionary(&wordExists))
	}
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
				Content: fmt.Sprintf(responseFormat,
					"Iyanu",                        // word
					"Yorùbá",                       // origin
					"name",                         // type
					"God's miracle",                // definition
					"Iyanuoluwa",                   // fullWord
					"Iyanu likes playing football", // sentence
					"Iyanu means miracle and oluwa means God", // etymology
				)},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("what about %s, use the above format in a stringified json, if you don't know just return null, always return stringified json, take your time", myData.Message),
			},
		},
	}

	// Create the chat completion.
	response, err := client.CreateChatCompletion(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msg{Message: fmt.Sprintf("ChatCompletion error: %v", err)})
	}

	// Process the response.
	responseString := response.Choices[0].Message.Content
	fmt.Println(responseString)

	var wordDef Dictionary
	err = json.Unmarshal([]byte(responseString), &wordDef)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, msg{Message: fmt.Sprintf("Error unmarshalling response: %v", err)})
	}

	go apiCfg.DB.CreateWord(c.Request().Context(), database.CreateWordParams{
		Word: myData.Message,
		Origin: sql.NullString{
			String: wordDef.Origin,
			Valid:  true,
		},
		Fullword: sql.NullString{
			String: wordDef.Fullword,
			Valid:  true,
		},
		Definition: sql.NullString{
			String: wordDef.Definition,
			Valid:  true,
		},
		Etymology: sql.NullString{
			String: wordDef.Etymology,
			Valid:  true,
		},
		Type: sql.NullString{
			String: wordDef.Type,
			Valid:  true,
		},
		Sentence: sql.NullString{
			String: wordDef.Sentence,
			Valid:  true,
		},
	})

	// Return the response as JSON.
	return c.JSON(http.StatusOK, wordDef)
}
