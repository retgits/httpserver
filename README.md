# httpserver
A simple httpserver in the idea of `python -m SimpleHTTPServer`

## usage
```
$ httpserver [--port] [--root]
```

* port is the HTTP port that will be used (defaults to `8080`)
* root is the root directory (defaults to current dir `.`)

## install
```
$ go get -u github.com/retgits/httpserver
```

## thanks
The code is originally from [this](https://www.chrismytton.uk/2013/07/17/golang-static-http-file-server/) blog