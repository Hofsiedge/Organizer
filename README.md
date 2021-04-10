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
* Rule (fuzzy logic)

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
	* go-kit
	* gorilla/mux
	* pgx v4
	* zap
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

PostgreSQL can be accessed from your system through `pgbouncer` connection pooler listening on port `6432`. Specify this port when using `psql`, `PgAdmin`, `DataGrip`, etc.

## Production environment

**TODO**


## Release v0.01
- [ ] Basic DB
	- [ ] Task tracker
		- [x] Task (with dependencies)
		- [ ] Resource - basically, a proxy for images, documents, etc [ ] Environment
	- [ ] Event (just as a log for now)
	- [ ] Object
- [ ] Basic REST API
	- [x] Task
	- [ ] Resource
	- [ ] Event
	- [ ] Object
- [ ] Basic CLI client
	- [ ] Task
	- [ ] Resource
	- [ ] Event
	- [ ] Object
- [x] Basic Docker
	- [x] PostgreSQL service
	- [x] Go web app service
	- [x] nginx service
	- [x] docker-compose.yaml
- [ ] Testing
	- [ ] Go
	- [ ] PL/pgSQL
