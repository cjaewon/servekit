# Servekit
⚡️ Micro HTTP server for serving static files. 

## Motivation
Imagine that you have frontend static files, backend server, nignix in Docker Container. the problem arises from this point appear, how to serve static files in another container? Install 1 more nignix in frontend container is overkill. this is why Servekit is made.

## Quick Start
#### Start with local
```sh
# below url is a MacOS binary file 
wget https://github.com/cjaewon/servekit/releases/download/v0.0.1/servekit_0.0.1_darwin_amd64.tar.gz
tar xzvf servekit_0.0.1_darwin_amd64.tar.gz -C .

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