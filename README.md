# json-schema-service
A service for validating JSON schemas

# Design Decisions
I chose to containerise this service with Docker so to that it can run in any environment.
This also means that this project could be deployed to Kubernetes.

I also chose to write tests from the beginning of the project. Having previously worked in 
software testing I believe quality must be built into a project (not added at the end).
Write tests at the same time as the code, ensures that code is written for testability.

## Main
In `main.go`, `func main()` was written to have a tiny abstraction. The `main` function
does not perform any setup at all, it only calls `run` then handles any errors it 
receives. This avoids the need for repeating repetitive error handling in the `main` function. 
This also allows the `run` function to simply return an error if an error is encountered 
during setup, again avoiding the need for repetitive error handling at each setup step.

## Server
I chose to create a `server` struct and store dependencies inside it. This makes it very
clear what a server needs to run. This also avoids holding dependencies in global state.
By not using global variables, this project is automaticaly safe from a whole class of
bugs and errors.

I have implemented a `ServeHTTP` method on `server` so that it satisfies the `http.Handler`
interface. This means a `server` can be used wherever a `http.Handler` is needed, e.g.
`ListenAndServe`.

I also chose to create a `newServer` function, which is a constructor for `server`. 
As server does not have many dependencies, they are passed to this constructor as 
arguments. The `server` constructor sets up routes for the server as this only needs
to be done once at creation. 

## Routes
I chose to move all routing code to `routes.go`. If a bug report came in, a URL would 
usually be included in the report. Having all routes in one place allows for very fast
identification of the handler which needs to be looked at.

## Handlers
I chose to have all handlers as methods on the `server` struct. This gives each handler
access to the dependencies inside the `server` struct. There is a trade off here which 
is that each incomming http request will run in a new goroutine, with each goroutine 
having pointer access to the `server` struct. Pointers aren't concurrency safe so code
does need to be written with data races in mind. For this project, the advantages were
greater than the disadvantages which is why I chose to do this.

Each handler method returns an anonymous `http.HandlerFunc` handler function following
a closure pattern. Each anonymous handler function has its request and response types 
declared inside. This was done to not only make the types easy to find, but also to 
reduce namespace pollution as request and reponse types no longer need unique names.

# Running this project
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
