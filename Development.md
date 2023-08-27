# Development Notes

## Setup

### Podman Machine

```sh
podman machine init --cpus=4 --memory=4000
```

### Create Cluster

This requires port 8080 and 8443 to be available on the host.

```sh
kind create cluster --config kind-ingress.yaml
```

### Install Ingress Controller

```sh
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s

# Patch the controler to listen on ports 8080 and 8443 in place of 80 and 443
kubectl -n ingress-nginx patch svc ingress-nginx-controller --type merge --patch-file nginx-patch.yaml
```

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
# or
skaffold run --default-repo ghcr.io/tonygilkerson 
```

## Test

```sh
open http://about.127.0.0.1.nip.io:8080/env/version/badge
```
