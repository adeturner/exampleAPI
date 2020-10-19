# ExampleAPI

Unsupported and currently not fully tested because I quickly made a copy of another project to create the example

That said I'm not expecting too much drama (if any) to get it working)

## Overview

- Implements CRUD APIs for a "Source" object
- Via a Golang API server with Gorilla Mux
- The API server an be run locally or as a Google Cloud Function (GCP Function) protected by Google Endpoints; or... tbd Azure Functions
- Persistence uses Google Cloud Firestore as a backend
- The code does not know the backend, access is abstracted via [https://github.com/adeturner/persistenceServices]

The project is setup to enable additional APIs to be added into the project:

- use the sourcesapi directory as boilerplate, just clone the directory and edit
- add the new api into root directory api_*go files
- add the api to doctype_model.go
- clone and edit the api/swagger2-sourcesapi*yaml
- customise as required

### Useful background

If you want just Google Cloud Firestore, any web searches that mention *Firebase* Firstore should be treated with a pinch of salt
... unless you really want it! There be devils lurking.

### API Routes

See the openapi spec in the API directory, for clarity the Gorilla Mux routes are added in sources_default.go

```go
r = append(r, utils.Route{Name: name + "Add", Method: strings.ToUpper("Post"), Pattern: "/api/v1/" + api, HandlerFunc: uac.Add})
r = append(r, utils.Route{Name: name + "Update", Method: strings.ToUpper("Put"), Pattern: "/api/v1/" + api + "/{id}", HandlerFunc: uac.Update})
r = append(r, utils.Route{Name: name + "Delete", Method: strings.ToUpper("Delete"), Pattern: "/api/v1/" + api + "/{id}", HandlerFunc: uac.Delete})
r = append(r, utils.Route{Name: name + "FindById", Method: strings.ToUpper("Get"), Pattern: "/api/v1/" + api + "/{id}", HandlerFunc: uac.FindById})
r = append(r, utils.Route{Name: name + "FindByTags", Method: strings.ToUpper("Get"), Pattern: "/api/v1/" + api, HandlerFunc: uac.FindByTags})
```

## Setting up

### Get the code

Install go

```bash
cd ${MYROOT}
mkdir pkg src bin
cd src
git clone https://github.com/adeturner/exampleAPI
```

Create a secret environment file...

```bash
cat > ~/secrets/apiv01-exampleAPI-env.sh << EOF
# Private file
export GOPATH="${MYROOT}"
export CGO_ENABLED=0
#
export GCP_PROJECT=myproject
export SECRETS_DIR=~/secrets
export CLOUDEVENT_DOMAIN=exampleapi.com
export GOOGLE_APPLICATION_CREDENTIALS=${SECRETS_DIR}/${GCP_PROJECT}-persistenceServices.json
export AUTH0_SECRET=${SECRETS_DIR}/auth0_client_secret.json
export FIRESTORE_SECRET=${SECRETS_DIR}/apiv01-firestore.json
export API=exampleAPI
export AUTH0_TOKEN_ENDPOINT=https://mydomain.eu.auth0.com/oauth/token
EOF
```

## Base infra setup

WARNING!

You need Firestore and Pubsub; the build directory has some scripts, they are not tested for automated build. Please run by hand until sure.

## Local testing

### Starting the server

To run the server locally, follow these simple steps:

```bash
export LOG_LEVEL=INFO
export USE_FIRESTORE=true
export USE_PUBSUB=true
export USE_CQRS=true
export STDERR_THRESHOLD=true

go run cmd/main.go
```

### Testing locally

tests/localtests.sh has some example curl commands; this script is an example and not fully automated so run line by line

## Running the server in GCP

### Deploy to Google Cloud Function

Follow the scripts [a, b, c, d] in gcp_3_Functions; these scripts are also not fully automated so run line by line

Note you can skip [b, c] if you just want to play with a publicly accessible Cloud Function. The API works fine.

Authentication needs an AUTH0 account or equivalent to get access tokens

Authentication is applied at GOOGLE ENDPOINTS through the concatenation/sed of the API/*yaml files in c

### Testing on GCP

tests/gcptests.sh has some example curl commands; this script is an example and not fully automated so run line by line

