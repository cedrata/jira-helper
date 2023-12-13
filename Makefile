build:
	go build -C cmd/v1 -o ../../bin/v1

run:
	./bin/v1

test:
	go test ./...

coverage:
	go test -coverprofile="cover.out" ./...
