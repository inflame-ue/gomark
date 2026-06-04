package grammar

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/inflame-ue/gomark/internal/markdown"
	"github.com/inflame-ue/gomark/internal/utils"
)

func HandleGrammarCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodErr := fmt.Errorf("expected the HTTP method to be POST, got %v instead", r.Method)

		log.Print(methodErr.Error())
		err := utils.WriteError(w, methodErr, http.StatusBadRequest)
		if err != nil {
			log.Print(err)
		}

		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		headerErr := fmt.Errorf("expected the markdown file to be passed as application/json, got %v instead", r.Header.Get("Content-Type"))

		log.Print(headerErr.Error())
		err := utils.WriteError(w, headerErr, http.StatusBadRequest)
		if err != nil {
			log.Print(err)
		}

		return
	}

	var markdown markdown.Markdown
	if err := json.NewDecoder(r.Body).Decode(&markdown); err != nil {
		marshalErr := errors.New("failed to unmarshal the json into the Makrdown struct, the syntax is probably invalid")

		log.Print(marshalErr.Error())
		err := utils.WriteError(w, marshalErr, http.StatusBadRequest)
		if err != nil {
			log.Print(err)
		}

		return
	}

	err := utils.WriteJSON(w, markdown)
	if err != nil {
		log.Print(err)
	}
}
