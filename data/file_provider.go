// file_provider.go
package data

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/beevik/guid"
)

// FileProvider implements CRUDProvider
type FileProvider struct {
	notes map[guid.Guid]*Note
	mutex sync.RWMutex
}

func NewFileProvider() *FileProvider {
	return &FileProvider{
		notes: make(map[guid.Guid]*Note),
	}
}

func (p *FileProvider) SaveNote(note *Note) error {
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

func (p *FileProvider) ListNotes(query string) ([]*Note, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var result []*Note
	query = strings.ToLower(query)

	// Do the query
	// TODO: make this better
	for _, note := range p.notes {
		if note.Active && (query == "" ||
			strings.Contains(strings.ToLower(note.Title), query) ||
			strings.Contains(strings.ToLower(note.Contents), query) ||
			strings.Contains(strings.ToLower(note.Author), query)) {
			result = append(result, note)
		}
	}
	if len(result) == 0 {
		return result, fmt.Errorf("could not find any notes from the query: %s", query)
	}
	return result, nil
}

func (p *FileProvider) LoadNote(id guid.Guid) (*Note, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	// Find the note by ID
	if note, exists := p.notes[id]; exists {
		return note, nil
	}

	return nil, fmt.Errorf("note %s not found", id.String())
}

func (p *FileProvider) ListNoteVersions(id guid.Guid) ([]int, error) {
	return nil, errors.New("method not implemented")
}

func (p *FileProvider) LoadNoteVersion(id guid.Guid, version int) (*Note, error) {
	return nil, errors.New("method not implemented")
}

func (p *FileProvider) UpdateNote(id guid.Guid, updatedNote *Note) error {
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

func (p *FileProvider) RestoreNote(id guid.Guid, version int) (*Note, error) {
	return nil, errors.New("method not implemented")
}

func (p *FileProvider) DeleteNote(id guid.Guid) error {
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
