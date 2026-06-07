package db

import (
	"database/sql"
	"fmt"
)

type MetaNoteResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DetailedNoteResponse struct {
	ID        int64          `json:"id"`
	Title     string         `json:"title"`
	Filepath  string `json:"filepath"`
	Content   string         `json:"text"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

type NotesResponse struct {
	Notes []MetaNoteResponse `json:"notes"`
}

func (db *DB) CreateNote(title, content string) (int64, error) {
	result, err := db.Connection.Exec("INSERT INTO notes (title, content) VALUES (?, ?)", title, content)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (db *DB) CreateNoteFromUpload(title, content, filepath string) (int64, error) {
	result, err := db.Connection.Exec("INSERT INTO notes (title, content, file_path) VALUES (?, ?, ?)", title, content, filepath)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (db *DB) ListNotes() (*NotesResponse, error) {
	var notes NotesResponse
	notes.Notes = []MetaNoteResponse{} // initiliaze to get [], instead of null for empty

	rows, err := db.Connection.Query("SELECT id, title, created_at, updated_at FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note MetaNoteResponse
		if err := rows.Scan(&note.ID, &note.Title, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
		notes.Notes = append(notes.Notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &notes, nil
}

func (db *DB) GetNoteById(id int) (*DetailedNoteResponse, error) {
	var note DetailedNoteResponse

	row := db.Connection.QueryRow("SELECT id, title, content, COALESCE(file_path, ''), created_at, updated_at FROM notes WHERE id = ?", id)
	if err := row.Scan(&note.ID, &note.Title, &note.Content, &note.Filepath, &note.CreatedAt, &note.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return &note, fmt.Errorf("no note with id: %d", id)
		}
		return &note, fmt.Errorf("scan of the row did not succeed: %v", err)
	}

	return &note, nil
}
