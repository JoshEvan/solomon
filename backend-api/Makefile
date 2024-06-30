#!/bin/bash

build:
	@echo "Building backend-api"
	go build -o api .\

run:
	(CONFIG_PATH=./etc/dev.yml ./api)

recompile-docker:
	BUILDKIT_PROGRESS=plain docker compose up --force-recreate --build
