SRC=$(shell find . -name "*.go")

ifeq (, $(shell which golangci-lint))
$(warning "could not find golangci-lint in $(PATH), run: curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh")
endif

.PHONY: fmt

all: fmt test

build:
	$(info ************ building to ./bin/jira-helper *****************)
	go build -o ./bin/jira-helper

test:
	$(info ******************** running tests *************************)
	go test ./...

coverage:
	$(info ******************** checking coverage *********************)
	go test -coverprofile="cover.out" ./...

format:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)
