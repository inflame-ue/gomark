package grammar

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/inflame-ue/gomark/internal/markdown"
	"github.com/inflame-ue/gomark/internal/request"
	"github.com/inflame-ue/gomark/internal/respond"
)

func HandleGrammarCheck(w http.ResponseWriter, r *http.Request) {
	if err := request.RequireMethod(r, http.MethodPost); err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
		return
	}

	if err := request.RequireContentType(r, "application/json"); err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	var markdown markdown.Markdown
	if err := json.NewDecoder(r.Body).Decode(&markdown); err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	grammarIssues, err := checkGrammar(&markdown)
	if err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	err = respond.WriteJSON(w, grammarIssues)
	if err != nil {
		log.Print(err)
	}
}
