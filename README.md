# Transactional Key Value Store

- [About the project](#about-the-project)
  - [Design](#design)
- [Usage](#usage)
  - [Layout](#layout)
- [Notes](#notes)

## About the project

The goal of this project is to build an interactive command-line interface to the in-memory transactional key-value store. The data is not persisted to disk when the interactive session is terminated.

## Design

To support nested transactions we use the stack data structure to generalize transaction elements where each transaction has its linked list as a local store. Committing transaction will append data to update the global store subsequently.

This project code structure follows the [Clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and [Domain Driven Design (DDD)](https://martinfowler.com/tags/domain%20driven%20design.html) approach which supports incoming providers or external services for scalability and enable specific service testing.

## Usage

### Prerequisites

- installed [Golang 1.17](https://golang.org/)
- or run using [Docker](https://www.docker.com/)

### Start application

Build from source

```sh
make build // to compile and generate executable binary file.
make test // to run all the tests
```

Build using Docker

```sh
make docker-build
make docker-run
make docker-test
```

### Commands and example usage

The shell accepts the following commands

Set, Get, Delete a value:

```
> SET foo 123
> GET foo
123
> DELETE foo
> GET foo
key not set
```

Count the number of occurrences of a value

```
> SET foo 123
> SET bar 456
> SET baz 123
> COUNT 123
2
> COUNT 456
1
```

Commit and Rollback transaction

```
> BEGIN
> SET foo 456
> COMMIT
> ROLLBACK
> no transaction
> GET foo
> 456
```

## Project Layout

```tree
├── .gitignore
├── CHANGELOG.md
├── Makefile
├── README.md
├── external
│   └── redis
├── cmd
│   └── 
└── pkg
    └── localmeory
```

A brief description of the layout:

- `.gitignore` varies per project, but all projects need to ignore `bin` directory.
- `CHANGELOG.md` contains auto-generated changelog information.
- `README.md` is a detailed description of the project.
- `pkg` places most of functional and logic.
- `test` holds all tests.
- `external` for all external store implementation

## Notes

- Makefile **MUST NOT** change well-defined command semantics, see Makefile for details.