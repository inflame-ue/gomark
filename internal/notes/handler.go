package notes

import (
	"encoding/json"
	"net/http"

	"github.com/inflame-ue/gomark/internal/db"
	"github.com/inflame-ue/gomark/internal/utils"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"text"`
}

func HandlePostNote(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if err := utils.RequireMethod(r, http.MethodPost); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
	}

	if err := utils.RequireContentType(r, "application/json"); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusBadRequest)
	}

	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusBadRequest)
	}

	err := db.CreateNote(note.Title, note.Content)
	if err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusInternalServerError)
	}

}
