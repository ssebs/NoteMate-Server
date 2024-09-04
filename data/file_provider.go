// file_provider.go
package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/beevik/guid"
	"github.com/ssebs/padpal-server/util"
)

/*
Init:
- Load metadata from file/db
- Load local files
- Add metadata to state

Save note:
- Validate
- Save file locally
- Create metadata in state
- <notify clients of new file>
- Sync metadata to file/db

Update note:
- Is metadata found
- No: Save note
- Yes:
  - Validate
  - Save file locally
  - Update metadata in state
  - <notify clients of updated file>
  - Sync metadata to file/db

Get note by id:
- Meta param?
- Yes: send metadata
- No:
  - Upload/send local file

Get notes (query):
- Query / filter metadata
- Send metadata

*/

// FileProvider implements CRUDProvider
type FileProvider struct {
	notes      map[string]*Note // guid string
	mutex      sync.RWMutex
	dirName    string
	configPath string
}

func NewFileProvider(dirName, jsonConfigFullPath string) *FileProvider {
	fp := &FileProvider{
		notes:      make(map[string]*Note),
		dirName:    path.Clean(dirName),
		configPath: jsonConfigFullPath,
	}
	// Load metadata from file

	// List local files

	// Compare metadata list to local files list

	// Update metadata if needed

	// Load local files with LoadNote()

	return fp
}

func (p *FileProvider) loadMetadata() error {
	return nil
}

func (p *FileProvider) saveMetadata() error {
	// Save notes as json to file
	file, err := os.Create(p.configPath)
	if err != nil {
		return fmt.Errorf("could not create file %s", err)
	}

	jsonData, err := json.Marshal(p.notes)
	if err != nil {
		return fmt.Errorf("could not marshall notes into json file %s", err)
	}
	_, err = file.Write(jsonData)
	return err
}

func (p *FileProvider) SaveNote(note *Note) error {
	/*
		- Validate
		- Save file locally
		- Create metadata in state
		- <notify clients of new file>
	*/
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Validate
	if err := util.ValidateNoteContents(note.Contents); err != nil {
		return fmt.Errorf("could not validate note contents, %e", err)
	}

	// Check if a note with the same ID already exists
	if p.checkNoteFound(note.ID.String()) {
		return fmt.Errorf("note with the same ID already exists, id: %s", note.ID)
	}

	// Save file
	fn := fmt.Sprintf("./%s/%s.md", p.dirName, note.ID)
	if err := os.WriteFile(fn, []byte(note.Contents), 0644); err != nil {
		if e, ok := err.(*os.PathError); ok {
			return fmt.Errorf("have you created the ./%s dir? err: %s", p.dirName, e)
		}

		return fmt.Errorf("error writing file: %s", err.Error())
	}

	// Create metadata in state
	p.notes[note.ID.String()] = note
	if err := p.saveMetadata(); err != nil {
		// TODO: logger warn
		fmt.Println(err)
		return err
	}

	// TODO: notify all clients

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
		if query == "" ||
			strings.Contains(strings.ToLower(note.Title), query) ||
			strings.Contains(strings.ToLower(note.Contents), query) ||
			strings.Contains(strings.ToLower(note.Author), query) {
			result = append(result, note)
		}
	}
	// TODO: candidate for unit test..
	if len(result) == 0 {
		return result, fmt.Errorf("could not find any notes from the query: %s", query)
	}
	return result, nil
}

func (p *FileProvider) LoadNote(id guid.Guid) (*Note, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	// Find the note by ID
	if note, exists := p.notes[id.String()]; exists {
		return note, nil
	}

	return nil, fmt.Errorf("note %s not found", id.String())
}

func (p *FileProvider) UpdateNote(id guid.Guid, updatedNote *Note) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if the note exists
	if !p.checkNoteFound(id.String()) {
		return fmt.Errorf("note %s not found", id)
	}

	// Update the note
	p.notes[id.String()] = updatedNote
	return nil
}

func (p *FileProvider) DeleteNote(id guid.Guid) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if the note exists
	if !p.checkNoteFound(id.String()) {
		return fmt.Errorf("note %s not found", id)
	}

	// Mark the note as inactive
	return nil
}

func (p *FileProvider) checkNoteFound(id string) bool {
	if _, exists := p.notes[id]; !exists {
		return false
	}
	return true
}
