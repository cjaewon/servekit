# Servekit
⚡️ Micro HTTP server for serving static files. 

## Motivation
Imagine that you have a frontend static files, backend server and nignix(reverse proxy), which are running in the Docker Container. The problem occurs at this point, how to serve static files to reverse proxy? Install 1 more nignix in frontend container, that is overkill. this is why Servekit is made.

## Quick Start
#### Start with local
Grab the latest binary from the [releases page](https://github.com/cjaewon/servekit/releases)
```sh
./servekit # serving ./static directory
```

#### Start with Docker
```Dockerfile
FROM cjaewon/servekit:1.0.0
COPY ./static /static

EXPOSE 3000
```

#### Start with docker compose
```yml
version: "3.8"

services:
  app:
    image: cjaewon/servekit:1.0.0
    volumes:
      - ./static:/static
    environment:
      - SERVEKIT_SERVER_PORT: :3000
      - SERVEKIT_SERVER_PATH: ./static
      - SERVEKIT_SERVER_MODE: none
      - SERVEKIT_SERVER_OVERVIEW: false
```

## Config
### With config file
You can set a port, serving path and etc.. with config file or environment variable.
Servekit scans a `.servekit.toml` at `$HOME`, `.` Directories

```toml
### Config file sample
[server]
  port=":3000" # :3000 (default)
  path="./static" # ./static (default)
  
  mode="none" # none (default), client-side-rendering
  # if you are using client-side-rendering, 404 page will return index.html

  overview=false # false (default), true
  # if your are using true, directory's file list will be show
```

### With environment variable
with docker, you can set environment variable like below. (A undefined config will use a default value)
```Dockerfile
ENV SERVEKIT_SERVER_PORT :3000
ENV SERVEKIT_SERVER_PATH ./static
ENV SERVEKIT_SERVER_MODE none
ENV SERVEKIT_SERVER_OVERVIEW false
```

