.PHONY: help build build-native run

help:
	@echo "build - to build project"
	@echo "build-native - compile rust dependecy"
	@echo "run - to run project"

build:
	go build cmd/codewars-api/main.go

build-native:
	cargo build --manifest-path internal/native/Cargo.toml --release
	mv internal/native/target/release/libnative.so internal/service/

run: build-native
	go run cmd/codewars-api/main.go
