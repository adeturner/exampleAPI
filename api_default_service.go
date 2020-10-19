package exampleAPI

import (
	"github.com/adeturner/exampleAPI/sourcesapi"
	"github.com/adeturner/observability"
)

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
	// {templates}apisvc        {templates}api.{Templates}apiService
	sourcessapisvc sourcesapi.SourcesapiService
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() *DefaultApiService {

	observability.SetAppName("exampleapi")
	observability.Logger("Info", "New default API Servicer")

	// t := {templates}api.New{Templates}apiService(DOCUMENT_TYPE_SOURCES)

	u := sourcesapi.NewSourcesapiService(DOCUMENT_TYPE_SOURCES)

	d := DefaultApiService{*u}

	return &d
}
