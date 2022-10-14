# Go HTTP Server

This is a simple Go http server

## How to run

```
go build
```

e.g.
```
➜  go-http-server go build
➜  go-http-server ls
README.md      go-http-server server.go      static
➜  go-http-server ./go-http-server 
Starting server at port 8080
```


## Endpoints

- GET /
- GET /hello
- GET /form

## Usage

`http://localhost:8080/` is the landing page which points you to the form at `http://localhost:8080/form.html`