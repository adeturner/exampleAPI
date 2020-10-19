package sourcesapi

import (
	"fmt"

	"github.com/adeturner/exampleAPI/sharedinterfaces"
	"github.com/adeturner/observability"
	"github.com/adeturner/persistenceServices"

	"github.com/google/uuid"
)

type SourcesapiService struct {
	persistenceLayer *persistenceServices.PersistenceLayer
}

// NewSourcesapiService creates a default api service
func NewSourcesapiService(docType sharedinterfaces.DocumentType) *SourcesapiService {

	observability.Logger("Info", "New Sources API Servicer")
	persistenceLayer, err := persistenceServices.GetPersistenceLayer(docType)

	if err != nil {
		observability.Logger("Error", fmt.Sprintf("NewSourcesapiService.GetPersistenceLayer : ERROR=%v", err))
	}
	return &SourcesapiService{persistenceLayer}
}

// AddSource -
// curl -d '{"name": "value1", "tag": "value2"}' -X POST http://localhost:8080/api/v1/Sources
func (srv *SourcesapiService) Add(i interface{}) (interface{}, error) {
	// generate a uuid
	UUID := uuid.New().String()
	newSource := i.(Source) // type assertion
	newSource.Id = UUID
	return nil, srv.persistenceLayer.AddDocument(UUID, newSource)
}

// UpdateSource -
// curl -d '{"id": "", "name": "value1", "tag": "value2"}' -X PUT http://localhost:8080/api/v1/Sources/{id}
func (srv *SourcesapiService) Update(id string, i interface{}) (interface{}, error) {
	return nil, srv.persistenceLayer.UpdateDocument(id, i.(Source))
}

// DeleteSource -
// curl -v -X DELETE http://localhost:8080/api/v1/Sources/f0518bfd-06c4-4c7d-9ace-5353f5bdd8e4
func (srv *SourcesapiService) Delete(id string) (interface{}, error) {

	// persistenceLayer.DeleteDocument copes with Id being passed as a string, or as the value struct
	// ensure both are populated
	s := Source{Id: id}
	return nil, srv.persistenceLayer.DeleteDocument(id, s)
}

// FindSourceById -
// curl -v -X GET http://localhost:8080/api/v1/Sources/dd50a804-45c4-454c-ac3e-839d8adddd48
func (srv *SourcesapiService) FindById(id string) (interface{}, error) {

	return srv.persistenceLayer.FindById(id, Source{})
}

/*
FindSources -
curl -v -X GET http://localhost:8080/api/v1/Sources?tags=value2\&limit=4

rpc error: code = FailedPrecondition desc = The query requires an index. You can create it here

gcloud firestore indexes composite create --collection-group=Sources --field-config field-path=Tag,order=ascending --field-config field-path=Name,order=ascending --async


-- takes several minutes even when empty, run the following until it completes

gcloud firestore indexes composite list
┌──────────────┬──────────────────┬─────────────┬───────┬─────────────┬───────────┬──────────────┐
│     NAME     │ COLLECTION_GROUP │ QUERY_SCOPE │ STATE │ FIELD_PATHS │   ORDER   │ ARRAY_CONFIG │
├──────────────┼──────────────────┼─────────────┼───────┼─────────────┼───────────┼──────────────┤
│ CICAgJim14AK │ Sources            │ COLLECTION  │ READY │ tag         │ ASCENDING │              │
│              │                  │             │       │ name        │ ASCENDING │              │
└──────────────┴──────────────────┴─────────────┴───────┴─────────────┴───────────┴──────────────┘

gcloud firestore indexes composite delete CICAgJim14AK
*/

func (srv *SourcesapiService) FindByTags(tags []string, strlimit string) (interface{}, error) {
	var apiType Source
	var apiTypes []Source
	return srv.persistenceLayer.FindByTags(tags, strlimit, apiType, apiTypes)
}

// PingGet - Server heartbeat operation
func (srv *SourcesapiService) PingGet() (interface{}, error) {
	observability.Logger("Info", "PingGet success")
	return nil, nil
}
