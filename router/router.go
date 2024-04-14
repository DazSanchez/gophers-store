package router

import (
	"github.com/gin-gonic/gin"
)

type api struct{}

// API defines routes under the /api path, intended to hold RESTful services.
var API api = api{}

// Init registers all configured routes over API's endpoint.
func (a api) Init(e *gin.Engine) {
	r := e.Group("api")

	addCategoryRoutes(r)
	addGopherRoutes(r)
	addTagRoutes(r)
	addPingRoutes(r)
}
