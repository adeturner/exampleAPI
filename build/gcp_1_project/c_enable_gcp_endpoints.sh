
# gcloud auth login

. gcp_env.sh

export ESPV2_REGION=${FUNCTIONS_REGION}

gcloud config set project ${GCP_PROJECT}

gcloud services enable servicecontrol.googleapis.com
gcloud services enable endpoints.googleapis.com

gcloud services enable run.googleapis.com
gcloud config set run/region $ESPV2_REGION
