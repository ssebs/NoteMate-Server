// notes_rt.go - /notes/ routing handlers
package api

import (
	"fmt"

	"github.com/beevik/guid"
	"github.com/gin-gonic/gin"
	"github.com/ssebs/padpal-server/data"
)

/*
GET
	/notes?q=
	/notes/:id
	/notes/:id?version=
	TODO: /versions/notes
POST
	/notes
PUT
	/notes/:id
	/notes/:id?version=
DELETE
	/notes/:id
	TODO: /notes/:id?version=
*/

// GET //
func GETNotesHandler(provider data.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		qry := c.Query("q")

		notes, err := provider.ListNotes(qry)
		if err != nil {
			errorHandler(404, err, c)
			return
		}
		c.JSON(200, notes)
	}
}

func GETNoteByIDHandler(provider data.CRUDProvider) gin.HandlerFunc {
	// Get ID from param
	return func(c *gin.Context) {
		// Parse id as GUID if possible
		id, err := guid.ParseString(c.Param("id"))
		if err != nil {
			errorHandler(400, fmt.Errorf("invalid id given, could not convert to guid: err: %s", err), c)
			return
		}
		// Then get the note from that GUID & return
		note, err := provider.LoadNote(*id)
		if err != nil {
			errorHandler(404, err, c)
			return
		}
		c.JSON(200, note)
	}
}

// POST //
func POSTNotesHandler(provider data.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Map post data to NoteBind, then create Note from that
		var nb data.NoteBind
		if err := c.ShouldBind(&nb); err != nil {
			errorHandler(400, err, c)
			return
		}
		note := data.NewNoteFromBind(nb)

		// Save the new note
		err := provider.SaveNote(note)
		if err != nil {
			errorHandler(500, err, c)
			return
		}
		c.JSON(201, note)
	}
}

// PUT //
func PUTNoteHandler(provider data.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// DELETE //
func DELETENoteHandler(provider data.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
