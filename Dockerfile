FROM golang

ADD . /go/src/github.com/mjd/basic-svc

# Build our app for Linux
RUN go install github.com/mjd/basic-svc

# Run the basic-svc command by default when the container starts.
ENTRYPOINT /go/bin/basic-svc

# Document that the service listens on port 8083
EXPOSE 8083