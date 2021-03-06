# Go API Starter Pack

This repo contains the basics of what you need to get started in created a Go web API. It has an example api endpoint and instructions about how to test it, and then adjust it for your own needs when adding more endpoints.

# Getting Started

In order to use this you need to have a working [go](https://golang.org/doc/install) environment set up and [docker](https://docs.docker.com/get-started/) installed. Once you have this set up you can clone this repo into your go environment by running:

    git clone https://github.com/mmcken3/go-api-starter.git

Once in the repo you can run the make commands to build or run the tool:

    make build (this will only install the go code)
    make cert (only need to run once, sets up local certs for https)
    make run (this will install and run the api)

When the api is up and running successfully you will see a message that says:

    go-api-starter | Go API Starter is now running!

With this there are two test API endpoints that can be hit at the address `http://localhost:8002`

## Go Dep

This api starter is now set up to be using go mods with the vendors committed as of v2.0. If you would like to use the version of this api that uses dep it was switched over after release v1.3. Here are steps for how to get yourself in a new branch with the go api starter that was using dep.

```bash
git clone https://github.com/mmcken3/go-api-starter.git
cd go-api-starter
git checkout tags/v1.3 -b api-with-dep
```

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

# Open Sorce Go Packages Used

- [github.com/go-chi/chi](https://github.com/go-chi/chi)
- [github.com/go-chi/render](https://github.com/go-chi/render)
- [github.com/joho/godotenv](https://github.com/joho/godotenv)
- [github.com/kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig)
- [github.com/pkg/errors](https://github.com/pkg/errors)
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
