
export ESP_PROJECT_ID=$GCP_PROJECT
export CLOUD_RUN_SERVICE_NAME=espv2-cloudrun-${API_LOWERCASE}


export CLOUD_RUN_SERVICE_URL=`gcloud run services describe ${CLOUD_RUN_SERVICE_NAME} --platform=managed --format text | grep status.address.url | awk '{print $NF}'`
export CLOUD_RUN_HOSTNAME=`echo $CLOUD_RUN_SERVICE_URL | sed 's/https:\/\///'`


export ACCESS_TOKEN=`curl -s --request POST --url ${AUTH0_TOKEN_ENDPOINT} --header 'content-type: application/json' --data "@${AUTH0_SECRET}" | jq -r '.access_token'`


# {"code":401,"message":"Jwt is missing"}
# 401 unauthorised
curl -X GET https://${CLOUD_RUN_HOSTNAME}/api/v1/sources
curl -X GET https://${CLOUD_RUN_HOSTNAME}/api/v1/sources/${DOCREF}
curl -X DELETE https://${CLOUD_RUN_HOSTNAME}/api/v1/sources/${DOCREF}
curl -X GET https://${CLOUD_RUN_HOSTNAME}/api/v1/sources/${DOCREF}
curl -d '{"name": "myname1", "tag": "mytag"}' -X POST https://${CLOUD_RUN_HOSTNAME}/api/v1/sources
curl -X GET https://${CLOUD_RUN_HOSTNAME}/api/v1/sources?tags=value2\&limit=4




curl -s -X GET --header "Authorization: Bearer ${ACCESS_TOKEN}" --url https://${CLOUD_RUN_HOSTNAME}/api/v1/sources?tags=\&limit=4

curl -X POST --header "Authorization: Bearer ${ACCESS_TOKEN}" -d '{"name": "newade!", "tag": "newade!"}' --url https://${CLOUD_RUN_HOSTNAME}/api/v1/sources

curl -s -X GET --header "Authorization: Bearer ${ACCESS_TOKEN}" --url https://${CLOUD_RUN_HOSTNAME}/api/v1/sources?tags=newade\!\&limit=4

export DOCREF=`curl -s -X GET --header "Authorization: Bearer ${ACCESS_TOKEN}" --url https://${CLOUD_RUN_HOSTNAME}/api/v1/sources?tags=newade\!\&limit=4 | jq -r '.[].id'`

curl -X GET --header "Authorization: Bearer ${ACCESS_TOKEN}" --url https://${CLOUD_RUN_HOSTNAME}/api/v1/sources/${DOCREF}

curl -X PUT --header "Authorization: Bearer ${ACCESS_TOKEN}" -d '{"name": "updated newade!", "tag": "updated newade!"}' --url https://${CLOUD_RUN_HOSTNAME}/api/v1/sources/${DOCREF}

curl -X DELETE --header "Authorization: Bearer ${ACCESS_TOKEN}" --url https://${CLOUD_RUN_HOSTNAME}/api/v1/sources/${DOCREF}


