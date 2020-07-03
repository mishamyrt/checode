.PHONY: lint clear checode

all: dist/checode

dist:
	mkdir dist

clear:
	rm -rf dist

checode: dist/checode
	dist/checode

lint:
	golangci-lint run -E lll -E misspell -E prealloc -E stylecheck -E gocritic

dist/checode: dist
	env CGO_ENABLED=0 GOGC=off go build -v -o dist/checode ./checode.go