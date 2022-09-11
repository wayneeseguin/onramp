all: build stop start parse load

.PHONY: start stop parse load reports psql bash

PORT=3000

start:
	./bin/docker start

stop:
	./bin/docker stop
	
build:
	./bin/docker build

logs:
	./bin/docker logs

parse:
	./bin/docker parse

load:
	./bin/docker pload

psql:
	./bin/docker psql

pshell:
	./bin/docker pshell

shell:
	./bin/docker shell

test:
	./bin/docker tests

web:
	open http://127.0.0.1:${PORT}

clean: 
	./bin/clean

