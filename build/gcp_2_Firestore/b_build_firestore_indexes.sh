
. gcp_env.sh

echo Note will pause, the composite create command takes several minutes even when the collection is empty

set -x

# required for c.FindUsers
gcloud firestore indexes composite create \
     --collection-group=Sources \
     --field-config field-path=Tag,order=ascending \
     --field-config field-path=Name,order=ascending

set +x


# alternate option
# add the --async option to the composite create above
# then run the following until it completes
# gcloud firestore indexes composite list

# ┌──────────────┬──────────────────┬─────────────┬───────┬─────────────┬───────────┬──────────────┐
# │     NAME     │ COLLECTION_GROUP │ QUERY_SCOPE │ STATE │ FIELD_PATHS │   ORDER   │ ARRAY_CONFIG │
# ├──────────────┼──────────────────┼─────────────┼───────┼─────────────┼───────────┼──────────────┤
# │ CICAgJim14AK │ Users            │ COLLECTION  │ READY │ tag         │ ASCENDING │              │
# │              │                  │             │       │ name        │ ASCENDING │              │
# └──────────────┴──────────────────┴─────────────┴───────┴─────────────┴───────────┴──────────────┘
# 
# gcloud firestore indexes composite delete CICAgJim14AK



