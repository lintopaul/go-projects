# Bookstore

This is a simple bookstore implemented in Go with MySQL as the backend store

## Code organization ##
```
.
├── Makefile
├── README.md
├── bookstore
├── cmd
│   └── main
│       └── main.go
├── go.mod
├── go.sum
├── main.go
└── pkg
    ├── config
    │   └── app.go
    ├── controllers
    │   └── book-controller.go
    ├── models
    │   └── book.go
    ├── routes
    │   └── bookstore-routes.go
    └── utils
        └── utils.go
```
## How to Build and Run ##

bookstore can be run as a binary or executable and doesn't require Go tooling on your system. You may build or test it using `make`
e.g.

```
make build
```

For the source code, these are the targets and their purposes for the Makefile:

**make**: without any target after the “make” command will run my default target, in this case “build”.

**build**: builds bookstore binary for darwin(macos) arm64 (this is the default option)

**buildlinux**: builds bookstore binary for linux amd64

**buildmac**: builds bookstore binary for darwin(macos) amd64

**buildmacm1**: builds bookstore binary for darwin(macos) arm64

**clean**: runs “go clean” and then removes object files

**test**: to run the tests.

**cover**”: runs test coverage.”run”: to run the program.

**lint**: to run the linter. 

## Routes ##

- GET /book
- GET /book/{bookID}
- POST /book
- PUT /book/{bookID}
- DELETE /book/{bookID}

## Prerequisite

MySQL database should be setup already and the MySQL connection in config should be adjusted depending on how it is setup (local vs docker container).