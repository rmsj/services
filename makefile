SHELL := /bin/bash

# ==============================================================================
# Building containers

all: sales-api

sales-api:
	docker build \
		-f ./zarf/docker/dockerfile.sales-api . \
		-t sales-api-amd64:1.0 \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”` \

# ==============================================================================
# Running from within k8s/dev

kind-up:
	kind create cluster --image kindest/node:v1.19.1 --name ardan-starter-cluster --config zarf/k8s/dev/kind-config.yaml

kind-down:
	kind delete cluster --name ardan-starter-cluster

kind-load:
	kind load docker-image sales-api-amd64:1.0 --name ardan-starter-cluster

kind-services:
	kustomize build zarf/k8s/dev | kubectl apply -f -

kind-update: sales-api
	kind load docker-image sales-api-amd64:1.0 --name ardan-starter-cluster
	kubectl delete pods -lapp=sales-api

kind-logs:
	kubectl logs -lapp=sales-api --all-containers=true -f

kind-status:
	kubectl get nodes
	kubectl get pods --watch

kind-status-full:
	kubectl describe pod -lapp=sales-api

kind-shell:
	kubectl exec -it $(shell kubectl get pods | grep sales-api | cut -c1-26) --container app -- /bin/sh

kind-database:
	# ./admin --db-disable-tls=1 migrate
	# ./admin --db-disable-tls=1 seed

kind-delete:
	kustomize build zarf/k8s/dev | kubectl delete -f -

# ==============================================================================	
run:
	go run app/sales-api/main.go

runadmin:
	go run app/sales-admin/main.go	

kill:
	lsof -t -i tcp:3000 | xargs kill

tidy:
	go mod tidy
	go mod vendor

# Running tests within the local computer
# count=1 means don't use the cache
test:
	go test -v ./... -count=1
	staticcheck ./...	