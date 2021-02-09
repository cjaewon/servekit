# Servekit
⚡️ Micro HTTP server for serving static files. 

## Motivation
Imagine that you have frontend static files, backend server, nignix in Docker Container. the problem arises from this point appear, how to serve static files in another container? Install 1 more nignix in frontend container is overkill. this is why Servekit is made.

## Quick Start
#### Start with local
Grab the latest binary from the [releases page](https://github.com/cjaewon/servekit/releases)
```sh
./servekit # serving ./static directory
```

#### Start with Docker
```Dockerfile
FROM cjaewon/servekit:0.0.1
COPY ./static ./static

EXPOSE 3000
```

## Config
You can set a port and serving path with environment variable

- `SERVEKIT_PORT` : default is `:3000`
- `SERVEKIT_PATH` : default is `./static`

with docker, you can set environment variable like below
```Dockerfile
ENV SERVEKIT_PORT :3000
ENV SERVEKIT_PATH ./static
```