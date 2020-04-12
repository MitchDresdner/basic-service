# basic-svc
### _A basic microservice which responds to GET requests for:_
* /crash
* /info (help)
* /health
* /version


The default port is configured for **8083**

Build: docker build -t mitchd/basic-svc .

Run:   docker run -itd -p 8083:8083 --name mysvc mitchd/basic-svc

---

Experiments using Docker desktop for windows v 19.03.08  [buildx](https://docs.docker.com/buildx/working-with-buildx/)

See my notes [Some assembly required](https://bestow.info/some-assembly-required/)

linux Build: 
  docker buildx use nix-arm
  docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 --push -t mitchd/basic-svc-linux -f Dockerfile-linux  .
  
  docker buildx imagetools inspect mitchd/basic-svc-linux:latest 



   
