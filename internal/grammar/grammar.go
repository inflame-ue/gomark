package grammar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/inflame-ue/gomark/internal/markdown"
	"github.com/joho/godotenv"
)

type languageToolResponse struct {
	Matches []struct {
		Message      string `json:"message"`
		Replacements []struct {
			Value string `json:"value"`
		} `json:"replacements"`
		Offset  int `json:"offset"`
		Length  int `json:"length"`
		Sentence string `json:"sentence"`
		Type     struct {
			TypeName string `json:"typeName"`
		} `json:"type"`
	} `json:"matches"`
}

func checkGrammar(content *markdown.Markdown) (*languageToolResponse, error) {
	err := godotenv.Load() 
	if err != nil {
		return nil, fmt.Errorf("failed to laod the .env file: %v", err)
	}

	languageToolHost, languageToolAddr := os.Getenv("LANGUAGE_TOOL_HOST"), os.Getenv("LANGUAGE_TOOL_PORT")
	resp, err := http.Post(languageToolHost + ":" + languageToolAddr, "application/text", strings.NewReader(content.Content))
	if err != nil {
		return nil, fmt.Errorf("failed to get a response from language tool API: %v", err)
	}
	defer resp.Body.Close()

	var langToolResp languageToolResponse
	if err := json.NewDecoder(resp.Body).Decode(&langToolResp); err != nil {
		return nil, fmt.Errorf("failed to decode the response body into the languageToolResponse struct: %v", err)
	}

	return &langToolResp, nil
}