[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/ToJen/go-shortii)
[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/github.com/ToJen/go-shortii)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/ToJen/go-shortii/blob/master/LICENSE)


# go-shortii
Simple URL Shortener (using TinyURL's Create API) to practice Go.

Feel free to make PRs! ;)

## Run
- directly
    - `go run cmd/main.go`
- or build first
    - `go build -o bin/go-shortii cmd/main.go`
    - `./bin/go-shortii`

## Test
`go test ./pkg/shorten`
