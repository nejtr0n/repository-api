package ports

import (
	"net/http"

	"github.com/nejtr0n/repository-api/src/app"
	"github.com/nejtr0n/repository-api/src/app/command"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{app: app}
}

type HttpServer struct {
	app app.Application
}

func (h HttpServer) createProjectFromDull(c *gin.Context) {
	var cmd command.CreateProjectFromDull

	err := c.BindJSON(&cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.app.Commands.CreateProjectFromDull.Handle(c, cmd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
	return
}

func NewRouter(server HttpServer) *gin.Engine {
	router := gin.Default()
	router.POST("/create-project-from-dull", server.createProjectFromDull)

	return router

}
