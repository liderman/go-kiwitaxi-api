# go-kiwitaxi-api
Golang implementation KiwiTaxi API.

[![GoDoc](https://godoc.org/github.com/liderman/go-kiwitaxi-api?status.svg)](https://godoc.org/github.com/liderman/go-kiwitaxi-api)

This package implements the API partner for the service of the order of transfers across the world - [kiwitaxi](http://kiwitaxi.com).

Installation
-----------
	go get github.com/liderman/go-kiwitaxi-api

Usage
-----------
Getting places:
```go
kiwitaxiApi := kiwitaxi.NewKiwitaxiApi("TOKEN")
places, err := kiwitaxiApi.Places()
```

Requirements
-----------

* Need at least `go1.5` or newer.

Documentation
-----------

You can read package documentation [here](http:godoc.org/github.com/liderman/go-kiwitaxi-api).
