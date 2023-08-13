# Makefile for Stream Consumer API

# Variables
APP_NAME := stream-consumer
GO := go
PORT := 8089
STREAMING_URL := https://stream.upfluence.co/stream

# Targets
.PHONY: run test

run:
	@echo "Running $(APP_NAME)..."
	@STREAMING_URL=$(STREAMING_URL) PORT=$(PORT) $(GO) run cmd/main.go

test:
	@echo "Running tests..."
	@$(GO) test -v ./...
