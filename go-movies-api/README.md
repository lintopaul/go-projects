# Go Movies API

This is a simple Movies API without a database and entries are stored in memory.

## How to run

```
go build
```

e.g.
```
➜  go-movies-api git:(main) ✗ go build
➜  go-movies-api git:(main) ✗ ls
go-movies-api main.go
➜  go-movies-api git:(main) ✗ ./go-movies-api 
Starting server at port 8000
```


## Endpoints

- GET /movies
- GET /movies/{id}
- POST /movies
- PUT /movies/{id}
- DELETE /movies/{id}

## Usage

`http://localhost:8000/movies`