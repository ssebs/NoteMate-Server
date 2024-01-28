// notes_rt.go - /notes/ routing handlers
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ssebs/padpal-server/data"
)

/*
GET
	/notes?qry=
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
	// provider.SaveNote(data.NewNote("test1", "apiuser", "# test1"))
	// provider.SaveNote(data.NewNote("test2", "apiuser", "# test2"))
	return func(c *gin.Context) {
		notes, err := provider.ListNotes("")
		if err != nil {
			c.Error(err)
			c.JSON(404, err)
			return
		}
		c.JSON(200, notes)
	}
}

func GETNoteHandler(provider data.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// POST //
func POSTNotesHandler(provider data.CRUDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Map post data to NoteBind, then create Note from that
		var nb data.NoteBind
		if err := c.ShouldBind(&nb); err != nil {
			c.Error(err)
			c.JSON(500, err)
			return
		}
		note := data.NewNoteFromBind(nb)

		// Save the new note
		err := provider.SaveNote(note)
		if err != nil {
			c.Error(err)
			c.JSON(500, err)
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
