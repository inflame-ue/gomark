# gomark

A CRUD RESTful Markdown Note-taking App built for the [roadmap.sh Markdown Note-taking project](https://roadmap.sh/projects/markdown-note-taking-app).

## Features

- **Grammar check** — `POST /grammar` — checks English grammar via LanguageTool
- **Create note** — `POST /notes` — accepts markdown as JSON
- **Upload note** — `POST /notes/upload` — accepts `.md` file uploads
- **List notes** — `GET /notes` — returns note metadata
- **Get note** — `GET /notes/{id}` — returns full note details
- **Render HTML** — `GET /notes/{id}/html` — returns note rendered as HTML

## Stack

- Go 1.26 
- SQLite via `modernc.org/sqlite` 
- [LanguageTool](https://languagetool.org/) via Docker for grammar checking
- [gomarkdown/markdown](https://github.com/gomarkdown/markdown) for markdown rendering

## Getting Started

```bash
# Start LanguageTool
docker compose up -d

# Run database migrations
goose -dir migrations sqlite db/notes.db up

# Start the server
go run ./cmd/server/
```

## Environment

See `.example_env` for configuration (port, database path, LanguageTool host/port).
