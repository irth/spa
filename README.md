# `spa`

Simple HTTP server for serving single page apps locally (not to be used in
production.)

Meant to be an alternative for `python3 -m http.server`, where instead of
returning 404 it returns `index.html` for files that do not exist, so that SPAs
which use history API work.

## Usage

```shell
$ spa -h
Usage of spa:
  -addr string
        address to listen on (default ":8000")
  -dir string
        dir to serve (default ".")
```

## Installation

Download binary from [releases](https://github.com/irth/spa/releases), `chmod +x`, put it in path.

## Building

Clone repo, `make`.
