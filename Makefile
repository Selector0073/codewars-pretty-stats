.PHONY help

help:
	@echo "build - to build project"
	@echo "build-native - compile rust dependecy"
	@echo "run - to run project"

build:
	go build cmd/codewars-api/main.go

build-native:
	cd internal/native/
	cargo build --release
	cd ../..
	mv internal/native/target/release/libnative.so internal/service/

run: 
	cd internal/native/
	cargo build --release
	cd ../..
	mv internal/native/target/release/libnative.so internal/service/
	go run cmd/codewars-api/main.go