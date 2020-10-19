
. ../gcp_env.sh


export PROJECT_NUMBER=`gcloud projects describe ${GCP_PROJECT} --format text | grep projectNumber | awk '{print $NF}'`
export CLOUD_RUN_SERVICE_NAME=espv2-cloudrun-${API_LOWERCASE}
export CLOUD_RUN_SERVICE_URL=`gcloud run services describe ${CLOUD_RUN_SERVICE_NAME} --platform=managed --format text | grep status.address.url | awk '{print $NF}'`
export CLOUD_RUN_HOSTNAME=`echo $CLOUD_RUN_SERVICE_URL | sed 's/https:\/\///'`
export ESP_GCP_PROJECT=$GCP_PROJECT
export ESP_PROJECT_NUMBER=$PROJECT_NUMBER


# You must re-run this entire step:
# 1) Every time you modify and redeploy the Endpoints service configuration. Otherwise, changes to the service configuration will not be propagated to ESPv2 Beta.
# 2) When a new version of ESPv2 Beta is released. Otherwise, the older version of ESPv2 Beta will remain deployed.

cat $OPENAPI_YAML | sed 's|GOOGLE-CLOUD-ENDPOINT-URL|'"${CLOUD_RUN_HOSTNAME}"'|' \
                  | sed 's|FUNCTIONS_REGION-GCP_PROJECT|'"${FUNCTIONS_REGION}-${GCP_PROJECT}"'|' \
                  | sed 's|REPLACE_WITH_API|'"${API}"'|'> ${ENDPOINT_YAML}

egrep 'host|func|address' ${ENDPOINT_YAML}



gcloud endpoints services deploy ${ENDPOINT_YAML} --project $ESP_GCP_PROJECT

curl -o ${GCLOUD_BUILD_IMAGE} https://raw.githubusercontent.com/GoogleCloudPlatform/esp-v2/master/docker/serverless/gcloud_build_image
chmod +x ${GCLOUD_BUILD_IMAGE}



# this needs to change for subsequent deployments
export CONFIG_ID=`gcloud endpoints configs list --service ${CLOUD_RUN_HOSTNAME} --format text | grep ^id: | awk '{print $NF}' | head -1`
echo $CONFIG_ID
${GCLOUD_BUILD_IMAGE} -s $CLOUD_RUN_HOSTNAME -c $CONFIG_ID -p $ESP_GCP_PROJECT


# gcloud run revisions  list --platform managed --project=$ESP_GCP_PROJECT
# gcloud endpoints configs list --service ${CLOUD_RUN_HOSTNAME}
# gcloud endpoints services list --project $ESP_GCP_PROJECT
# gcloud endpoints services delete espv2-cloudrun-nj3vlrcugq-ew.a.run.app --project $ESP_GCP_PROJECT

export GCLOUD_BUILD_IMAGE_NAME=`gcloud container images list-tags gcr.io/${GCP_PROJECT}/endpoints-runtime-serverless \
                                --sort-by='~timestamp' --limit='1' --format='value(tags)' | tail -1`

export CLOUD_RUN_SERVICE_ACCOUNT=cloudrun-${API_LOWERCASE}-srvacct

gcloud run deploy $CLOUD_RUN_SERVICE_NAME \
  --image="gcr.io/${GCP_PROJECT}/endpoints-runtime-serverless:${GCLOUD_BUILD_IMAGE_NAME}" \
  --service-account=${CLOUD_RUN_SERVICE_ACCOUNT} \
  --platform=managed \
  --set-env-vars=ESPv2_ARGS=--cors_preset=basic \
  --project=$ESP_GCP_PROJECT

# gcloud run services list --platform=managed
# cloud iam service-accounts list

gcloud functions add-iam-policy-binding ${API} \
   --region ${FUNCTIONS_REGION} --member "serviceAccount:${ESP_PROJECT_NUMBER}-compute@developer.gserviceaccount.com" \
   --role "roles/cloudfunctions.invoker" --project ${GCP_PROJECT}

gcloud functions add-iam-policy-binding ${API} \
   --region ${FUNCTIONS_REGION} \
   --member "serviceAccount:${CLOUD_RUN_SERVICE_ACCOUNT}@${GCP_PROJECT}.iam.gserviceaccount.com" \
   --role "roles/cloudfunctions.invoker" --project ${GCP_PROJECT}

# gcloud functions get-iam-policy $API --region $FUNCTIONS_REGION
# bindings:
# - members:
#   - serviceAccount:1025159660614-compute@developer.gserviceaccount.com
#   - serviceAccount:cloudrun-usersapi-srvacct@apiv01.iam.gserviceaccount.com
#   role: roles/cloudfunctions.invoker

gcloud projects get-iam-policy $GCP_PROJECT


