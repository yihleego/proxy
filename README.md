# Proxy

[![GoDoc](https://godoc.org/github.com/yihleego/proxy?status.svg)](https://godoc.org/github.com/yihleego/proxy)
[![Go Report Card](https://goreportcard.com/badge/github.com/yihleego/proxy)](https://goreportcard.com/report/github.com/yihleego/proxy)

A reverse proxy for HTTP written in Go.

## Usage

```go
bind := "0.0.0.0:8080"
remote := "http://<host>:<port>"
p := &proxy{bind: bind, remote: remote}
p.start()
```

## License

This project is under the MIT license. See the [LICENSE](LICENSE) file for details.
