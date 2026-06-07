package mark

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/inflame-ue/gomark/internal/db"
	"github.com/inflame-ue/gomark/internal/request"
	"github.com/inflame-ue/gomark/internal/respond"
)

func HandleRenderNoteHTML(w http.ResponseWriter, r *http.Request, db *db.DB) {
	if err := request.RequireMethod(r, http.MethodGet); err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusMethodNotAllowed)
		return
	}

	noteID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		respond.WriteErrorAndLog(w, fmt.Errorf("conversion of note id to integer failed: %v", err), http.StatusBadRequest)
		return
	}

	note, err := db.GetNoteById(noteID)
	if err != nil {
		respond.WriteErrorAndLog(w, err, http.StatusBadRequest)
		return
	}

	html := markdownToHTML([]byte(note.Content))
	page := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head><title>%s</title></head>
<body>%s</body>
</html>`, note.Title, html)

	err = respond.WriteHTML(w, []byte(page))
	if err != nil {
		log.Print(err)
	}
}
