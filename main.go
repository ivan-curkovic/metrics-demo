package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var rootCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "root_request_count",
		Help: "No of request handled by root handler",
	},
)

func root(c *gin.Context) {
	rootCounter.Inc()
	c.String(http.StatusOK, "hello")
}

func main() {
	prometheus.MustRegister(rootCounter)

	r := gin.Default()

	r.GET("/", root)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8080")
}
