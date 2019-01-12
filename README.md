# grpc-go-api
Sample gRPC API server written in golang using Cloud Spanner on GCP (Google Cloud Platform) as storage layer.
This sample uses GKE (Google Kubernetes Engine) on GCP to host the gRPC API server.

To get started with GCP, please follow this [link](https://cloud.google.com/gcp/getting-started/).

## Define your project ID:
```
export PROJECT_ID=$(gcloud config list project --format "value(core.project)")
```

## Build Docker image using as parameters your own GCP project info:
Update the below command with your own GCP Project ID as well as Cloud Spanner instance/database names.
```
$ docker build \
--build-arg projectid=test-project \
--build-arg instance=test-instance \
--build-arg database=game-a  \
-t asia.gcr.io/${PROJECT_ID}/grpc-go-api:v1 .
```

## Then push the new Docker image to GCR (Google Container Repository):
```
$ gcloud docker -- push asia.gcr.io/${PROJECT_ID}/grpc-go-api:v1
```

## You can verify that the image has been successfully pushed using this commmand:
```
$ gcloud container images list-tags asia.gcr.io/${PROJECT_ID}/grpc-go-api
```

## Deploy the Web app to GKE (first deployment then service)
```
$ kubectl create -f grpcapi-deployment.yaml
$ kubectl create -f grpcapi-service.yaml
```

## Check that the Deployments and Services are created
```
kubectl get deployments
kubectl get svc
```

## Test your gRPC API using the client app located in cmd/client-grpc folder

```
$ EXTERNAL-IP=$(kubectl get service grpc-go-api-service --output jsonpath="{.status.loadBalancer.ingress[0].ip}")
$ go run main.go -grpc-address=${EXTERNAL-IP} -grpc-port=8080
```

## Secure your gRPC API using Cloud Endpoints
To first deploy Cloud Endpoints config without authentication, use api_config.yaml config file.

Create a proto descriptor file
```
$ protoc --include_imports --include_source_info --descriptor_set_out deployments/endpoints/player.pb api/proto/v1/player.proto
```

Replace PROJECT_ID with your own GCP Project ID in the following config files:
- deployments/endpoints/api_config.yaml
- deployments/k8s/grpcapi-deployment.yaml
- deployments/k8s/grpcapi-endpoints-deployment.yaml

Deploy the Endpoints configuration
```
$ cd deployments/endpoints
$ gcloud endpoints services deploy player.pb api_config.yaml
```

Then delete and redeploy your gRPC pods and service using the Cloud Endpoints ESP sidecar container
```
$ kubectl delete deployment grpc-go-api-deployment
$ kubectl delete svc grpc-go-api-service
$ kubectl create -f grpcapi-endpoints-deployment.yaml
$ kubectl create -f grpcapi-endpoints-service.yaml
```

If the deployment is successful, you can access the GCP Console and start seeing metrics from the Cloud Endpoints portal.