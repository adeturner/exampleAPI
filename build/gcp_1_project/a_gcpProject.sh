# gcloud auth login

. gcp_env.sh

gcloud projects create $GCP_PROJECT

export PROJECT_NUMBER=`gcloud projects describe ${GCP_PROJECT} --format text | grep projectNumber | awk '{print $NF}'`

gcloud config set project ${GCP_PROJECT}

# works only if we have one billing acccount!
export BILLING_ID=`gcloud beta billing accounts list --format text | grep ^name | awk 'BEGIN {FS="/"} {print $NF}'`

gcloud beta billing projects link ${GCP_PROJECT} --billing-account=${BILLING_ID}



