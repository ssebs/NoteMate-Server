// api.go - REST API for PadPal-Server
package api

import (
	"os"

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

	a.router.GET("/notes", GETNotesHandler(a.provider))
	a.router.GET("/notes/:id", GETNoteByIDHandler(a.provider))
	a.router.POST("/notes", POSTNotesHandler(a.provider))

	return a
}
func (a *API) RunAPI() error {
	return a.router.Run(a.hostPort)
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
