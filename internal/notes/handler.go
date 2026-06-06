package notes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/inflame-ue/gomark/internal/db"
	"github.com/inflame-ue/gomark/internal/utils"
)

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"text"`
}

type successMessage struct {
	Message string `json:"message"`
}

func HandleUploadNote(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if err := utils.RequireMethod(r, http.MethodPost); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
		return
	}

	const maxMemory = 10 << 20
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		utils.WriteErrorAndLog(w, fmt.Errorf("failed to parse form data: %v", err), http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.WriteErrorAndLog(w, fmt.Errorf("failed to open the form file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		utils.WriteErrorAndLog(w, fmt.Errorf("failed to read the file content to memory: %v", err), http.StatusInternalServerError)
		return
	}

	filepath, err := utils.SaveToDisk(data)
	if err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	_, err = db.CreateNoteFromUpload(title, string(data), filepath)
	if err != nil {
		utils.WriteErrorAndLog(w, fmt.Errorf("failed to create note: %v", err), http.StatusInternalServerError)
		return
	}

	successMsg := successMessage{
		Message: fmt.Sprintf("Note with title '%v' saved to disk and db successfully", title),
	}
	err = utils.WriteJSON(w, successMsg)
	if err != nil {
		log.Print(err)
	}
}

func HandlePostNote(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if err := utils.RequireMethod(r, http.MethodPost); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
		return
	}

	if err := utils.RequireContentType(r, "application/json"); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	var note CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	_, err := db.CreateNote(note.Title, note.Content)
	if err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	successMsg := successMessage{
		Message: fmt.Sprintf("Note with title '%v' created successfully", note.Title),
	}
	err = utils.WriteJSON(w, successMsg)
	if err != nil {
		log.Print(err)
	}
}

func HandleGetNotes(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if err := utils.RequireMethod(r, http.MethodGet); err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
		return
	}

	notes, err := db.ListNotes()
	if err != nil {
		utils.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	err = utils.WriteJSON(w, notes)
	if err != nil {
		log.Print(err)
	}
}
