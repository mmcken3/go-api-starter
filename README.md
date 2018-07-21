# Go API Starter Pack

This repo contains the basics of what you need to get started in created a Go web API. It has an example api endpoint and instructions about how to test it, and then adjust it for your own needs when adding more endpoints.

# Getting Started

In order to use this you need to have a working [go](https://golang.org/doc/install) environment set up and [docker](https://docs.docker.com/get-started/) installed. Once you have this set up you can clone this repo into your go environment by running:

    git clone https://github.com/mmcken3/go-api-starter.git

Once in the repo you can run the make commands to buidl or run the tool:

    make build (this will only build the containers)
    make run (this will build and run the containers)

When the containers are up and running successfully you will see a message that says:

    go-api-starter | Go API Starter is now running!

With this there are two test API endpoints that can be hit at the address `http://localhost:8002`

# Test API Endpoints

Here is a table of the two test endpoints that are already configfured into the API.

HTTP request | Description   | Link
------------ | ------------- | -------------
**GET** /v1/test    | This will just return a success test message | [http://localhost:8002/v1/test](http://localhost:8002/v1/test)
**POST** /v1/test/submit    | This will return a success message with the user | [http://localhost:8002/v1/test/submit](http://localhost:8002/v1/test/submit)

# Adding New API Endpoints

First you will go into the go/go-api-starter/cmd/api/routes.go file and add a handler to do what you want this endpoint to do. You can look at the two example endpoints for what a general format of a GET and POST request will look like.

Then you can go into the api.go file in the same directory and add the Handler to the server with the format of:

    router.HandleFunc("[endpoint path]", [route function]).Methods("[method]")

# Open Sorce Packages Used

[github.com/gorilla/mux](https://github.com/gorilla/mux)

[github.com/rs/cors](https://github.com/rs/cors)