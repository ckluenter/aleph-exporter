package main

import (
	"flag"
	"fmt"
	"github.com/ckluenter/aleph-exporter/pkg/observe"
	"github.com/ckluenter/aleph-exporter/pkg/web"
	"github.com/facebookgo/grace/gracehttp"
	"net/http"
	"time"
)

var (
	addr          = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	alephHost     = flag.String("aleph-host", "localhost", "hostname of aleph")
	alephToken    = flag.String("aleph-token", "asdfasdfasdfasdf", "token to be used for authentication")
	SkipTLSVerify = flag.Bool("skip-tls-verify", false, "Set to true to not verify the authenticity the server's TLS certificate.")
)

func main() {
	flag.Parse()
	r := web.NewRouter()
	r = observe.RegisterPrometheus(r)
	srv := &http.Server{
		Addr:         *addr,
		Handler:      r,
		ReadTimeout:  0,
		WriteTimeout: 0,
	}

	fmt.Printf("aleph exporter started. Listening on %s and exposing api from %s \n\n", *addr, alephUrl(*alephHost))
	go func() {
		for {
			requestBody := observe.GetAlephStatus(alephUrl(*alephHost), *alephToken, true)
			status := observe.ParseAlephStatus([]byte(requestBody))
			observe.UpdatePrometheus(status)
			time.Sleep(time.Duration(10 * time.Second))

		}
	}()

	gracehttp.Serve(srv)
}

func alephUrl(hostName string) string {
	return "https://" + hostName + "/api/2/status"
}
