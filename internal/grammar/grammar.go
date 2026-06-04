package grammar

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/inflame-ue/gomark/internal/markdown"
)

type languageToolResponse struct {
	Matches []struct {
		Message      string `json:"message"`
		Replacements []struct {
			Value string `json:"value"`
		} `json:"replacements"`
		Offset   int    `json:"offset"`
		Length   int    `json:"length"`
		Sentence string `json:"sentence"`
		Type     struct {
			TypeName string `json:"typeName"`
		} `json:"type"`
	} `json:"matches"`
}

func checkGrammar(content *markdown.Markdown) (*languageToolResponse, error) {
	languageToolHost, languageToolAddr := os.Getenv("LANGUAGE_TOOL_HOST"), os.Getenv("LANGUAGE_TOOL_PORT")
	parsedURL := fmt.Sprintf("http://%s:%s/v2/check", languageToolHost, languageToolAddr)
	data := url.Values{"text": {content.Content}, "language": {"en-US"}}

	resp, err := http.PostForm(parsedURL, data)
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
