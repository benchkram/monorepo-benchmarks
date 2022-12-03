MAKEFLAGS := --no-print-directory --always-make

build:
	go build -o rc
	cp rc ../rc

run: build
	rm -rf test/
	./rc
