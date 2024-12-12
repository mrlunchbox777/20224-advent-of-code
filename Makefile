.PHONY: default all build run

# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

default: all

all: build run

# builds the project
build:
	@echo "Building..."
	@go build -o bin/2024-advent-of-code main.go

run:
	@echo "make running with args ($(RUN_ARGS))..."
	go run . $(RUN_ARGS)
