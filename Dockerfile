# Use the official go docker image built on debian.
FROM golang:1.8

# Grab the source code and add it to the workspace.
ADD . /go/src/github.com/benhawker/go-json-api

# Install revel and the revel CLI.
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

# Install Gorm
RUN go get github.com/jinzhu/gorm

# Install PQ
RUN go get github.com/lib/pq

# Use the revel CLI to start up our application.
ENTRYPOINT revel run github.com/benhawker/go-json-api dev 9000

# Open up the port where the app is running.
EXPOSE 9000