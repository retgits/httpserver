// Package main implements a simple httpserver in the idea of
// `python -m SimpleHTTPServer`. Flags:
// port is the HTTP port that will be used (defaults to `8080`)
// root is the root directory (defaults to current dir `.`)
// Code originally from https://www.chrismytton.uk/2013/07/17/golang-static-http-file-server/
package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.String("port", "8080", "Define what TCP port to bind to")
var root = flag.String("root", ".", "Define the root filesystem path")

func main() {
	flag.Parse()
	fmt.Printf("Starting HTTP server on %s\n", *port)
	fmt.Printf("Serving up %s\n", *root)
	panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*root))))
}
