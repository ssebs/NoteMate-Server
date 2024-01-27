package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ssebs/padpal-server/data"
	"github.com/ssebs/padpal-server/util"
)

// HandleAndServe will handle the routes and serve HTTP
// contains a list of route handlers
func HandleAndServe(host string, port int) {
	hostPort := fmt.Sprintf("%s:%d", host, port)

	router := gin.Default()

	// Handlers
	router.GET("/", rootHandler)
	router.GET("/notes", notesHandler)

	// Run & log the server once it errors
	log.Fatal(router.Run(hostPort))
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

func notesHandler(c *gin.Context) {
	c.JSON(200, data.GetSampleNotes())
}
