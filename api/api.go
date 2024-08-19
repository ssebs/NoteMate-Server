// api.go - REST API for PadPal-Server
package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ssebs/padpal-server/data"
	"github.com/ssebs/padpal-server/util"
)

// DoEverything takes a hostPort (e.x. 0.0.0.0:8080)
func DoEverything(hostPort string) error {
	// TODO: replace_me with an env var or CLI flag
	provider := data.NewFileProvider()

	router := gin.Default()
	router.GET("/", rootHandler)
	router.GET("/notes", GETNotesHandler(provider))
	router.GET("/notes/:id", GETNoteByIDHandler(provider))
	router.POST("/notes", POSTNotesHandler(provider))

	// Run the server
	return router.Run(hostPort)
}

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
	c.JSON(code, err.Error())
}
