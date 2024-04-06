# Makefile for a s01 project
# Customize the BINARY name to your project's output name

BINARY=s01
GOARCH=amd64

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

# Build the project
build:
	go build -o ${BINARY} ${LDFLAGS}

# Install dependencies
deps:
	go mod download

# Update dependencies
update:
	go get -u ./...
	go mod tidy

# Test the project
test:
	go test -v ./...

# Clean up
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

# Run the project
run:
	go run .

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o ${BINARY}-linux-${GOARCH} ${LDFLAGS}

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=${GOARCH} go build -o ${BINARY}-windows-${GOARCH}.exe ${LDFLAGS}

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=${GOARCH} go build -o ${BINARY}-mac-${GOARCH} ${LDFLAGS}

# .PHONY defines parts of the makefile that are not dependent on any files
# It helps avoid conflicts with files of the same name and improves performance
.PHONY: build deps update test clean run build-linux build-windows build-mac
