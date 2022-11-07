FROM golang:1.19
LABEL maintainer="rosie@rosie.dev"

RUN mkdir /json-schema-service

# Copy go.mod and go.sum files to container
COPY go.* /json-schema-service/ 

WORKDIR /json-schema-service

# Fetch dependencies
RUN go mod download

# Copy go files to container
COPY *.go /json-schema-service/

RUN go build -o app 
CMD ["./app"]
