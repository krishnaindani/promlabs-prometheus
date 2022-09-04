package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	temp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius",
		Help: "The current home temperature in degrees Celsius",
	})

	prometheus.MustRegister(temp)

	temp.Set(42)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
