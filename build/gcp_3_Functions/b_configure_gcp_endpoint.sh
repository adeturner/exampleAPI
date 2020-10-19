

# gcloud auth login

. gcp_env.sh

export ESP_GCP_PROJECT=$GCP_PROJECT
export CLOUD_RUN_SERVICE_NAME=espv2-cloudrun-${API_LOWERCASE}



gcloud config set project $GCP_PROJECT

# note option to run the following command with --service-account=...
gcloud run deploy ${CLOUD_RUN_SERVICE_NAME} \
                --image="gcr.io/endpoints-release/endpoints-runtime-serverless:2" \
                --allow-unauthenticated --platform managed --project=$ESP_GCP_PROJECT

# gcloud run services list --platform managed --project=$ESP_GCP_PROJECT

export CLOUD_RUN_SERVICE_URL=`gcloud run services describe ${CLOUD_RUN_SERVICE_NAME} --platform=managed --format text | grep status.address.url | awk '{print $NF}'`
export CLOUD_RUN_HOSTNAME=`echo $CLOUD_RUN_SERVICE_URL | sed 's/https:\/\///'`

export CLOUD_RUN_SERVICE_ACCOUNT=cloudrun-${API_LOWERCASE}-srvacct

gcloud iam service-accounts create ${CLOUD_RUN_SERVICE_ACCOUNT} --description="${CLOUD_RUN_SERVICE_ACCOUNT}" --display-name="${CLOUD_RUN_SERVICE_ACCOUNT}"
# gcloud iam service-accounts list

gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
           --member "serviceAccount:${CLOUD_RUN_SERVICE_ACCOUNT}@${GCP_PROJECT}.iam.gserviceaccount.com" \
           --role roles/servicemanagement.serviceController


