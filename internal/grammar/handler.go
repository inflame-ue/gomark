package grammar

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/inflame-ue/gomark/internal/markdown"
	"github.com/inflame-ue/gomark/internal/utils"
)

func HandleGrammarCheck(w http.ResponseWriter, r *http.Request) {
	if err := utils.RequireMethod(r, http.MethodPost); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
		return
	}

	if err := utils.RequireContentType(r, "application/json"); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	var markdown markdown.Markdown
	if err := json.NewDecoder(r.Body).Decode(&markdown); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	grammarIssues, err := checkGrammar(&markdown)
	if err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, grammarIssues)
	if err != nil {
		log.Print(err)
	}
}
