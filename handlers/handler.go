// handlers/handler.go
package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/elsierra/go-echo-zik/internal/database"
	"github.com/labstack/echo/v4"
)

type msg struct {
	Message string `json:"message"`
}
type ApiConfig struct {
	DB *database.Queries
}
type Dictionary struct {
	ID         int32 `json:"id"`
	Word       string `json:"word"`
	Origin     string `json:"origin"`
	Fullword   string `json:"fullword"`
	Definition string `json:"definition"`
	Etymology  string`json:"etymology"`
	Type       string `json:"type"`
	Sentence   string `json:"sentence"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}




func  ToDictionary( dictionary *database.Dictionary) Dictionary {

	toString := func(origin sql.NullString) string {
		if origin.Valid {
			return origin.String
		}
		return ""
	}
	
	return Dictionary{
		ID:         dictionary.ID,
		Word:       dictionary.Word,
		Origin:     toString(dictionary.Origin),
		Fullword:  toString(dictionary.Fullword),
		Definition: toString(dictionary.Definition),
		Etymology: toString(dictionary.Etymology),
		Type:       toString(dictionary.Type),
		Sentence:   toString(dictionary.Sentence),
		CreatedAt:  dictionary.CreatedAt,
		UpdatedAt:  dictionary.UpdatedAt,
	}
}

func (apiConfig *ApiConfig) HomeHandler(c echo.Context) error {
	
dictionary, err := apiConfig.DB.CreateWord(c.Request().Context(),database.CreateWordParams{
	Word: "Iyanud",
	Origin: sql.NullString{
		String: "Yorùbá",
		Valid: true,
	},
	Fullword: sql.NullString{
		String: "Iyanuoluwa",
		Valid: true,
	},
	Definition: sql.NullString{
		String: "God's miracle",
		Valid: true,
	},
	Etymology: sql.NullString{
		String: "Iyanu means miracle and oluwa means God",
		Valid: true,
	},
	Type: sql.NullString{
		String: "name",
		Valid: true,
	},
})

if err != nil {
	return c.JSON(http.StatusInternalServerError, msg{Message:err.Error()})
}
return c.JSON(http.StatusOK, ToDictionary(&dictionary))
}
