package main

import (
	"exporter/collector"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	IP := ""
	PORT := ""
	if len(os.Args) < 2 {
		IP = "127.0.0.1"
		PORT = "8080"
	} else {
		IP = os.Args[1]
		PORT = os.Args[2]
	}

	foo := collector.NewwRamCollector()
	prometheus.MustRegister(foo)
	lisen := IP + ":" + PORT
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(lisen, nil))
}
