// crud_providers = Save, Load, List, Update, Delete Providers.
// example of a provider: file, google drive, etc.
package data

import "github.com/beevik/guid"

// Interface for CRUD stuff, if you want to save / load / update a file then use this.
// If you want to add your own storage mechanism, implement this.
type CRUDProvider interface {
	// Single note stuff

	// Save note to disk
	SaveNote(note *Note) error
	// Load note from disk by guid ID
	LoadNote(id guid.Guid) (*Note, error)
	// Update note, append version
	UpdateNote(id guid.Guid, updatedNote *Note) error
	// Delete a note (archive it)
	DeleteNote(id guid.Guid) error

	// Version stuff
	// List all versions of a note
	ListNoteVersions(id guid.Guid) ([]int, error)
	// Load note from disk by guid ID + version
	LoadNoteVersion(id guid.Guid, version int) (*Note, error)
	// UpdateNote to a specific version of a note, append version
	RestoreNote(id guid.Guid, version int) (*Note, error)

	// Multi-Notes stuff
	ListNotes(query string) ([]*Note, error)
}
