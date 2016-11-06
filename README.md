### go-plugin

Simple app to play with the new feature of Golang 1.8 - plugin

More: https://tip.golang.org/pkg/plugin/

### Requirements 

* Golang 1.8 

### Build to library

```
cd printer-impl && go build -buildmode=plugin printer.go
cd processor-impl && go build -buildmode=plugin processor.go
```

### Test

```
go run main.go
```