
export FIRESTORE_LOCATION_ID=europe-west
export FUNCTIONS_REGION=europe-west1
export API_LOWERCASE=`echo $API | tr '[:upper:]' '[:lower:]'`
export OPENAPI_YAML=$GOPATH/src/${API}/api/swagger2-${API}.yaml
export ENDPOINT_YAML=/tmp/endpoint_${API}.yaml
export GCLOUD_BUILD_IMAGE=/tmp/gcloud_build_image