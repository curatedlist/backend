# Step-by-step guide

## Init the go modules

From https://blog.golang.org/using-go-modules
```bash
$ go mod init github.com/curatedlist/backend
```

## Code Layout

https://github.com/golang-standards/project-layout

## Install Gin server

From https://github.com/gin-gonic/gin#installation
```bash
$ go get -u github.com/gin-gonic/gin
```

Run it from:
```bash
$ go run cmd/backend/main.go
```

## Add Dockefile

## Build image

```bash
$ docker build .
```

## Run image

```bash
$ docker run -p8080:8080 IMAGE 
```

## Google Cloud Compute

### Install the SDK

https://cloud.google.com/sdk/docs#mac
```bash
$ ./google-cloud-sdk/install.sh
$ ./google-cloud-sdk/bin/gcloud init
```

### Create instance

```bash
$  gcloud compute instances create-with-container [INSTANCE_NAME] --container-image [DOCKER_IMAGE]
```

### Docker setup

```bash
$ gcloud auth configure-docker
```

#### Push image

```bash
$ gcloud projects create curatedlist-project --name=curatedlist --set-as-default
$ gcloud sql instances create curatedlist --region="europe-west1" --zone="europe-west1-d" --no-assign-ip
```

### Deploy new version 

```bash
$ gcloud compute instances update-container nginx-vm --container-image gcr.io/cloud-marketplace/google/nginx1:latest
```

### SSH container

```bash
$ gcloud compute ssh [INSTANCE_NAME] --container [CONTAINER_NAME]
```

### Visualize logs

```bash
$ gcloud compute instances get-serial-port-output [INSTANCE_NAME]
```

## Google Cloud Run

### Build the image

```bash
gcloud builds submit --tag gcr.io/curatedlist-project/back 
```

### Deploy the image

```bash
$ gcloud run deploy --image gcr.io/curatedlist-project/back --platform managed
```
