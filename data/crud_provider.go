// crud_provider.go = Create, Read, Update, Delete Providers.
// More funcs are available than just CRUD...
// example of a provider: file, google drive, etc.
package data

import "github.com/beevik/guid"

// Interface for CRUD stuff, if you want to save / load / update a file then use this.
// If you want to add your own storage mechanism, implement this.
type CRUDProvider interface {
	// CREATE //
	SaveNote(note *Note) error

	// READ //
	ListNotes(query string) ([]*Note, error)
	LoadNote(id guid.Guid) (*Note, error)

	// UPDATE //
	UpdateNote(id guid.Guid, updatedNote *Note) error

	// DELETE //
	DeleteNote(id guid.Guid) error
}
