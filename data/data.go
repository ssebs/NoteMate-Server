package data

import (
	"time"

	"github.com/beevik/guid"
)

// Note
// ID is a GUID
// LastUpdated is a time.Time in UTC
type Note struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	Author      string    `json:"author"`
	LastUpdated time.Time `json:"last_updated"`
	Version     int       `json:"version"`
	Active      bool      `json:"active"`
}

// NewNote will create a new note from the title, author, and contents.
// LastUpdated, Version, and Active will be set by default
// Retuns a *Note
func NewNote(title, author, contents string) *Note {
	return &Note{
		ID:          guid.NewString(),
		Title:       title,
		Contents:    contents,
		Author:      author,
		LastUpdated: time.Now().UTC(),
		Version:     1,
		Active:      true,
	}
}
