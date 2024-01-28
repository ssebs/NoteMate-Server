// note.go
package data

import (
	"encoding/json"
	"time"

	"github.com/beevik/guid"
)

// NoteBind used for form data to create a Note
// TODO: replace the name of this!
type NoteBind struct {
	Title    string `json:"title"`
	Contents string `json:"contents"`
	Author   string `json:"author"`
}

// Note
// ID is a GUID
// LastUpdated is a time.Time in UTC
type Note struct {
	ID          *guid.Guid `json:"-"`
	Title       string     `json:"title"`
	Contents    string     `json:"contents"`
	Author      string     `json:"author"`
	LastUpdated time.Time  `json:"last_updated"`
	Version     int        `json:"version"`
	Active      bool       `json:"active"`
}

// NewNote will create a new note from the title, author, and contents.
// LastUpdated, Version, and Active will be set by default
// Retuns a *Note
func NewNote(title, author, contents string) *Note {
	// TODO: Move Version to a db, or somewhere else..
	return &Note{
		ID:          guid.New(),
		Title:       title,
		Contents:    contents,
		Author:      author,
		LastUpdated: time.Now().UTC(),
		Version:     1,
		Active:      true,
	}
}

// NewNoteFromBind will create a new note from a mapped NoteBind
func NewNoteFromBind(nb NoteBind) *Note {
	note := NewNote(nb.Title, nb.Author, nb.Contents)
	return note
}

// MarshalJSON customizes the JSON marshaling for the Note struct.
// The ID field is represented as a string in the JSON output.
func (n *Note) MarshalJSON() ([]byte, error) {
	type Alias Note
	return json.Marshal(&struct {
		ID string `json:"id"`
		*Alias
	}{
		ID:    n.ID.String(),
		Alias: (*Alias)(n),
	})
}
