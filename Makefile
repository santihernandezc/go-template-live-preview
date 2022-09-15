.PHONY: server

wasm:
	GOOS=js GOARCH=wasm go build -o ./go-template-live-preview/src/example.wasm

server:
	go build -o ./server/server ./server
	./server/server

build:
	cd ui && npm run build

all: wasm server