package main

import (
	"flag"
	"fmt"
	"github.com/ckluenter/aleph-exporter/pkg/observe"
	"github.com/ckluenter/aleph-exporter/pkg/web"
	"github.com/facebookgo/grace/gracehttp"
	"net/http"
	"os"
	"time"
)

var (
	port          = flag.String("listen-port", env("ALEPH_LISTEN_PORT"), "The port to listen on for HTTP requests. e.g. :8080")
	alephHost     = flag.String("aleph-host", env("ALEPH_HOST"), "hostname of aleph")
	alephToken    = flag.String("aleph-token", env("ALEPH_TOKEN"), "token to be used for authentication")
	SkipTLSVerify = flag.Bool("skip-tls-verify",envBool("ALEPH_SKIP_TL_VERIFY"), "Set to true to not verify the authenticity the server's TLS certificate.")
)

func main() {
	if *port == "" {
		*port = ":8080"
	}

	if *alephHost == "" {
		*alephHost = "localhost"
	}
	flag.Parse()
	r := web.NewRouter()
	r = observe.RegisterPrometheus(r)
	srv := &http.Server{
		Addr:         *port,
		Handler:      r,
		ReadTimeout:  0,
		WriteTimeout: 0,
	}

	fmt.Printf("aleph exporter started. Listening on %s and exposing api from %s \n\n", *port, alephUrl(*alephHost))
	go func() {
		for {
			err,requestBody := observe.GetAlephStatus(alephUrl(*alephHost), *alephToken, true)
			if err != nil {
				fmt.Print(err, "\n")
				observe.AlephApiUp(*alephHost,false)
			} else {
				status := observe.ParseAlephStatus([]byte(requestBody))
				observe.UpdatePrometheus(status)
				observe.AlephApiUp(*alephHost,true)
			}
			time.Sleep(time.Duration(10 * time.Second))

		}
	}()

	gracehttp.Serve(srv)
}

func alephUrl(hostName string) string {
	return "https://" + hostName + "/api/2/status"
}

func env(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		return ""
	}
	return value
}

func envBool(name string) bool {
	value, ok := os.LookupEnv(name)
	if !ok {
		return false
	}
	return value == "false"
}