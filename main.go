package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "dockerhub_ratelimit"
)

var (
	addr     = flag.String("listen", "127.0.0.1:9768", "The address to listen on for HTTP requests.")
	username = flag.String("username", "", "Username for use in authentication")
	password = flag.String("password", "", "Password for use in authentication")
)

func main() {
	flag.Parse()
	c := Collector{
		username: *username,
		password: *password,
	}
	prometheus.MustRegister(c)

	http.Handle("/metrics", promhttp.Handler())
	log.Println("start exporter")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
