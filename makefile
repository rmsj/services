SHELL := /bin/bash

run:
	go run app/sales-api/main.go

runadmin:
	go run app/admin/main.go	

kill:
	lsof -t -i tcp:3000 | xargs kill

tidy:
	go mod tidy
	go mod vendor