build:
	go build -C ./main/core -o ../../bin/jhelp

test:
	go test ./...

coverage:
	go test -coverprofile="cover.out" ./...
