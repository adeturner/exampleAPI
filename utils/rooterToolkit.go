package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// EncodeJSONResponse uses the json encoder to write an interface to the http response with an optional status code
func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	var cors = os.Getenv("CORS_ENABLE")
	if cors == "true" {
		enableCors(&w)
	}

	return json.NewEncoder(w).Encode(i)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// ReadFormFileToTempFile reads file data from a request form and writes it to a temporary file
func ReadFormFileToTempFile(r *http.Request, key string) (*os.File, error) {
	r.ParseForm()
	formFile, _, err := r.FormFile(key)
	if err != nil {
		return nil, err
	}

	defer formFile.Close()
	file, err := ioutil.TempFile("tmp", key)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	fileBytes, err := ioutil.ReadAll(formFile)
	if err != nil {
		return nil, err
	}

	file.Write(fileBytes)
	return file, nil
}

// parseIntParameter parses a sting parameter to an int64
func parseIntParameter(param string) (int64, error) {
	return strconv.ParseInt(param, 10, 64)
}
