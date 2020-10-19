. gcp_env.sh

gcloud config set project ${GCP_PROJECT}

gcloud services enable cloudbuild.googleapis.com
gcloud services enable cloudfunctions.googleapis.com
gcloud services enable containerregistry.googleapis.com
gcloud services enable pubsub.googleapis.com




