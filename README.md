# go-data
Device data API with Go Micro

Note: this is just a simple REST API to get fake data for a device. This API checks with a device API if the device is active. The device API is written with Go Micro.

The Dockerfile uses the empty scratch image and requires that you build a static exe with the following command:

`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .`

To run in a container, you can use environment variables to specify options like the registry to use like so:

`docker run -d -p 8080:8080 --env MICRO_REGISTRY=mdns image_tag`

Use go-data-dep.yaml to deploy to Kubernetes like so:

`kubectl create -f go-data-dep.yaml`

