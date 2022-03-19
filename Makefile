.DEFAULT: all

.PHONY: all
all: build run

.PHONY: build
build:
	wire
	go build


.PHONY: run
run:
	./wire-demo