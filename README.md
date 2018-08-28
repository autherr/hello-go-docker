# Hello Go and Docker

This is a simple test app to play with golang, docker, containers, etc.

## Build

```sh
docker build -t hello-go-docker .
```

This will build the container using a multi stage build.

Final container is based on [distroless](https://github.com/GoogleContainerTools/distroless).

## Run

```sh
docker run --rm -p 3000:3000 hello-go-docker
```

And then browse or `curl` [localhost:3000](http://localhost:3000)
