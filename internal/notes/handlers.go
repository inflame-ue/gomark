package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/inflame-ue/gomark/internal/db"
	"github.com/inflame-ue/gomark/internal/utils"
)

type Note struct {
	Title   string `json:"title"`
	Content string `json:"text"`
}

func HandlePostNote(w http.ResponseWriter, r *http.Request, db *db.DB) {
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
		headerErr := fmt.Errorf("expected the title and content informat to be passed as application/json, got %v instead",
			r.Header.Get("Content-Type"))
		log.Print(headerErr.Error())
		err := utils.WriteError(w, headerErr, http.StatusBadRequest)
		if err != nil {
			log.Print(err)
		}
		return
	}

	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		unmarshalErr := errors.New("failed to unmarshal the json into the Note struct, the syntax is probably invalid")
		log.Print(unmarshalErr.Error())
		err := utils.WriteError(w, unmarshalErr, http.StatusBadRequest)
		if err != nil {
			log.Print(err)
		}
		return
	}

	err := db.CreateNote(note.Title, note.Content)
	if err != nil {
		log.Print(err.Error())
		err := utils.WriteError(w, err, http.StatusInternalServerError)
		if err != nil {
			log.Print(err)
		}
		return
	}

	success := struct{
		message string
	}{
		message: "created the note with sucess",
	}
	err = utils.WriteJSON(w, success)
	if err != nil {
		log.Print(err)
	}
}
