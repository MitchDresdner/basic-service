# basic-svc
### _A basic microservice which responds to GET requests for:_
* /crash
* /info (help)
* /health
* /version


The default port is configured for **8083**

~~Build: docker build -t mitchd/basic-svc .~~   # switched to experimental buildx

Local Build: docker buildx build --platform linux/amd64 -t mitchd/basic-svc .

Run:   docker run -itd -p 8083:8083 --name mysvc mitchd/basic-svc

---

Experiments using Docker desktop for windows v 19.03.08  [buildx](https://docs.docker.com/buildx/working-with-buildx/)

10 April 2020 - 

Attempting to build multiple cross platform builds resulted in an error

  docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 --push -t mitchd/basic-svc-amd64-arm64-armv7  .
  
  ERROR exporting to image
  
Individual builds completed successfully when built locally, haven't tested results after promoting to github to trigger automatic build yet.

  docker buildx build --platform linux/amd64  --push -t mitchd/basic-svc-amd64  .
  docker buildx build --platform linux/arm64  --push -t mitchd/basic-svc-arm64  .
  docker buildx build --platform linux/arm/v7 --push -t mitchd/basic-svc-armv7  .