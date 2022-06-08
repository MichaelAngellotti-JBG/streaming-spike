GO_FILES    = $(shell find . -name vendor -prune -o -name '*.go' -print)
GO_PACKAGE  := https://github.com/MichaelAngellotti-JBG/streaming-spike
APP_NAME    := streaming
VERSION     ?= $(shell git log -n 1 --format=%h)
GIT_HASH    := $(shell echo $(VERSION) | cut -c 1-7)
USER        := $(shell whoami)
TARGET      ?= $(shell uname | tr A-Z a-z)
PID         = /tmp/go-app-$(APP_NAME).pid
APP_BINARY  = $(APP_NAME)-$(TARGET)

$(APP_NAME): ./bin/$(APP_BINARY)

# Build target, can be invoked locally or within Docker build image
./bin/$(APP_BINARY): $(GO_FILES)
ifeq ($(DEV_MODE),true)
	@echo "dev"
	@# Faster build (-i relies on pkg cache)
	GOOS=$(TARGET) go build -i -o ./bin/$(APP_BINARY) cmd/app/*.go
else
	@# Normal build
	@echo "normal"
	GOOS=$(TARGET) go build -o ./bin/$(APP_BINARY) -ldflags "-X main.Version=$(GIT_HASH)" cmd/app/*.go
endif

serve: restart
	$(eval DEV_MODE=true)
	@fswatch -r -e ".*" -i "\\.go$$" -i "\\.yml$$" -o . | xargs -n1 -I{}  $(MAKE) restart || $(MAKE) kill

kill:
	@kill `cat $(PID) 2> /dev/null` 2>/dev/null || true

before:

restart: before kill
	@$(MAKE) ./bin/$(APP_BINARY) DEV_MODE=true
	@./bin/$(APP_BINARY) & echo $$! > $(PID)


################################################################################
############################ Report Card stuff  ################################
################################################################################

.PHONY: vet
vet:
	@echo "Running go vet"
	@go vet ./...

.PHONY: format
format:
	@echo "Running gofmt"
	@gofmt -w -s ${GO_FILES}

.PHONY: spell
spell:
	@echo "Running spell check"
	@misspell ${GO_FILES}

.PHONY: cyclo
cyclo:
	@echo "Running cyclomatic check"
	@gocyclo --over 15 ${GO_FILES}

.PHONY: lint
lint:
	@echo "Running lint check"
	@golint -set_exit_status ${APP_PACKAGES}

.PHONY: ineffassign
ineffassign:
	@echo "Running ineffectual assign check"
	@ineffassign .

.PHONY: install-go-tools
install-go-tools:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2

.PHONY: commit-check
commit-check: vet format spell lint ineffassign cyclo test
	@echo "Done!"


##############################################################

test:
	@go test -count=1 ./...

clean:
	@rm -f ./bin/*

# List all makefile targets
list:
	@$(MAKE) -rpn | sed -n -e '/^$$/ { n ; /^[^ .#][^ ]*:/ { s/:.*$$// ; p ; } ; }' | egrep --color '^[^ ]*'

.PHONY: serve kill restart test clean list
