SHELL := /bin/bash

# ==============================================================================
# Building containers

all: sales-api

sales-api:
	docker build \
		-f zarf/docker/dockerfile.sales-api \
		-t sales-api-amd64:1.0 \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”` \

# ==============================================================================
run:
	go run app/sales-api/main.go

runadmin:
	go run app/admin/main.go	

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