package main

import (
	"flag"
	"github.com/ckluenter/aleph-exporter/pkg/observe"
	"github.com/ckluenter/aleph-exporter/pkg/web"
	"net/http"
	"github.com/facebookgo/grace/gracehttp"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	r := web.NewRouter()
	r = observe.RegisterPrometheus(r)
	srv := &http.Server{
		Addr:              *addr,
		Handler:           r,
		ReadTimeout:       0,
		WriteTimeout:      0,
	}
	gracehttp.Serve(srv)
}

