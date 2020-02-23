#basic-svc
###_basic microservice which contains a response to GET requests for:_ 
* version
* health
* greet

The default port is configured for **8083**

Build: docker build -t mitchd/basic-svc .

Run:   docker run -itd -p 8083:8083 --name mysvc mitchd/basic-svc