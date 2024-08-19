// sample_provider.go
package data

import (
	"errors"
	"fmt"
	"strings"
	"sync"

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
	query = strings.ToLower(query)

	// Do the query
	// TODO: make this better
	for _, note := range p.notes {
		if query == "" ||
			strings.Contains(strings.ToLower(note.Title), query) ||
			strings.Contains(strings.ToLower(note.Contents), query) ||
			strings.Contains(strings.ToLower(note.Author), query) {
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

	return nil, fmt.Errorf("note %s not found", id.String())
}

// UpdateNote updates a note in the in-memory provider
func (p *SampleProvider) UpdateNote(id guid.Guid, updatedNote *Note) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if the note exists
	if _, exists := p.notes[id]; !exists {
		return fmt.Errorf("note %s not found", id)
	}

	// Update the note
	p.notes[id] = updatedNote
	return nil
}

// DeleteNote deletes a note (archives it) by ID
func (p *SampleProvider) DeleteNote(id guid.Guid) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if the note exists
	if _, exists := p.notes[id]; !exists {
		return fmt.Errorf("note %s not found", id)
	}

	return nil
}
