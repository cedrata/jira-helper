build:
	go build -o ./bin/jhelp

test:
	go test ./...

coverage:
	go test -coverprofile="cover.out" ./...
