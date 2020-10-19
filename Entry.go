package exampleAPI

import (
	"log"
	"net/http"

	"github.com/adeturner/observability"
)

var router = NewRouter(NewDefaultApiController(NewDefaultApiService()))

// StartLocal -
func Prepare() {
	log.Printf("Server started")
}

// FunctionEntry -
// GCP expects signature func(http.ResponseWriter, *http.Request)
func FunctionEntry(w http.ResponseWriter, r *http.Request) {
	observability.SetAppName("exampleapi")
	observability.LogEnvVars()
	router.ServeHTTP(w, r)

}

func LocalEntry() {
	observability.SetAppName("exampleapi")
	observability.LogEnvVars()
	log.Fatal(http.ListenAndServe(":8080", router))
}
