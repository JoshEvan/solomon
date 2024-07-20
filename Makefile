#!/bin/bash

build:
	@echo "Building backend-api"
	cd ./backend-api && go build -o api .\

run:
	(CONFIG_PATH=./etc/dev.yml cd ./backend-api && ./api)

recompile-docker:
	docker compose up --force-recreate --build

build-fe:
	cd frontend && docker build -f ./frontend.Dockerfile . -t solomon-fe

run-fe:	
	cd frontend && docker run solomon-fe

run-fe-nodocker:
	cd frontend && NODE_OPTIONS=--openssl-legacy-provider npm run dev
