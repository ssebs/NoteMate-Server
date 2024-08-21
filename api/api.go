// api.go - REST API for PadPal-Server
package api

import (
	"fmt"
	"os"

	"github.com/beevik/guid"
	"github.com/gin-gonic/gin"
	"github.com/ssebs/padpal-server/data"
	"github.com/ssebs/padpal-server/util"
)

type API struct {
	hostPort string
	provider data.CRUDProvider
	router   *gin.Engine
}

func NewAPI(hostPort, dirName string) *API {
	a := &API{
		hostPort: hostPort,
		provider: data.NewFileProvider(dirName),
		router:   gin.Default(),
	}
	a.router.GET("/", rootHandler)

	a.router.POST("/notes", a.createNoteHandler())
	a.router.GET("/notes", a.getNotesHandler())
	a.router.GET("/notes/:id", a.getNoteByIDHandler())
	a.router.PUT("/notes/:id", a.updateNoteHandler())
	a.router.DELETE("/notes/:id", a.deleteNoteHandler())

	return a
}

func (a *API) RunAPI() error {
	return a.router.Run(a.hostPort)
}

func (a *API) createNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Map post data to NoteBind, then create Note from that
		var nb data.NoteBind
		if err := c.ShouldBind(&nb); err != nil {
			errorHandler(400, err, c)
			return
		}
		note := data.NewNoteFromBind(nb)

		// Save the new note
		err := a.provider.SaveNote(note)
		if err != nil {
			errorHandler(500, err, c)
			return
		}
		c.JSON(201, gin.H{"status": "Note Created", "note": note})
	}
}

func (a *API) getNotesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		qry := c.Query("q")

		notes, err := a.provider.ListNotes(qry)
		if err != nil {
			errorHandler(404, err, c)
			return
		}
		c.JSON(200, notes)
	}
}

func (a *API) getNoteByIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse id as GUID if possible
		id, err := guid.ParseString(c.Param("id"))
		if err != nil {
			errorHandler(400, fmt.Errorf("invalid id given, could not convert to guid: err: %s", err), c)
			return
		}
		// Then get the note from that GUID & return
		note, err := a.provider.LoadNote(*id)
		if err != nil {
			errorHandler(404, err, c)
			return
		}
		c.JSON(200, note)
	}
}

func (a *API) updateNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	}
}

func (a *API) deleteNoteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	}
}

// func (a *API) Handler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.JSON(200, gin.H{"status": "ok"})
// 	}
// }

// rootHandler renders the REST-API.md file as HTML for /
func rootHandler(c *gin.Context) {
	// get contents of ./REST-API.md and return
	md, err := os.ReadFile("./REST-API.md")
	if err != nil {
		c.Error(err)
		c.JSON(500, err)
		return
	}

	// convert to html
	html := util.ParseMDToHTML(md)
	c.Header("content-type", "text/html")
	c.String(200, string(html))
}

// errorHandler will log the error and return JSON with the message
func errorHandler(code int, err error, c *gin.Context) {
	c.Error(err)
	c.JSON(code, gin.H{"status": "Error", "message": err.Error()})
}
