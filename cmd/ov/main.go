package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"

	"github.com/directionless/osquery-verifier/pkg/echoserver"
	"github.com/directionless/osquery-verifier/pkg/osquery"
	"github.com/kolide/kit/httputil"
	"github.com/oklog/run"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upath := path.Clean(r.URL.Path)
		fmt.Printf("Request for %s\n", upath)
		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}

func main() {

	var g run.Group
	{
		logic := echoserver.New()
		osqFrontend := loggingMiddleware(osquery.NewTLSServer(logic))

		srv := httputil.NewServer("127.0.0.1:4433", osqFrontend)
		g.Add(func() error {
			fmt.Println("Starting port 443")
			return srv.ListenAndServeTLS("build/server.crt", "build/server.key")
		}, func(err error) {
			srv.Close()
			return
		})
	}

	{
		// this actor handles an os interrupt signal and terminates the server.
		sig := make(chan os.Signal, 1)
		g.Add(func() error {
			signal.Notify(sig, os.Interrupt)
			<-sig
			fmt.Println("beginning shutdown")
			return nil
		}, func(err error) {
			fmt.Printf("process interrupted: %v\n", err)
			close(sig)
		})
	}

	g.Run()
}
