# gowiki

Golang webapp sandbox at `https://golang.org/doc/articles/wiki/`

## Build and test
```Bash
$ go build -o ./build/wiki wiki.go
$ ./build/wiki
```
Visit test webapp at `http://localhost:8080/view/test`

## Memo

- `./.gitignore`
```
build/
```
- `./Dockerfile`
```
FROM golang:1.9-alpine

ADD . /gowiki
RUN cd /gowiki && go build -o ./build/wiki wiki.go

EXPOSE 8080
ENTRYPOINT ["./build/wiki"]
```
- case sensitive

