# Organizer

## About

**TODO**

This project is intended to be a task tracker and organizer with a knowledge graph.

Goals:
* Learn Go
* Improve PostgeSQL skills
* Implement some graph ML algoritms
* Learn knowledge base algorithms and inductive (axiomatic) inference


## Design

* Task
* Event
* Possibility

### Tasks

Task states:
```
* Created    (C)
* In process (I)
* Failed     (F)
* Done       (D)
* Planned    (P)
* Archived   (A)
```

Operations:
```
* **Create**:      {}     -> Created
* **Edit**:         *     -> *
* **Fail**:      {I, P}   -> Failed
* **Plan**:      {C, I}   -> Planned
* **Complete**: {C, I, P} -> Done
* **Archive**:      *     -> Archive
* **Restore**:      A     -> *
* **Start**:     {C, P}   -> In process
```


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


**Note that the `./postgres` folder and its content must be in `a+rx` mode in order to be sourced by psql in `postgres` service**

## Development environment

`docker-compose up` starts the services. Note that `postgres` is initialized with `postgres/init.dev.sh`

`docker-compose down -v` shuts the services down and removes volumes

`docker-compose exec postgres ash` opens `ash` terminal in `postgres` container

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
	[ ] Task tracker
		[ ] Task (with dependencies)
		[ ] Resource - basically, a proxy for images, documents, etc
	[ ] Environment
		[ ] Event (just as a log for now)
		[ ] Object
[ ] Basic REST API
	[ ] Task
	[ ] Resource
	[ ] Event
	[ ] Object
[ ] Basic CLI client
	[ ] Task
	[ ] Resource
	[ ] Event
	[ ] Object
[ ] Basic Docker
	[ ] PostgreSQL service
	[ ] Go web app service
	[ ] nginx service
	[ ] docker-compose.yaml
[ ] Testing
	[ ] Go
	[ ] PL/pgSQL
