PROJECT := uwdog

OS ?= linux
ARCH ?= amd64

PLATFORMS := \
	linux/amd64 \
	linux/arm64 \
	windows/amd64 \
	windows/arm64 \
	darwin/amd64 \
	darwin/arm64


run:
	go run cmd/uwdog/main.go

test: 
	go test -cover ./...

fmt:
	go fmt ./...

clean:
	rm -rf build/

build:
	@OUTPUT=build/$(OS)/$(ARCH)/$(PROJECT); \
	if [ "$(OS)" = "windows" ]; then OUTPUT="$$OUTPUT.exe"; fi; \
	echo "Building $$OUTPUT ..."; \
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags='-s -w' -o $$OUTPUT cmd/uwdog/main.go

build-all:
	@for platform in $(PLATFORMS); do \
		OS=$${platform%/*}; \
		ARCH=$${platform#*/}; \
		OUTPUT="build/$${OS}/$${ARCH}/$(PROJECT)"; \
		if [ "$$OS" = "windows" ]; then OUTPUT="$$OUTPUT.exe"; fi; \
		echo "Building $$OUTPUT ..."; \
		CGO_ENABLED=0 GOOS=$$OS GOARCH=$$ARCH go build -ldflags='-s -w' -o $$OUTPUT cmd/uwdog/main.go; \
	done
