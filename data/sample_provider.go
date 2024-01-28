// sample_provider.go
package data

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/beevik/guid"
)

// SampleProvider is an in-memory implementation of CRUDProvider
type SampleProvider struct {
	notes map[guid.Guid]*Note
	mutex sync.RWMutex
}

// NewSampleProvider creates a new SampleProvider instance
func NewSampleProvider() *SampleProvider {
	return &SampleProvider{
		notes: make(map[guid.Guid]*Note),
	}
}

// SaveNote saves a note to the in-memory provider
func (p *SampleProvider) SaveNote(note *Note) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if a note with the same ID already exists
	if _, exists := p.notes[*note.ID]; exists {
		return errors.New("note with the same ID already exists")
	}

	// Save the note to the map
	p.notes[*note.ID] = note
	return nil
}

// ListNotes lists all active notes that match the query
// Errors if there are no notes found
func (p *SampleProvider) ListNotes(query string) ([]*Note, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var result []*Note

	// Do the query
	for _, note := range p.notes {
		if note.Active && (query == "" || strings.Contains(note.Title, query) || strings.Contains(note.Contents, query)) {
			result = append(result, note)
		}
	}
	if len(result) == 0 {
		return result, fmt.Errorf("could not find any notes from the query: %s", query)
	}
	return result, nil
}

// LoadNote loads a note from the in-memory provider by ID
func (p *SampleProvider) LoadNote(id guid.Guid) (*Note, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	// Find the note by ID
	if note, exists := p.notes[id]; exists {
		return note, nil
	}

	return nil, fmt.Errorf("note %s not found", id)
}

// ListNoteVersions lists all versions of a note by ID
func (p *SampleProvider) ListNoteVersions(id guid.Guid) ([]int, error) {
	return nil, errors.New("method not implemented")
}

// LoadNoteVersion loads a specific version of a note by ID and version number
func (p *SampleProvider) LoadNoteVersion(id guid.Guid, version int) (*Note, error) {
	return nil, errors.New("method not implemented")
}

// UpdateNote updates a note in the in-memory provider
func (p *SampleProvider) UpdateNote(id guid.Guid, updatedNote *Note) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if the note exists
	if _, exists := p.notes[id]; !exists {
		return fmt.Errorf("note %s not found", id)
	}

	if p.notes[id].Version <= updatedNote.Version {
		updatedNote.Version += 1
	}

	// Update the note
	p.notes[id] = updatedNote
	return nil
}

// RestoreNote restores a Note to a specific version, append version #
func (p *SampleProvider) RestoreNote(id guid.Guid, version int) (*Note, error) {
	return nil, errors.New("method not implemented")
}

// DeleteNote deletes a note (archives it) by ID
func (p *SampleProvider) DeleteNote(id guid.Guid) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if the note exists
	if _, exists := p.notes[id]; !exists {
		return fmt.Errorf("note %s not found", id)
	}

	// Mark the note as inactive
	p.notes[id].Active = false
	return nil
}

// OLD
// GetSampleNotes will return sample notes
func GetSampleNotes() []Note {
	_now := time.Now().UTC()
	notes := []Note{
		{Title: "firstTitle", Contents: "# First Note!\nTest\n- Foo\n- Bar", Author: "Seb",
			LastUpdated: _now, Version: 1, Active: true, ID: guid.New()},
		{Title: "secondTitle", Contents: "# Second Note!\nTest\n- one\n- two", Author: "Seb",
			LastUpdated: _now, Version: 1, Active: true, ID: guid.New()},
		{Title: "thirdTitle", Contents: "# Third Note!\nTest\n- Nest\n\t- it", Author: "Seb",
			LastUpdated: _now, Version: 1, Active: true, ID: guid.New()},
	}

	return notes
}
