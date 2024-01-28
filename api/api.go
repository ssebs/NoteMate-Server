// api.go - REST API for PadPal-Server
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

	// TODO: replace_me with an env var or CLI flag
	provider := data.NewSampleProvider()

	// init gin + HTTP handlers
	router := gin.Default()
	initHandlers(router, provider)

	// Run the server
	log.Fatal(router.Run(hostPort))
}

// initHandlers is where to define new HTTP handlers
// requires the current router, and a CRUDProvider
func initHandlers(router *gin.Engine, provider data.CRUDProvider) {
	// Default handler
	router.GET("/", rootHandler)
	// Notes handlers
	router.GET("/notes", GETNotesHandler(provider))
	router.POST("/notes", POSTNotesHandler(provider))

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
