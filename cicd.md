## Steps to build cicd pipline on GCP
1. Install google cloud sdk.
2. env set `CLOUDSDK_CORE_DISABLE_PROMPTS=1` and `GIT_SHA=$(git rev-parse HEAD)`
2. Auth to gcloud using service acount (`gcloud auth activate-service-account --key-file service-account.json`).
3. create service account on GCP and give to this account permission to manage k8s cluster and storage to be able to push images.
4. set project_id (`gcloud config set project "project id"`).
5. set zone (`gcloud config set compute/zone "zone"`).
6. tell to gcloud command which k8s cluster should interact with  (`gcloud container clusters get-credentials "name of the cluster"`).
7. make a script where to build images.
8. push images to gcp container.
9. kubectl apply -f k8s
10. kubectl se image deployments/server 

