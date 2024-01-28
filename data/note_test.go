// note_test.go
package data

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/beevik/guid"
)

func TestNewNote(t *testing.T) {
	_title := "Test Title"
	_author := "Test Author"
	_contents := "Test Contents"

	note := NewNote(_title, _author, _contents)

	if note.ID == nil {
		t.Error("expected non-nil ID, got nil")
	}
	if note.Title != _title {
		t.Errorf("expected title %s, got %s", _title, note.Title)
	}
	if note.Contents != _contents {
		t.Errorf("expected contents %s, got %s", _contents, note.Contents)
	}
	if note.Author != _author {
		t.Errorf("expected author %s, got %s", _author, note.Author)
	}
	if note.LastUpdated.IsZero() {
		t.Error("expected non-zero LastUpdated time, got zero time")
	}
	if note.Version != 1 {
		t.Errorf("expected version 1, got %d", note.Version)
	}
	if !note.Active {
		t.Error("expected Active to be true, got false")
	}
}

func TestMarshalJSON(t *testing.T) {
	_id, _ := guid.ParseString("123e4567-e89b-12d3-a456-426614174001")
	_title := "Test Title"
	_author := "Test Author"
	_contents := "Test Contents"
	_lastUpdated := time.Date(2024, 1, 27, 12, 0, 0, 0, time.UTC)

	note := &Note{
		ID:          _id,
		Title:       _title,
		Contents:    _contents,
		Author:      _author,
		LastUpdated: _lastUpdated,
		Version:     1,
		Active:      true,
	}

	expectedJSON := `{"id":"123e4567-e89b-12d3-a456-426614174001","title":"Test Title","contents":"Test Contents","author":"Test Author","last_updated":"2024-01-27T12:00:00Z","version":1,"active":true}`

	resultJSON, err := json.Marshal(note)
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
	}

	if string(resultJSON) != expectedJSON {
		t.Errorf("Expected JSON:\n%s\nGot:\n%s", expectedJSON, string(resultJSON))
	}
}
