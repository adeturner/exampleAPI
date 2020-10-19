

. gcp_env.sh

export CLOUD_RUN_SERVICE_NAME=espv2-cloudrun-${API_LOWERCASE}
export CLOUD_RUN_SERVICE_URL=`gcloud run services describe ${CLOUD_RUN_SERVICE_NAME} --platform=managed --format text | grep status.address.url | awk '{print $NF}'`
export CLOUD_RUN_HOSTNAME=`echo $CLOUD_RUN_SERVICE_URL | sed 's/https:\/\///'`

export ACCESS_TOKEN=`curl -s --request POST --url ${AUTH0_TOKEN_ENDPOINT} --header 'content-type: application/json' --data "@${AUTH0_SECRET}" | jq -r '.access_token'`

curl --request GET --header "Authorization: Bearer ${ACCESS_TOKEN}" --url https://${CLOUD_RUN_HOSTNAME}/api/v1/sources?tags=value2\&limit=4

hey -n 10 -c 5 -m GET -H "Authorization: Bearer ${ACCESS_TOKEN}" https://${CLOUD_RUN_HOSTNAME}/api/v1/sources?tags=value2\&limit=4


