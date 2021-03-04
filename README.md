# Organizer

## About

**TODO**

This project is intended to be a task tracker and organizer with a knowledge graph.

Goals:
* Learn Go
* Improve PostgeSQL skills
* Implement some graph ML algoritms
* Learn knowledge base algorithms and inductive (axiomatic) inference

## Stack

* Docker, Docker Compose
* Go
	* Fiber
	* pgx v4
* PostgreSQL
	* PL/pgSQL
	* pgTAP
* Javascript
* Tensorflow
* \[Kotlin\]

## Development environment

`docker-compose up` starts the services. Note that `postgres` is initialized with `postgres/init.dev.sh`

`docker-compose down -v` shuts the services down and removes volumes

`docker-compose exec postgres bash` opens `bash` terminal in `postgres` container

`docker-compose exec web ash` opens `ash` terminal in `web` container

PostgreSQL can be accessed from your system through `pgbouncer` connection pooler listening on port `6432`. Specify this port when using `psql` or `PgAdmin`.

**TODO**
* Compilation
* Running tests
	* unit
	* mutational
	* load

## Production environment

**TODO**


## Release v0.01
[ ] Basic DB
	[ ] Task (with dependencies)
	[ ] Resource - basically, a proxy for images, documents, etc
	[ ] Event (just as a log for now)
[ ] Basic REST API
	[ ] Task
	[ ] Resource
	[ ] Event
[ ] Basic CLI client
	[ ] Task
	[ ] Resource
	[ ] Event
[ ] Basic Docker
	[ ] PostgreSQL service
	[ ] Go web app service
	[ ] nginx service
	[ ] docker-compose.yaml
[ ] Testing
	[ ] Go
	[ ] PL/pgSQL
