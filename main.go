package main

import (
	"flag"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	portptr := flag.String("port", ":9125", "Provide the port to listen on")
	addr := flag.String("web.listen-address", "127.0.0.1", "The address to listen on for HTTP requests.")
	flag.Parse()
	port := *addr + *portptr
	//Create a new instance of the foocollector and
	//register it with the prometheus client.
	col := collector.newLxdCollector()
	prometheus.MustRegister(col)

	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on address ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
