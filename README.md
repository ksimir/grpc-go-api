# grpc-go-api
Sample gRPC API server using Cloud Spanner on GCP as storage layer.
This sample uses GKE (Google Kubernetes Engine) on GCP to host the gRPC API server.

## Build Docker image:
```
$ docker build -t asia.gcr.io/${PROJECT_ID}/grpc-go-api:v1 .
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

