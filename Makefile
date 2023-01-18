PROJECTNAME=$(shell basename "$(PWD)")
DIR=$(shell pwd)
GOBIN=$(DIR)/bin

# Redirect error output to a file, so we can show it in development mode.
STDERR=.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID=.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

load-env:
ifneq ("$(wildcard ./.env)","")
include .env
export
endif

restart: stop start
start: compile start-server
stop: stop-server

start-cli: compile
	$(eval CLI_BIN := $(wildcard ./bin/cli*))
	@-$(CLI_BIN)

start-server: load-env
	@echo ">  $(PROJECTNAME) is available at http://$(APP_HOST):$(APP_PORT)";
	$(eval SERVER_BIN := $(wildcard ./bin/server*))
	@-$(SERVER_BIN) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/\>  PID: /"

stop-server:
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)

restart-server: stop-server start-server

watch:
	@GOBIN=$(GOBIN) yolo -i . -e vendor -e bin -c "$(run)"

compile:
	@touch $(STDERR)
	@rm $(STDERR)
	@-"$(MAKE)" -s go-compile 2> $(STDERR)
	@cat $(STDERR) 1>&2
	@rm $(STDERR)

go-compile: go-clean go-mod-vendor go-build

go-build:
	@echo "> Building apps:"
	$(eval MODULES := $(wildcard cmd/*))
	@for module in $(MODULES) ; do \
		echo ">> $$module" ; \
		GOBIN=$(GOBIN) go build -mod=vendor -o bin/  $(PROJECTNAME)/$$module ; \
	done

go-mod-vendor:
	@echo "> Downloading vendor packages"
	GOBIN=$(GOBIN) go mod vendor

go-clean:
	@echo "> Cleaning build cache"
	GOBIN=$(GOBIN) go clean

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo