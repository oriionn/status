build:
	go build -o bin/status

test:
	go run .

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/status-linux-x86_64
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o bin/status-linux-x86
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/status-linux-aarch64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o bin/status-linux-armv7
	cp bin/status-linux-armv7 bin/status-linux-armhf
