APP     := procpipe
DIST    := dist
LDFLAGS := -ldflags="-s -w"

.PHONY: build build-all clean

build:
	CGO_ENABLED=0 go build $(LDFLAGS) -o $(DIST)/$(APP) .

build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(DIST)/$(APP)-linux-amd64 .

build-linux-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(DIST)/$(APP)-linux-arm64 .

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(DIST)/$(APP)-windows.exe .

build-darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(DIST)/$(APP)-darwin-amd64 .

build-darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(DIST)/$(APP)-darwin-arm64 .

build-all: build-linux-amd64 build-linux-arm64 build-windows build-darwin-amd64 build-darwin-arm64

clean:
	rm -rf $(DIST)
