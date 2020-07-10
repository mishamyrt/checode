.PHONY: lint clear docker-image

GC = go build
GO_FLAGS = CGO_ENABLED=0 GOGC=off
ENTRYPOINT = ./checode.go

all: dist

dist:
	mkdir dist
	make dist/checode_darwin64
	make dist/checode_linux64
	make dist/checode_windows64.exe

clear:
	rm -rf dist

lint:
	golangci-lint run -E lll -E misspell -E prealloc -E stylecheck -E gocritic

dist/checode_darwin64:
	env GOOS=darwin \
		GOARCH=amd64 \
		$(GO_FLAGS) $(GC) -o dist/checode_darwin64 $(ENTRYPOINT)

dist/checode_linux64:
	env GOOS=linux \
		GOARCH=amd64 \
		$(GO_FLAGS) $(GC) -o dist/checode_linux64 $(ENTRYPOINT)

dist/checode_windows64.exe:
	env GOOS=windows \
		GOARCH=amd64 \
		$(GO_FLAGS) $(GC) -o dist/checode_windows64.exe $(ENTRYPOINT)

docker-image:
	docker build --pull --target checode .
