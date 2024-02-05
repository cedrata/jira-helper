build:
	go build -C ./main/app -o ../../bin/jhelp

test:
	go test ./...

coverage:
	go test -coverprofile="cover.out" ./...
