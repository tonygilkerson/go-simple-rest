# Development Notes

## Run Local

```sh
export KO_DOCKER_REPO="ghcr.io/tonygilkerson/go-simple-rest"
export VERSION="v0.0.0-dev"

export ABOUT_CONTEXT_FILE="$PWD/internal/env/test-ace-context.yaml"
go run cmd/about/main.go


ko build -L ./cmd/about
```
