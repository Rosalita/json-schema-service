# json-schema-service
A service for validating JSON schemas

# Plan
The service should be inside a docker container so that it runs in any environment.
The service needs a datastore, planning to eventually use MongoDb
Will need to use docker-compose to start the service and also start MongoDb
Plan for today:
Write minimal http server that contains three placeholder endpoints.
Containerise this http server into a docker image
Use docker compose and make to run the docker image locally.

# Running this project
I chose to containerise this project. The reason for doing this was to guarantee that it works on any machine. This also means that this project could be deployed to Kubernetes.

## Dependencies
To build and run this project, you will need `docker` and `make`

1. Install [Docker desktop](https://docs.docker.com/desktop/)
2. Install `Make` using the package manager for your OS, e.g.
    * Ubuntu: `sudo apt install make`
    * Mac: `brew install make`
    * Windows: `choco install make`

## Building and starting the service
1. Build the docker image from the dockerfile with the command `make build`.
2. Run docker-compose to start the app and dependencies with the command `make start`.
3. The service can be stopped using the command `make stop`.

While the service is running, the command `docker ps` should show the following active containers.
* json-schema-service
