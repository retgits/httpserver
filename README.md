# httpserver

A simple httpserver in the idea of `python -m SimpleHTTPServer`

## usage

```go
$ httpserver [-port] [-dir] [-tlscert] [-tlskey]
```

* port is the HTTP port that will be used (defaults to `8080`)
* dir is the root directory (defaults to current dir `.`)
* tlscert is the name of the TLS certificate file
* tlskey is the name of the TLS key file

Using all parameters, the command would be

```go
$ httpserver -port 8080 -dir . -tlscert local.pem -tlskey local-key.pem
```

If you're looking for a simple zero-config tool to make locally trusted development certificates with any names you'd like, check out [mkcert](https://github.com/FiloSottile/mkcert)

## install

```go
$ go get -u github.com/retgits/httpserver
```

## thanks

The code is originally from [this](https://www.chrismytton.uk/2013/07/17/golang-static-http-file-server/) blog
