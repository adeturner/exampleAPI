
# gcloud auth login

. build/gcp_env.sh

gcloud iam service-accounts create ${API_LOWERCASE}-serviceaccount --description="Service account for ${API}" --display-name="${API}-serviceaccount"

gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
   --member "serviceAccount:${API_LOWERCASE}-serviceaccount@${GCP_PROJECT}.iam.gserviceaccount.com" --role "roles/datastore.owner"

gcloud projects add-iam-policy-binding ${GCP_PROJECT} \
   --member "serviceAccount:${API_LOWERCASE}-serviceaccount@${GCP_PROJECT}.iam.gserviceaccount.com" --role "roles/pubsub.admin"

# must be executed in the root directory of the project!
gcloud functions deploy ${API} \
 --entry-point FunctionEntry \
 --trigger-http  \
 --region ${FUNCTIONS_REGION} \
 --service-account ${API_LOWERCASE}-serviceaccount@${GCP_PROJECT}.iam.gserviceaccount.com \
 --runtime go113 \
 --clear-labels  \
 --set-env-vars CORS_ENABLE=true,GCP_PROJECT=${GCP_PROJECT},CLOUDEVENT_DOMAIN=${CLOUDEVENT_DOMAIN},USE_FIRESTORE=true,USE_PUBSUB=true,USE_CQRS=true,DEBUG=true \
 --update-labels env=smoketest,version=0_1 \
  --quiet

# gcloud functions delete ${API} --region ${$FUNCTIONS_REGION}



