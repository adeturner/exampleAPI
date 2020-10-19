
. gcp_env.sh

gcloud projects list

# app engine seems to be mandatory
gcloud services enable appengine.googleapis.com  --project=${GCP_PROJECT}

gcloud alpha firestore databases create --project=${GCP_PROJECT} --region=${FIRESTORE_LOCATION_ID}

gcloud iam service-accounts create ${GCP_PROJECT}-firestore --display-name ${GCP_PROJECT}-firestore

gcloud projects add-iam-policy-binding ${GCP_PROJECT} --member "serviceAccount:${GCP_PROJECT}-firestore@${GCP_PROJECT}.iam.gserviceaccount.com" --role "roles/owner"

gcloud projects add-iam-policy-binding ${GCP_PROJECT} --member "serviceAccount:${GCP_PROJECT}-firestore@${GCP_PROJECT}.iam.gserviceaccount.com" --role "roles/datastore.owner"

mkdir -p ${SECRETS_DIR}

export GOOGLE_APPLICATION_CREDENTIALS=${FIRESTORE_SECRET}
rm $GOOGLE_APPLICATION_CREDENTIALS

gcloud iam service-accounts keys create ${GOOGLE_APPLICATION_CREDENTIALS} --iam-account ${GCP_PROJECT}-firestore@${GCP_PROJECT}.iam.gserviceaccount.com

