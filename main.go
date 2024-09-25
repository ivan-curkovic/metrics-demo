package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var rootCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "root_request_count",
		Help: "No of request handled by root handler",
	},
)

func root(w http.ResponseWriter, req *http.Request) {
	rootCounter.Inc()
	fmt.Fprintf(w, "hello")
}

func main() {
	prometheus.MustRegister(rootCounter)

	http.HandleFunc("/", root)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
