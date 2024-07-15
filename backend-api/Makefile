#!/bin/bash

build:
	@echo "Building backend-api"
	go build -o api ./cmd/app

build-consumer:
	@echo "Building backend-api event consumers"
	go build -o consumer ./cmd/consumer

run:
	(CONFIG_PATH=./etc/dev.yml ./api)

recompile-docker:
	BUILDKIT_PROGRESS=plain docker compose up --force-recreate --build
