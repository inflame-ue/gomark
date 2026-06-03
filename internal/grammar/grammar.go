package grammar

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/inflame-ue/gomark/internal/markdown"
)


func HandleGrammarCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("expected the HTTP method to be POST, got %v instead", r.Method)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Printf("expected the markdown file to be passed as application/json, got %v instead", r.Header.Get("Content-Type"))
		return
	}
	
	var markdown markdown.Markdown
	if err := json.NewDecoder(r.Body).Decode(&markdown); err != nil {
		log.Printf("failed to unmarshal the json into the Makrdown struct: %v", err)
		return
	}

	log.Printf("received the following json markdown: %v", markdown)
}
