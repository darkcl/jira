.PHONY: build install all

build:
	@go build -o build/jira main.go

install:
	@cp build/jira /usr/local/bin

all: build install
