# Development Notes

## Dev Local

```sh
export KO_DOCKER_REPO="ghcr.io/tonygilkerson"
export VERSION="v0.0.0-dev"

export ABOUT_CONTEXT_FILE="$PWD/internal/env/test-ace-context.yaml"
go run cmd/about/main.go

ko build ./cmd/about

## Dev Cluster

```sh
skaffold dev --default-repo ghcr.io/tonygilkerson 
```

DEVTODO 

- next step mount cm in deployment
- set ABOUT_CONTEXT_FILE on env for container


