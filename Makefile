#!/bin/bash

build:
	@echo "Building backend-api"
	cd ./backend-api && go build -o api .\

run:
	(CONFIG_PATH=./etc/dev.yml ./api)

recompile-docker:
	docker compose up --force-recreate --build
