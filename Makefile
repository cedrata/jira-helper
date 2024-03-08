build:
	go build -o ./bin/jira-helper

test:
	go test ./...

coverage:
	go test -coverprofile="cover.out" ./...
