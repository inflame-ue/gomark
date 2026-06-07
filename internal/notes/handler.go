package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/inflame-ue/gomark/internal/db"
	"github.com/inflame-ue/gomark/internal/request"
	"github.com/inflame-ue/gomark/internal/respond"
	"github.com/inflame-ue/gomark/internal/storage"
)

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"text"`
}

type successMessage struct {
	Message string `json:"message"`
}

func HandleUploadNote(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if err := request.RequireMethod(r, http.MethodPost); err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
		return
	}

	const maxMemory = 10 << 20
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		respond.WriteErrorAndLog(w, fmt.Errorf("failed to parse form data: %v", err), http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	file, _, err := r.FormFile("file")
	if err != nil {
		respond.WriteErrorAndLog(w, fmt.Errorf("failed to open the form file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		respond.WriteErrorAndLog(w, fmt.Errorf("failed to read the file content to memory: %v", err), http.StatusInternalServerError)
		return
	}

	filepath, err := storage.SaveToDisk(data)
	if err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	_, err = db.CreateNoteFromUpload(title, string(data), filepath)
	if err != nil {
		respond.WriteErrorAndLog(w, fmt.Errorf("failed to create note: %v", err), http.StatusInternalServerError)
		return
	}

	successMsg := successMessage{
		Message: fmt.Sprintf("Note with title '%v' saved to disk and db successfully", title),
	}
	err = respond.WriteJSON(w, successMsg)
	if err != nil {
		log.Print(err)
	}
}

func HandlePostNote(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if err := request.RequireMethod(r, http.MethodPost); err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
		return
	}

	if err := request.RequireContentType(r, "application/json"); err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	var note CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	_, err := db.CreateNote(note.Title, note.Content)
	if err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusInternalServerError)
		return
	}

	successMsg := successMessage{
		Message: fmt.Sprintf("Note with title '%v' created successfully", note.Title),
	}
	err = respond.WriteJSON(w, successMsg)
	if err != nil {
		log.Print(err)
	}
}

func HandleGetNotes(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if err := request.RequireMethod(r, http.MethodGet); err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
		return
	}

	notes, err := db.ListNotes()
	if err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	err = respond.WriteJSON(w, notes)
	if err != nil {
		log.Print(err)
	}
}

func HandleGetNote(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if err := request.RequireMethod(r, http.MethodGet); err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
		return
	}

	noteID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		respond.WriteErrorAndLog(w, errors.New("the id provided must be a valid integer"), http.StatusBadRequest)
		return
	}

	note, err := db.GetNoteById(noteID)
	if err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	err = respond.WriteJSON(w, note)
	if err != nil {
		log.Print(err)
	}
}
