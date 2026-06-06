package db

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
