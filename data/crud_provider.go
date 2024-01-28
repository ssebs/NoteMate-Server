// crud_provider.go = Create, Read, Update, Delete Providers.
// More funcs are available than just CRUD...
// example of a provider: file, google drive, etc.
package data

import "github.com/beevik/guid"

// Interface for CRUD stuff, if you want to save / load / update a file then use this.
// If you want to add your own storage mechanism, implement this.
type CRUDProvider interface {
	// CREATE //
	// Save note to disk
	SaveNote(note *Note) error
	// TODO: CopyNote(id guid.Guid) error

	// READ //
	// List all active notes
	ListNotes(query string) ([]*Note, error)
	// Load note from disk by guid ID
	LoadNote(id guid.Guid) (*Note, error)
	// List all versions of a note
	ListNoteVersions(id guid.Guid) ([]int, error)
	// Load note from disk by guid ID + version
	LoadNoteVersion(id guid.Guid, version int) (*Note, error)

	// UPDATE //
	// Update note to given data, append version #
	UpdateNote(id guid.Guid, updatedNote *Note) error
	// Restore a Note to a specific version, append version #
	RestoreNote(id guid.Guid, version int) (*Note, error)

	// DELETE //
	// Delete a note (archive it)
	DeleteNote(id guid.Guid) error
	// TODO: DeleteNoteVersion(id guid.Guid, version int) error
}
