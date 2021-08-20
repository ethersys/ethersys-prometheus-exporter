package main

import (
	"exporter/collector"
	"log"
	"crypto/subtle"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func basicAuth(username string, password string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="metrics"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised\n"))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	IP := "127.0.0.1"
	PORT := "8080"
	basicAuthUser := ""
	basicAuthPass := ""
	enableAuth := false
	if len(os.Args) > 2 {
		IP = os.Args[1]
		PORT = os.Args[2]
	}
	if len(os.Args) > 3 {
		basicAuthUser = os.Args[3]
		basicAuthPass = os.Args[4]
		enableAuth = true
	}

	foo := collector.NewwRamCollector()
	prometheus.MustRegister(foo)
	listen := IP + ":" + PORT

	if enableAuth {
		http.Handle("/metrics", basicAuth(basicAuthUser, basicAuthPass, promhttp.Handler()))
	} else {
		http.Handle("/metrics", promhttp.Handler())
	}

	log.Fatal(http.ListenAndServe(listen, nil))
}
