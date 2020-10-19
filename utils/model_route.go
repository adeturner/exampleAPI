package utils

import "net/http"

// A Route defines the parameters for an api endpoint
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
