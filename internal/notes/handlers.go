package notes

import "net/http"

func HandlePostNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This endpoint will sooner or later handle uploading notes via JSON."))
}
