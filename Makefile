.PHONY: build clean help
SHELL = /bin/sh

IMAGE_TAG = josuablejeru/netor:latest

help:  ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run: ## start running the application
	go run main.go

build:  ## Build and tag the docker image
	docker build -t $(IMAGE_TAG) .

clean:  ## Delete the local docker image
	docker image rm $(IMAGE_TAG)

upgrade:  ## Upgrade all Go packages in $GOPATH
	go get -u all
