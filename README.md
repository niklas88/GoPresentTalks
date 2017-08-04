# Go Talks

This repository holds my talks done wit the present tool from the Go tools
repository. The talks themselves are organized in folders containing
a `<title>.slide`, which contains the talk itself as well as a number of
runnable and non-runnable `.go` files with example code.

## Installing the present tool
Since the presentations include runnable Go examples you need Go installed and
when installing the present tool from source also a workspace set up.

For these steps consult the [Go Getting Started Guide](https://golang.org/doc/install)

After this you can install the present tool with

```
go get -u golang.org/x/tools/present
```

## Viewing a presentation

To view a presentation run `present` within this folder, then open your browser
at [127.0.0.1:3999](http://127.0.0.1:3999) and navigat to the specific slides.
