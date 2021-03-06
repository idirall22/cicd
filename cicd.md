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
11. Adding secrets using shell manualy.
12. install helm (helm client and tiller the server) because by default GCP enable RBAC.
13. create service account for tiller on kube-system namespace
14. create a cluster role binding to allow tiller service account to make changes on ower cluster.
15. helm init --service-account tiller --upgrade
16. install nginx-ingress

## Setup HTTPS
1. buy a domain from domain.google.com
2. install cert-manager usin helm get https setup automatically.
3. create an issuer.yaml and certificate.yaml files that will be used by cert-manager to generate cert.
4. update ingress to tell it to use certs.
