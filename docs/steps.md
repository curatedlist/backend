# Step-by-step guide
## Init the go modules
From https://blog.golang.org/using-go-modules
$ go mod init github.com/curatedlist/backend

## Code Layout
https://github.com/golang-standards/project-layout

## Install Gin server

From https://github.com/gin-gonic/gin#installation
$ go get -u github.com/gin-gonic/gin

Run it from:
$ go run cmd/backend/main.go

## Google Cloud

### Install the SDK
https://cloud.google.com/sdk/docs#mac
$ ./google-cloud-sdk/install.sh
$ ./google-cloud-sdk/bin/gcloud init

### Create instance

$  gcloud compute instances create-with-container [INSTANCE_NAME] --container-image [DOCKER_IMAGE]


### Docker setup
$ gcloud auth configure-docker

#### Push image

$ gcloud projects create curatedlist-project --name=curatedlist --set-as-default
$ gcloud sql instances create curatedlist --region="europe-west1" --zone="europe-west1-d" --no-assign-ip

### Deploy new version 
$ gcloud compute instances update-container nginx-vm --container-image gcr.io/cloud-marketplace/google/nginx1:latest

### SSH container

$ gcloud compute ssh [INSTANCE_NAME] --container [CONTAINER_NAME]

### Visualize logs

$ gcloud compute instances get-serial-port-output [INSTANCE_NAME]
