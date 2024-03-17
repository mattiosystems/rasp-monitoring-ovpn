GOCMD = go
GOBUILD = $(GOCMD) build
BINARY_NAME = monitoring
OS=linux
ARCH=arm
ARM=7

build:
	GOOS=$(OS) GOARCH=$(ARCH) GOARM=$(ARM) $(GOBUILD) -o $(BINARY_NAME) -v