.PHONY help

help:
	@echo "build - to build project"
	@echo "run - to run project"

build:
	cd internal/native/
	cargo build --release
	cd ../..
	mv internal/native/target/release/libnative.so internal/service/
	go build cmd/codewars-api/main.go

run: 
	cd internal/native/
	cargo build --release
	cd ../..
	mv internal/native/target/release/libnative.so internal/service/
	go run cmd/codewars-api/main.go