package sourcesapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
	r = append(r, utils.Route{Name: name + "Find", Method: strings.ToUpper("Get"), Pattern: "/api/v1/" + api, HandlerFunc: uac.Find})
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

// Find -
// Takes queryParams from the URL and uses them to build a search
// e.g. /api/v1/users?Email=dave.jones@maersk.com will search for matches on User.Email=dave.jones
// e.g. /api/v1/users?Email=dave.jones@maersk.com?Email=a.b@c.com will apply an Or to the email addresses
func (uac *SourcesApiController) Find(w http.ResponseWriter, r *http.Request) {

	/*
		Repeating some info on url.Values to show how they work

		v.Set("name", "Ava")
		v.Add("friend", "Jess")
		// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
		fmt.Println(v.Get("name"))
		fmt.Println(v.Get("friend"))
		fmt.Println(v["friend"])
	*/

	this := uac.Service

	// boilerplate:
	var err error
	var params url.Values
	params = r.URL.Query()

	result, err := this.Find(params)

	if err != nil {
		observability.Logger("Error", fmt.Sprintf("%v", err))
		w.WriteHeader(500)
		return
	}

	utils.EncodeJSONResponse(result, nil, w)
}
