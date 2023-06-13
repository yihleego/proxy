package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type handle struct {
	address string
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s %s %s", r.RemoteAddr, r.Method, r.URL.String(), r.Proto, r.UserAgent())
	target, err := url.Parse(h.address)
	if err != nil {
		log.Fatalln(err)
		return
	}
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return dialer.DialContext(ctx, network, addr)
	}
	r.Host = target.Host
	p := httputil.NewSingleHostReverseProxy(target)
	p.ServeHTTP(w, r)
}

type proxy struct {
	bind   string
	remote string
	svr    http.Server
}

func (p *proxy) start() {
	log.Printf("Listening on %s, forwarding to %s", p.bind, p.remote)
	p.svr.Addr = p.bind
	p.svr.Handler = &handle{address: p.remote}
	if err := p.svr.ListenAndServe(); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func (p *proxy) stop() {
	if err := p.svr.Shutdown(nil); err != nil {
		log.Println(err)
	}
}

func main() {
	bind := "0.0.0.0:8080"
	remote := "https://api.github.com"
	p := &proxy{bind: bind, remote: remote}
	p.start()
}
