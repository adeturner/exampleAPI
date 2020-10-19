package exampleAPI

import (
	"github.com/adeturner/exampleAPI/sourcesapi"
	"github.com/adeturner/exampleAPI/utils"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	// {templates}apiController       {templates}api.{Templates}ApiController
	sourcesapiController sourcesapi.SourcesApiController
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(d *DefaultApiService) Router {

	dac := DefaultApiController{}
	// dac.{templates}apiController = {templates}api.{Templates}ApiController{Service: &d.{templates}apisvc}
	dac.sourcesapiController = sourcesapi.SourcesApiController{Service: &d.sourcessapisvc}

	return &dac
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() utils.Routes {
	r := utils.Routes{}
	// r = c.{templates}apiController.AddRoutes(r)
	r = c.sourcesapiController.AddRoutes(r)
	return r
}
