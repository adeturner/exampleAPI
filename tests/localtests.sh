

export DEBUG=true
export USE_FIRESTORE=true
export USE_PUBSUB=true
export USE_CQRS=true
export STDERR_THRESHOLD=true

go run cmd/main.go

export TESTAPI=sources
export TESTNAME=beento
export TESTUPDATEDNAME=today
export TESTTAGS=

curl -s -w "%{http_code}\n" -X GET http://localhost:8080/api/v1/${TESTAPI}
# 500

curl -s -w "%{http_code}\n" -d '{"name": "'${TESTNAME}'", "tag": "'${TESTTAGS}'"}' -X POST http://localhost:8080/api/v1/${TESTAPI}

curl -s -w "%{http_code}\n" -X GET http://localhost:8080/api/v1/${TESTAPI}?tags=${TESTTAGS}\&limit=4

export TESTID=`curl -s -X GET http://localhost:8080/api/v1/${TESTAPI}?tags=${TESTTAGS}\&limit=4 | sed 's/..*Id":"//' | sed 's/"..*//'`

curl -s -w "%{http_code}\n" -X GET http://localhost:8080/api/v1/${TESTAPI}/${TESTID}

curl -s -w "%{http_code}\n" -d '{"name": "'${TESTUPDATEDNAME}'", "tag": "'${TESTTAGS}'"}' -X PUT http://localhost:8080/api/v1/${TESTAPI}/${TESTID}

curl -s -w "%{http_code}\n" -X GET http://localhost:8080/api/v1/${TESTAPI}/${TESTID}

curl -s -w "%{http_code}\n" -X DELETE http://localhost:8080/api/v1/${TESTAPI}/${TESTID}

curl -s -w "%{http_code}\n" -X GET http://localhost:8080/api/v1/${TESTAPI}/${TESTID}
# 500


###########

curl -s -w "%{http_code}\n" -d '{"name": "foo", "tag": "bar"}' -X POST http://localhost:8080/api/v1/${TESTAPI}

curl -s -w "%{http_code}\n" -X GET http://localhost:8080/api/v1/${TESTAPI}?Name=beento

curl -s -w "%{http_code}\n" -X GET http://localhost:8080/api/v1/${TESTAPI}?Name=beento\&Name=foo

curl -s -w "%{http_code}\n" -X GET http://localhost:8080/api/v1/${TESTAPI}?Name=foo
