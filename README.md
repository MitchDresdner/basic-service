# basic-svc
### _A basic microservice which responds to GET requests for:_
* /info (help)
* /health
* /version

The default port is configured for **8083**

Build: docker build -t mitchd/basic-svc .

Run:   docker run -itd -p 8083:8083 --name mysvc mitchd/basic-svc