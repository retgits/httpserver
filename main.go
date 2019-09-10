// Package main implements a simple httpserver in the idea of
// `python -m SimpleHTTPServer`.
//
// Flags:
// port is the HTTP port that will be used (defaults to `8080`)
// dir is the root directory (defaults to current dir `.`)
// tlscert is the name of the TLS certificate file
// tlskey is the name of the TLS key file
//
// Code originally from https://www.chrismytton.uk/2013/07/17/golang-static-http-file-server/
package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
)

var (
	port    = flag.String("port", "8080", "TCP port (defaults to 8080)")
	dir     = flag.String("dir", ".", "Filesystem directory to serve (defaults to current dir)")
	tlscert = flag.String("tlscert", "", "Name of the TLS certificate file")
	tlskey  = flag.String("tlskey", "", "Name of the TLS key file")
)

func main() {
	// Parse flags
	flag.Parse()

	// Get the correct directory name
	folder, err := filepath.Abs(filepath.Dir(*dir))
	if err != nil {
		panic(err)
	}

	fmt.Printf("HTTP Server\nPort     : %s\nDirectory: %s\n", *port, folder)

	if len(*tlscert) == 0 {
		fmt.Printf("\nYou're not using HTTPS... You can create selfsigned and trusted certificates for your machine\nusing mkcert: https://github.com/FiloSottile/mkcert")
		panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(folder))))
	}

	panic(http.ListenAndServeTLS(":"+*port, *tlscert, *tlskey, http.FileServer(http.Dir(folder))))
}
