package sourcesapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/adeturner/exampleAPI/utils"
	"github.com/adeturner/observability"
	"github.com/gorilla/mux"
)

type SourcesApiController struct {
	Service SourcesApiServicer
}

// AddRoutes -
func (uac *SourcesApiController) AddRoutes(r utils.Routes) utils.Routes {
	// customise
	name := "Sources"
	api := strings.ToLower(name)

	// boilerplate
	r = append(r, utils.Route{Name: name + "Add", Method: strings.ToUpper("Post"), Pattern: "/api/v1/" + api, HandlerFunc: uac.Add})
	r = append(r, utils.Route{Name: name + "Update", Method: strings.ToUpper("Put"), Pattern: "/api/v1/" + api + "/{id}", HandlerFunc: uac.Update})
	r = append(r, utils.Route{Name: name + "Delete", Method: strings.ToUpper("Delete"), Pattern: "/api/v1/" + api + "/{id}", HandlerFunc: uac.Delete})
	r = append(r, utils.Route{Name: name + "FindById", Method: strings.ToUpper("Get"), Pattern: "/api/v1/" + api + "/{id}", HandlerFunc: uac.FindById})
	r = append(r, utils.Route{Name: name + "FindByTags", Method: strings.ToUpper("Get"), Pattern: "/api/v1/" + api, HandlerFunc: uac.FindByTags})
	r = append(r, utils.Route{Name: name + "PingGet", Method: strings.ToUpper("Get"), Pattern: "/api/v1/" + api + "/ping", HandlerFunc: uac.PingGet})
	return r
}

// AddSource -
func (uac *SourcesApiController) Add(w http.ResponseWriter, r *http.Request) {

	// customise
	n := Source{} // was SourceDetails
	this := uac.Service

	// boilerplate
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		observability.Logger("Error", fmt.Sprintf("Decode: %v", err))
	} else if err == nil {

		result, err := this.Add(n)

		if err != nil {
			observability.Logger("Error", fmt.Sprintf("%v", err))
			w.WriteHeader(500)
		} else {
			utils.EncodeJSONResponse(result, nil, w)
		}
	}
}

// Update -
func (uac *SourcesApiController) Update(w http.ResponseWriter, r *http.Request) {

	// customise
	u := Source{}
	this := uac.Service

	// boilerplate
	var result interface{}
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&u)

	if u.Id == "" {
		u.Id = id
	} else if u.Id != id {
		err = errors.New(fmt.Sprintf("Path id=%s does not match body id=%s", id, u.Id))
	}

	if err == nil {
		result, err = this.Update(id, u)
	}

	if err != nil {
		observability.Logger("Error", fmt.Sprintf("%v", err))
		w.WriteHeader(500)
	} else {
		utils.EncodeJSONResponse(result, nil, w)
	}
}

// Delete -
func (uac *SourcesApiController) Delete(w http.ResponseWriter, r *http.Request) {

	// customise
	this := uac.Service

	// boilerplate
	params := mux.Vars(r)
	id := params["id"]

	result, err := this.Delete(id)
	if err != nil {
		observability.Logger("Error", fmt.Sprintf("Delete: %v", err))
		w.WriteHeader(500)
		return
	}

	utils.EncodeJSONResponse(result, nil, w)
}

// FindById -
func (uac *SourcesApiController) FindById(w http.ResponseWriter, r *http.Request) {

	// customise
	this := uac.Service

	// boilerplate
	params := mux.Vars(r)
	id := params["id"]
	result, err := this.FindById(id)
	if err != nil {
		observability.Logger("Error", fmt.Sprintf("FindById: %v", err))
		w.WriteHeader(500)
		return
	}

	utils.EncodeJSONResponse(result, nil, w)
}

// FindByTags -
func (uac *SourcesApiController) FindByTags(w http.ResponseWriter, r *http.Request) {

	// customise
	this := uac.Service

	// boilerplate
	query := r.URL.Query()
	tags := strings.Split(query.Get("tags"), ",")
	limit := query.Get("limit")

	result, err := this.FindByTags(tags, limit)
	if err != nil {
		observability.Logger("Error", fmt.Sprintf("FindByTags: %v", err))
		w.WriteHeader(500)
		return
	}

	utils.EncodeJSONResponse(result, nil, w)
}

// PingGet - Server heartbeat operation
func (uac *SourcesApiController) PingGet(w http.ResponseWriter, r *http.Request) {

	// customise
	this := uac.Service

	// boilerplate
	result, err := this.PingGet()

	if err != nil {
		observability.Logger("Error", fmt.Sprintf("PingGet: err=%v", err))
		w.WriteHeader(500)
		return
	}

	utils.EncodeJSONResponse(result, nil, w)
}
