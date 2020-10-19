package sourcesapi

import "net/http"

// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type SourcesApiRouter interface {
	Add(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	FindById(http.ResponseWriter, *http.Request)
	FindByTags(http.ResponseWriter, *http.Request)
	PingGet(http.ResponseWriter, *http.Request)
}

// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type SourcesApiServicer interface {
	Add(interface{}) (interface{}, error)
	Update(string, interface{}) (interface{}, error)
	Delete(string) (interface{}, error)
	FindById(string) (interface{}, error)
	FindByTags([]string, string) (interface{}, error)
	PingGet() (interface{}, error)
}
